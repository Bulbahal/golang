package fiberhttp

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"yourapp/domain"
)

type Handler struct {
	Books domain.BookRepo
}

type createBookDTO struct {
	Title           string `json:"title"`
	AuthorID        int    `json:"author_id"`
	PublicationYear int    `json:"publication_year"`
}

type updateBookDTO struct {
	Title           string `json:"title"`
	AuthorID        int    `json:"author_id"`
	PublicationYear int    `json:"publication_year"`
}

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	var in createBookDTO
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}
	id, err := h.Books.Create(context.Background(), domain.Book{
		Title:           in.Title,
		AuthorId:        in.AuthorID,
		PublicationYear: in.PublicationYear,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (h *Handler) GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	b, err := h.Books.GetByID(context.Background(), id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(b)
}

func (h *Handler) ListBooks(c *fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")

	var f domain.BookFilter
	if title != "" {
		f.Title = &title
	}
	if author != "" {
		f.Author = &author
	}

	list, err := h.Books.List(context.Background(), f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var in updateBookDTO
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}
	if err := h.Books.Update(context.Background(), domain.Book{
		Id:              id,
		Title:           in.Title,
		AuthorId:        in.AuthorID,
		PublicationYear: in.PublicationYear,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := h.Books.Delete(context.Background(), id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
