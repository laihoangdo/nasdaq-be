package media

import (
	"context"
	"fmt"
	"strings"

	"github.com/ninetwentyfour/go-wkhtmltoimage"
)

// Export exports images to storage and returns the URL of the image
func (e se) ExportImage(ctx context.Context, cfg Config, template Template) (string, error) {
	html, err := template.HTML()
	if err != nil {
		return "", err
	}

	imgBytes, err := e.generateImageFromHTML(cfg, html)
	if err != nil {
		return "", err
	}

	urlPath, err := e.imgSvc.Upload(cfg.OutputFileName, strings.NewReader(string(imgBytes)))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", cfg.CDNUrl, urlPath), nil
}

// generateImageFromHTML generates image on local storage and returns the images as bytes
func (e se) generateImageFromHTML(cfg Config, htmlStr string) ([]byte, error) {
	imgBytes, err := wkhtmltoimage.GenerateImage(&wkhtmltoimage.ImageOptions{
		Input:      "-",
		Format:     string(cfg.FileType),
		Html:       htmlStr,
		BinaryPath: defaultWkhtmltoimageBin,
		Quality:    cfg.Quality,
	})
	if err != nil {
		return nil, err
	}

	return imgBytes, nil
}
