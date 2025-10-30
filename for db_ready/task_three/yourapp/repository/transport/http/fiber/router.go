package fiberhttp

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, h *Handler) {
	api := app.Group("/api")

	api.Post("/books", h.CreateBook)
	api.Get("/books", h.ListBooks)
	api.Get("/books/:id", h.GetBook)
	api.Put("/books/:id", h.UpdateBook)
	api.Delete("/books/:id", h.DeleteBook)
}
