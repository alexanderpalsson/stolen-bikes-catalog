package reportshandler

import (
	"github.com/counter/officers/officerstorage"
	"github.com/counter/reports"
	"github.com/counter/reports/reportsstorage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	reportStorage  *reportsstorage.Storage
	officerStorage *officerstorage.Storage
}

func NewHandler(reportStorage *reportsstorage.Storage, officerStorage *officerstorage.Storage) *Handler {
	return &Handler{
		reportStorage:  reportStorage,
		officerStorage: officerStorage,
	}
}

type CreateReportRequest struct {
	Brand           string `json:"brand"`
	FrameNumber     string `json:"frameNumber"`
	Characteristics string `json:"characteristics"`
	ReportedBy      string `json:"reported_by"`
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateReportRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	var officerId string
	if officer, err := h.officerStorage.FindAvailableOfficer(ctx); err == nil {
		officer.Assigned = true
		if err = h.officerStorage.SetAssigned(ctx, officer.ID, true); err == nil {
			officerId = officer.ID.Hex()
		}
	}

	report := reports.StolenBikeReport{
		ID:              primitive.NewObjectID(),
		Brand:           req.Brand,
		FrameNumber:     req.FrameNumber,
		Characteristics: req.Characteristics,
		OfficerID:       officerId,
		ReportedBy:      req.ReportedBy,
		ReportTime:      time.Now(),
		Resolved:        false,
	}

	if err := h.reportStorage.Create(ctx, report); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stolen bike case"})
		return
	}

	c.JSON(http.StatusCreated, report)
}

func (h *Handler) ReportCaseSolved(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	ctx := c.Request.Context()
	report, err := h.reportStorage.FindByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find stolen bike case"})
		return
	}

	if err := h.reportStorage.CaseSolved(ctx, report.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stolen bike case"})
		return
	}

	_, err = h.reportStorage.FindAndAssign(ctx, report.OfficerID)
	if err != nil {
		oId, err := primitive.ObjectIDFromHex(report.OfficerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Corrupt officer id in report"})
			return
		}
		_ = h.officerStorage.SetAssigned(ctx, oId, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update officer status"})
		}
	}

	report.Resolved = true
	c.JSON(http.StatusOK, report)
}

func (h *Handler) FindByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	caseData, err := h.reportStorage.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stolen bike case not found"})
		return
	}

	c.JSON(http.StatusOK, caseData)
}

func (h *Handler) FindAll(c *gin.Context) {
	officer, err := h.reportStorage.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Officer not found"})
		return
	}

	c.JSON(http.StatusOK, officer)
}
