package endpoints

import (
	"image-receiver-api/repository"

	"github.com/gofiber/fiber/v2"
)

type FileHandle struct {
	Bucket repository.S3Repository
}

func NewFileHandle() *FileHandle {
	return &FileHandle{
		Bucket: *repository.NewS3Repository(),
	}
}

func (f *FileHandle) Upload(c *fiber.Ctx) error {
	// TODO: error handler
	file, err := c.FormFile("file")
	if err != nil {
		return nil
	}

	err = f.Bucket.Save(file)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString(file.Filename)
}
