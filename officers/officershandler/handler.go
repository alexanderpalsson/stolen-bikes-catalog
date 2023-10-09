package officershandler

import (
	"github.com/counter/officers"
	"github.com/counter/officers/officerstorage"
	"github.com/counter/reports/reportsstorage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Handler struct {
	officerStorage *officerstorage.Storage
	reportStorage  *reportsstorage.Storage
}

func NewHandler(os *officerstorage.Storage, rs *reportsstorage.Storage) *Handler {
	return &Handler{
		officerStorage: os,
		reportStorage:  rs,
	}
}

type CreateOfficerRequest struct {
	Name string `json:"name"`
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateOfficerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing name in create request"})
		return
	}

	ctx := c.Request.Context()
	officer := officers.Officer{
		ID:   primitive.NewObjectID(),
		Name: req.Name,
	}

	_, err := h.reportStorage.FindAndAssign(ctx, officer.ID.Hex())
	if err == nil {
		officer.Assigned = true
	}

	if err := h.officerStorage.Create(ctx, officer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create officer"})
		return
	}

	c.JSON(http.StatusCreated, officer)
}

type UpdateOfficerRequest struct {
	Name string `json:"name"`
}

func (h *Handler) Update(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req UpdateOfficerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.officerStorage.Update(c.Request.Context(), id, officers.Officer{Name: req.Name}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update officer"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.officerStorage.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete officer"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) FindAll(c *gin.Context) {
	resp, err := h.officerStorage.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Officer not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
