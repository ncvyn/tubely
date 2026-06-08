package main

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func (cfg apiConfig) getAssetPath(videoID string, fileHeader *multipart.FileHeader) (string, error) {
	fileExtension := path.Ext(fileHeader.Filename)
	if fileExtension != ".jpg" && fileExtension != ".jpeg" && fileExtension != ".png" {
		return "", fmt.Errorf("unsupported file type: %s", fileExtension)
	}
	fileName := fmt.Sprintf("%s%s", videoID, fileExtension)
	return path.Join(cfg.assetsRoot, fileName), nil
}

func (cfg apiConfig) getAssetURL(videoID string, fileHeader *multipart.FileHeader) string {
	mediaType := fileHeader.Header.Get("Content-Type")
	return fmt.Sprintf("http://localhost:%s/assets/%s.%s", cfg.port, videoID, mediaType)
}
