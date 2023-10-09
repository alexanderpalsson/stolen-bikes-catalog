package officerstorage

import (
	"context"
	"github.com/counter/officers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Storage struct {
	collection *mongo.Collection
}

func NewStorage(collection *mongo.Collection) *Storage {
	return &Storage{
		collection: collection,
	}
}

func (s *Storage) Create(ctx context.Context, officer officers.Officer) error {
	_, err := s.collection.InsertOne(ctx, officer)
	if err != nil {
		log.Println("Failed to create officer:", err)
		return err
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, id primitive.ObjectID, officer officers.Officer) error {
	_, err := s.collection.UpdateOne(ctx, primitive.M{"id": id}, primitive.M{"$set": primitive.M{"name": officer.Name}})
	if err != nil {
		log.Println("Failed to update officer:", err)
		return err
	}

	return nil
}

func (s *Storage) SetAssigned(ctx context.Context, id primitive.ObjectID, isAssigned bool) error {
	_, err := s.collection.UpdateOne(ctx, primitive.M{"id": id}, primitive.M{"$set": primitive.M{"assigned": isAssigned}})
	if err != nil {
		log.Println("Failed to update officer:", err)
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := primitive.M{"id": id}
	_, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Failed to delete officer:", err)
		return err
	}

	return nil
}

func (s *Storage) FindAvailableOfficer(ctx context.Context) (officers.Officer, error) {
	var officer officers.Officer
	err := s.collection.FindOne(ctx, primitive.M{"assigned": false}).Decode(&officer)
	if err != nil {
		log.Println("Failed to find officer:", err)
		return officers.Officer{}, err
	}

	return officer, nil
}

func (s *Storage) FindAll(ctx context.Context) ([]officers.Officer, error) {
	cursor, err := s.collection.Find(ctx, primitive.M{})
	if err != nil {
		log.Println("Failed to find officer:", err)
		return []officers.Officer{}, err
	}

	var resp []officers.Officer
	if err = cursor.All(ctx, &resp); err != nil {
		log.Println("Failed to find officer:", err)
		return []officers.Officer{}, err
	}
	return resp, nil
}
