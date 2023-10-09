package officers

import "go.mongodb.org/mongo-driver/bson/primitive"

// Officer defines the properties of a police officer
type Officer struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Assigned bool               `json:"assigned"`
}
