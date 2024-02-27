package uploader

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"nasdaqvfs/config"
)

//go:generate mockery --name ImageService --filename image.go
type ImageService interface {
	Upload(name string, file io.Reader) (filePath string, err error)
	Resize(path string, width int32, height int32) string
}

type imageService struct {
	host      string
	uploadKey string
	secretKey string
}

func NewImageService(cfg config.ImageService) ImageService {
	return &imageService{
		host:      cfg.Host,
		uploadKey: cfg.UploadKey,
		secretKey: cfg.SecretKey,
	}
}

func (i *imageService) Upload(name string, file io.Reader) (string, error) {
	const path = "/api/v1/upload"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileField, err := writer.CreateFormFile("image", name)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(fileField, file)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", i.host, path), body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", i.uploadKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := struct {
		Error string `json:"error"`
		Path  string `json:"path"`
	}{}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		var b []byte
		b, _ = io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload to img svc failed, code:%d , body: %s", resp.StatusCode, string(b))
	}

	var b []byte
	if b, err = io.ReadAll(resp.Body); err != nil {
		return "", err
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return "", err
	}
	if resp.StatusCode == http.StatusBadRequest {
		return "", fmt.Errorf("upload to img svc failed, code:%d , err: %s", resp.StatusCode, data.Error)
	}

	return data.Path, nil
}

func (i *imageService) Resize(path string, width int32, height int32) string {
	if path == "" {
		return ""
	}
	in := fmt.Sprintf("%s|%d|%d|%s", path, width, height, i.secretKey)
	h := md5.New()
	h.Write([]byte(in))
	hash := hex.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("%s/image/api/v1/resize/%s?width=%d&height=%d&hash=%s", i.host, path, width, height, hash)
}
