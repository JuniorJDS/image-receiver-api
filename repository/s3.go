package repository

import (
	"image-receiver-api/infra/aws"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct {
	Uploader *s3manager.Uploader
}

func NewS3Repository() *S3Repository {
	// svc := s3.New(aws.Client())
	// uploader := s3manager.NewUploader(aws.Client())
	session := aws.GetSessionAWS()
	uploader := s3manager.NewUploader(session, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024
	})
	return &S3Repository{
		Uploader: uploader,
	}
}

func (s *S3Repository) Save(file *multipart.FileHeader) error {
	f, _ := os.Open(file.Filename)
	defer f.Close()
	_, err := s.Uploader.Upload(&s3manager.UploadInput{
		Bucket: nil,
		Key:    nil,
		Body:   f,
	})
	if err != nil {
		return nil
	}
	return nil
}
