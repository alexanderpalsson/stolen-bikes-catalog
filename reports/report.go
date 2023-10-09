package reports

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// StolenBikeReport defines the properties of a bicycle theft StolenBikeReport
type StolenBikeReport struct {
	ID              primitive.ObjectID `json:"id"`
	Brand           string             `json:"brand"`
	FrameNumber     string             `json:"frameNumber"`
	Characteristics string             `json:"characteristics"`
	OfficerID       string             `json:"officer_id"`
	ReportedBy      string             `json:"reported_by"`
	ReportTime      time.Time          `json:"report_time"`
	Resolved        bool               `json:"resolved"`
}
