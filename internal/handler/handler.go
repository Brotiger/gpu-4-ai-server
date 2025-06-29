package handler

import (
	"gpu-4-ai-server/internal/service"
	"gpu-4-ai-server/internal/types"

	"github.com/Brotiger/gpu-4-ai-worker/proto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Svc *service.OllamaService
}

func NewHandler(workerAddr string) *Handler {
	return &Handler{Svc: service.NewOllamaService(workerAddr)}
}

func (h *Handler) HandleGenerate(c *fiber.Ctx) error {
	var req types.GenerateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	resp, err := h.Svc.Generate(&proto.GenerateRequest{
		Model:  req.Model,
		Prompt: req.Prompt,
		Stream: req.Stream,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(types.GenerateResponse{
		Response: resp.Response,
		Done:     resp.Done,
	})
}

func (h *Handler) HandleHealth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

func (h *Handler) HandleVersion(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"version": "0.5.1"})
}

func (h *Handler) HandleTags(c *fiber.Ctx) error {
	resp, err := h.Svc.Tags()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(types.TagsResponse{Models: resp.Models})
}

func (h *Handler) HandleShow(c *fiber.Ctx) error {
	var req types.ShowRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	resp, err := h.Svc.Show(&proto.ShowRequest{Model: req.Model})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(types.ShowResponse{Model: resp.Model, Details: resp.Details})
}

func (h *Handler) HandlePull(c *fiber.Ctx) error {
	var req types.PullRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	resp, err := h.Svc.Pull(&proto.PullRequest{Name: req.Name})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(types.PullResponse{Status: resp.Status})
}

func (h *Handler) HandleCreate(c *fiber.Ctx) error {
	var req types.CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	resp, err := h.Svc.Create(&proto.CreateRequest{Name: req.Name, Modelfile: req.ModelFile})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(types.CreateResponse{Status: resp.Status})
}

func (h *Handler) HandleDelete(c *fiber.Ctx) error {
	var req types.DeleteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	resp, err := h.Svc.Delete(&proto.DeleteRequest{Model: req.Model})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(types.DeleteResponse{Status: resp.Status})
}

// TODO: Реализация балансировки между воркерами
