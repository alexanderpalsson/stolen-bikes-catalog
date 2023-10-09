package reportsstorage

import (
	"context"
	"github.com/counter/reports"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewStorage(t *testing.T) {
	type args struct {
		collection *mongo.Collection
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStorage(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Assign(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx       context.Context
		id        primitive.ObjectID
		officerId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				collection: tt.fields.collection,
			}
			if err := s.Assign(tt.args.ctx, tt.args.id, tt.args.officerId); (err != nil) != tt.wantErr {
				t.Errorf("Assign() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_CaseSolved(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx context.Context
		id  primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				collection: tt.fields.collection,
			}
			if err := s.CaseSolved(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CaseSolved() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Create(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx      context.Context
		caseData reports.StolenBikeReport
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				collection: tt.fields.collection,
			}
			if err := s.Create(tt.args.ctx, tt.args.caseData); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_FindAll(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []reports.StolenBikeReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				collection: tt.fields.collection,
			}
			got, err := s.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_FindAndAssign(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx       context.Context
		officerId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    reports.StolenBikeReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				collection: tt.fields.collection,
			}
			got, err := s.FindAndAssign(tt.args.ctx, tt.args.officerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAndAssign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAndAssign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_FindByID(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx context.Context
		id  primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    reports.StolenBikeReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				collection: tt.fields.collection,
			}
			got, err := s.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
