package utils

import (
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
