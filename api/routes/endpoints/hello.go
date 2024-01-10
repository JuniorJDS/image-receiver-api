package endpoints

import "github.com/gofiber/fiber/v2"

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) GetHello(c *fiber.Ctx) error {
	return c.SendString("hello-world")
}
