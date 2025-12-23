package main

import (
    "fmt"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    app.Use(cors.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Backend Limbus rodando no Podman! 🚀")
    })

    app.Get("/config", func(c *fiber.Ctx) error {
        dbURL := os.Getenv("DATABASE_URL")
        return c.JSON(fiber.Map{"connected_to": dbURL})
    })

    fmt.Println("Server running on port 8080")
    app.Listen(":8080")
}