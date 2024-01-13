package repository

import (
	"fmt"
	"image-receiver-api/infra/aws"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type S3Repository struct {
	Uploader *s3manager.Uploader
}

func NewS3Repository() *S3Repository {
	session := aws.GetSessionAWS()
	uploader := s3manager.NewUploader(session, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024
	})
	return &S3Repository{
		Uploader: uploader,
	}
}

func (s *S3Repository) Save(file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err.Error())
	}
	defer f.Close()

	bucket := "images-claim-check"
	UUID := uuid.NewString()
	destination := fmt.Sprintf("raw/%s/%s", UUID, file.Filename)

	_, err = s.Uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &destination,
		Body:   f,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file to S3. Error: %s", err.Error())
	}
	return nil
}
