package endpoints

import (
	"github.com/gofiber/fiber/v2"
)

type FileHandle struct {
}

func NewFileHandle() *FileHandle {
	return &FileHandle{}
}

func (f *FileHandle) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return nil
	}

	return c.SendString(file.Filename)
}
