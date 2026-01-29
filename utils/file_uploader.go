package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var allowedImageExts = []string{".jpg", ".jpeg", ".png", ".webp"}

func UploadFile(ctx *fiber.Ctx, fieldname, destination string) (string, error) {
	file, err := ctx.FormFile(fieldname)

	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !slices.Contains(allowedImageExts, ext) {
		return "", fmt.Errorf("invalid file type. Allowed: %v", allowedImageExts)
	}

	if file.Size > 5*1024*1024 {
		return "", fmt.Errorf("file size exceeds the 5MB limit")
	}

	uploadPath := fmt.Sprintf("./uploads/%s", destination)
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	newFilename := fmt.Sprintf("%d-%s%s", time.Now().Unix(), uuid.New().String(), ext)
	finalPath := filepath.Join(uploadPath, newFilename)

	if err := ctx.SaveFile(file, finalPath); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	dbPath := fmt.Sprintf("uploads/%s/%s", destination, newFilename)
    return dbPath, nil
}