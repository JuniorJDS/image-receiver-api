package repository

import (
	"fmt"
	"image-receiver-api/config"
	"image-receiver-api/infra/aws"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct {
	Uploader *s3manager.Uploader
	Bucket   string
}

var settings = config.GetSettings()

func NewS3Repository() *S3Repository {
	session := aws.GetSessionAWS()
	uploader := s3manager.NewUploader(session, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024
	})
	return &S3Repository{
		Uploader: uploader,
		Bucket:   settings["BUCKET"],
	}
}

func (s *S3Repository) Save(file *multipart.FileHeader, ID string) error {
	f, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err.Error())
	}
	defer f.Close()

	destination := fmt.Sprintf("raw/%s/%s", ID, file.Filename)

	_, err = s.Uploader.Upload(&s3manager.UploadInput{
		Bucket: &s.Bucket,
		Key:    &destination,
		Body:   f,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file to S3. Error: %s", err.Error())
	}
	return nil
}
