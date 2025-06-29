package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gpu-4-ai-server/internal/handler"
	"gpu-4-ai-server/internal/config"
)

func main() {
	cfg := config.Load()

	h := &handler.Handler{WorkerAddr: cfg.WorkerAddr}

	app := fiber.New()
	app.Post("/api/generate", h.HandleGenerate)
	app.Post("/api/tags", h.HandleTags)
	app.Post("/api/show", h.HandleShow)
	app.Post("/api/pull", h.HandlePull)
	app.Post("/api/create", h.HandleCreate)
	app.Post("/api/delete", h.HandleDelete)
	app.Get("/health", h.HandleHealthFiber)

	log.Println("Fiber HTTP server listening on", cfg.HTTPAddr, "proxying to worker at", cfg.WorkerAddr)
	if err := app.Listen(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
} 