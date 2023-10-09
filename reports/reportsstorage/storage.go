package reportsstorage

import (
	"context"
	"github.com/counter/reports"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	collection *mongo.Collection
}

func NewStorage(collection *mongo.Collection) *Storage {
	return &Storage{
		collection: collection,
	}
}

func (s *Storage) Create(ctx context.Context, caseData reports.StolenBikeReport) error {
	_, err := s.collection.InsertOne(ctx, caseData)
	if err != nil {
		log.Println("Failed to create stolen bike case:", err)
		return err
	}
	return nil
}

func (s *Storage) Assign(ctx context.Context, id primitive.ObjectID, officerId string) error {
	update := primitive.M{"$set": primitive.M{"officer_id": officerId}}
	_, err := s.collection.UpdateOne(ctx, primitive.M{"id": id}, update)
	if err != nil {
		log.Println("Failed to update stolen bike case:", err)
		return err
	}
	return nil
}

func (s *Storage) CaseSolved(ctx context.Context, id primitive.ObjectID) error {
	update := primitive.M{"$set": primitive.M{"resolved": true}}
	_, err := s.collection.UpdateOne(ctx, primitive.M{"id": id}, update)
	if err != nil {
		log.Println("Failed to update stolen bike case:", err)
		return err
	}
	return nil
}

func (s *Storage) FindAll(ctx context.Context) ([]reports.StolenBikeReport, error) {
	cursor, err := s.collection.Find(ctx, primitive.M{})
	if err != nil {
		log.Println("Failed to find stolen bike case:", err)
		return []reports.StolenBikeReport{}, err
	}

	var cases []reports.StolenBikeReport
	if err = cursor.All(ctx, &cases); err != nil {
		log.Println("Failed to find stolen bike case:", err)
		return []reports.StolenBikeReport{}, err
	}
	return cases, nil
}

func (s *Storage) FindByID(ctx context.Context, id primitive.ObjectID) (reports.StolenBikeReport, error) {
	var caseData reports.StolenBikeReport
	err := s.collection.FindOne(ctx, primitive.M{"id": id}).Decode(&caseData)
	if err != nil {
		log.Println("Failed to find stolen bike case:", err)
		return reports.StolenBikeReport{}, err
	}

	return caseData, nil
}

func (s *Storage) FindAndAssign(ctx context.Context, officerId string) (reports.StolenBikeReport, error) {
	var caseData reports.StolenBikeReport

	err := s.collection.FindOne(ctx, primitive.M{"officer_id": ""}, &options.FindOneOptions{Sort: primitive.M{"report_time": -1}}).Decode(&caseData)
	if err != nil {
		log.Println("Failed to find stolen bike case:", err)
		return reports.StolenBikeReport{}, err
	}

	update := primitive.M{"$set": primitive.M{"officer_id": officerId}}
	_, err = s.collection.UpdateOne(ctx, primitive.M{"id": caseData.ID}, update)
	if err != nil {
		log.Println("Failed to find stolen bike case:", err)
		return reports.StolenBikeReport{}, err
	}

	return caseData, nil
}
