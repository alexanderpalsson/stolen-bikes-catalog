package reportshandler

import (
	"github.com/counter/officers/officerstorage"
	"github.com/counter/reports/reportsstorage"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	type fields struct {
		reportStorage  *reportsstorage.Storage
		officerStorage *officerstorage.Storage
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
				reportStorage:  tt.fields.reportStorage,
				officerStorage: tt.fields.officerStorage,
			}
			h.Create(tt.args.c)
		})
	}
}

func TestHandler_FindAll(t *testing.T) {
	type fields struct {
		reportStorage  *reportsstorage.Storage
		officerStorage *officerstorage.Storage
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
				reportStorage:  tt.fields.reportStorage,
				officerStorage: tt.fields.officerStorage,
			}
			h.FindAll(tt.args.c)
		})
	}
}

func TestHandler_FindByID(t *testing.T) {
	type fields struct {
		reportStorage  *reportsstorage.Storage
		officerStorage *officerstorage.Storage
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
				reportStorage:  tt.fields.reportStorage,
				officerStorage: tt.fields.officerStorage,
			}
			h.FindByID(tt.args.c)
		})
	}
}

func TestHandler_ReportCaseSolved(t *testing.T) {
	type fields struct {
		reportStorage  *reportsstorage.Storage
		officerStorage *officerstorage.Storage
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
				reportStorage:  tt.fields.reportStorage,
				officerStorage: tt.fields.officerStorage,
			}
			h.ReportCaseSolved(tt.args.c)
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		reportStorage  *reportsstorage.Storage
		officerStorage *officerstorage.Storage
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
			if got := NewHandler(tt.args.reportStorage, tt.args.officerStorage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
