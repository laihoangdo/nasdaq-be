package media

import (
	"context"

	"nasdaqvfs/pkg/uploader"
)

// Template forces data to be exported must be written in HTML
//
//go:generate mockery --name Template
type Template interface {
	HTML() (string, error)
}

// ScheduleExporter defines method to export data as required
//
//go:generate mockery --name ScheduleExporter
type ScheduleExporter interface {
	// ExportImage exports template in HTML to output file
	ExportImage(ctx context.Context, cfg Config, template Template) (string, error)
}

// Config defines usertest's request of the custom media file
type Config struct {
	CDNUrl         string
	FileType       FileType
	Width          int
	Height         int
	Quality        int
	OutputFileName string
}

type se struct {
	imgSvc uploader.ImageService
}

// NewScheduleExporterService inits schedule exporter constructor
func NewScheduleExporterService(imgSvc uploader.ImageService) ScheduleExporter {
	return se{imgSvc: imgSvc}
}
