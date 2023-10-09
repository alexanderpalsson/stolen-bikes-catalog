package officershandler

import (
	"github.com/counter/officers/officerstorage"
	"github.com/counter/reports/reportsstorage"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	type fields struct {
		officerStorage *officerstorage.Storage
		reportStorage  *reportsstorage.Storage
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				officerStorage: tt.fields.officerStorage,
				reportStorage:  tt.fields.reportStorage,
			}
			h.Create(tt.args.c)
		})
	}
}

func TestHandler_Delete(t *testing.T) {
	type fields struct {
		officerStorage *officerstorage.Storage
		reportStorage  *reportsstorage.Storage
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				officerStorage: tt.fields.officerStorage,
				reportStorage:  tt.fields.reportStorage,
			}
			h.Delete(tt.args.c)
		})
	}
}

func TestHandler_FindAll(t *testing.T) {
	type fields struct {
		officerStorage *officerstorage.Storage
		reportStorage  *reportsstorage.Storage
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				officerStorage: tt.fields.officerStorage,
				reportStorage:  tt.fields.reportStorage,
			}
			h.FindAll(tt.args.c)
		})
	}
}

func TestHandler_Update(t *testing.T) {
	type fields struct {
		officerStorage *officerstorage.Storage
		reportStorage  *reportsstorage.Storage
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				officerStorage: tt.fields.officerStorage,
				reportStorage:  tt.fields.reportStorage,
			}
			h.Update(tt.args.c)
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		os *officerstorage.Storage
		rs *reportsstorage.Storage
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.os, tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
