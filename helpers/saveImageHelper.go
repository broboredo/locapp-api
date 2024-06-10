package helpers

import (
	"fmt"
	"github.com/broboredo/locapp-api/handler"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveImage(file *multipart.FileHeader, fileName string, subDirectory string) string {
	uploadedFile, err := file.Open()
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo de imagem: %v\n", err)
		return ""
	}
	defer uploadedFile.Close()

	uploadDir := "./static/images/"
	if subDirectory != "" {
		uploadDir = filepath.Join(uploadDir, subDirectory)
	}
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	ext := filepath.Ext(file.Filename)

	filePath := filepath.Join(uploadDir, fileName+ext)

	newFile, err := os.Create(filePath)
	if err != nil {
		handler.Logger.Errorf("ERROR: SaveImage create file: %v\n", err)
		return ""
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, uploadedFile)
	if err != nil {
		handler.Logger.Errorf("ERROR: SaveImage copy file: %v\n", err)
		return ""
	}

	return filePath
}
