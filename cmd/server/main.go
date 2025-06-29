package main

import (
	"log"

	"gpu-4-ai-server/internal/config"
	"gpu-4-ai-server/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	h := handler.NewHandler(cfg.WorkerAddr)

	app := fiber.New()
	app.Get("/", h.HandleHealth)
	app.Post("/api/generate", h.HandleGenerate)
	app.Get("/api/version", h.HandleVersion)
	app.Get("/api/tags", h.HandleTags)
	app.Post("/api/show", h.HandleShow)
	app.Post("/api/pull", h.HandlePull)
	app.Post("/api/create", h.HandleCreate)
	app.Post("/api/delete", h.HandleDelete)

	log.Println("Fiber HTTP server listening on", cfg.HTTPAddr, "proxying to worker at", cfg.WorkerAddr)
	if err := app.Listen(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}
