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
	fileExtension, err := getFileExtension(fileHeader.Filename)
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s%s", videoID, fileExtension)
	return path.Join(cfg.assetsRoot, fileName), nil
}

func (cfg apiConfig) getAssetURL(videoID string, fileHeader *multipart.FileHeader) string {
	fileExtension, err := getFileExtension(fileHeader.Filename)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("http://localhost:%s/assets/%s%s", cfg.port, videoID, fileExtension)
}

func getFileExtension(filename string) (string, error) {
	fileExtension := path.Ext(filename)
	if fileExtension != ".jpg" && fileExtension != ".jpeg" && fileExtension != ".png" {
		return "", fmt.Errorf("unsupported file type: %s", fileExtension)
	}
	return fileExtension, nil
}
