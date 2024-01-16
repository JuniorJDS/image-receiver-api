package service

import (
	"image-receiver-api/repository"
	"image-receiver-api/schema"
	"image-receiver-api/service/event"
	"mime/multipart"

	"github.com/google/uuid"
)

type FileHandler struct {
	Bucket repository.S3Repository
	Event  event.EventSNS
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		Bucket: *repository.NewS3Repository(),
		Event:  *event.NewEventSNS(),
	}
}

func (f *FileHandler) Process(file *multipart.FileHeader) error {
	fileID := uuid.NewString()
	err := f.Bucket.Save(file, fileID)
	if err != nil {
		return err
	}

	message := schema.Message{
		Type: "image-raw",
		Data: map[string]interface{}{
			"id": fileID,
		},
	}

	err = f.Event.Publish(message)
	if err != nil {
		return err
	}

	return nil
}
