package officerstorage

import (
	"context"
	"github.com/counter/officers"
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

func TestStorage_Create(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx     context.Context
		officer officers.Officer
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
			if err := s.Create(tt.args.ctx, tt.args.officer); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Delete(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx context.Context
		id  int
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
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
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
		want    []officers.Officer
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

func TestStorage_FindAvailableOfficer(t *testing.T) {
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
		want    officers.Officer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				collection: tt.fields.collection,
			}
			got, err := s.FindAvailableOfficer(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAvailableOfficer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAvailableOfficer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_SetAssigned(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx        context.Context
		id         primitive.ObjectID
		isAssigned bool
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
			if err := s.SetAssigned(tt.args.ctx, tt.args.id, tt.args.isAssigned); (err != nil) != tt.wantErr {
				t.Errorf("SetAssigned() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Update(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx     context.Context
		id      primitive.ObjectID
		officer officers.Officer
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
			if err := s.Update(tt.args.ctx, tt.args.id, tt.args.officer); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
