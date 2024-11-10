package utils

import (
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SaveSingleFileToStorage(c *gin.Context, basePath string, formFileName string, customFileName string) error {
	file, err := c.FormFile(formFileName)
	if err != nil {
		return err
	}
	path := filepath.Join(basePath, customFileName)
	if err := c.SaveUploadedFile(file, path); err != nil {
		return err
	}
	return nil
}

func ConvertVideo(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-c:v", "libx264", "-c:a", "aac", outputPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func FootageFileName(originalName string) string {
	return ("footage_" + originalName)
}
