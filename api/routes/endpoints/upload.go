package endpoints

import (
	"image-receiver-api/service"

	"github.com/gofiber/fiber/v2"
)

type UploadRoute struct {
	FileHandlerService service.FileHandler
}

func NewUploadRoute() *UploadRoute {
	return &UploadRoute{
		FileHandlerService: *service.NewFileHandler(),
	}
}

func (f *UploadRoute) Upload(c *fiber.Ctx) error {
	// TODO: error handler
	file, err := c.FormFile("file")
	if err != nil {
		return nil
	}

	err = f.FileHandlerService.Process(file)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("Ok!")
}
