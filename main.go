package main

import (
    "os"
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func getPort() string {
    port := os.Getenv("BETA_LISTEN_PORT")
   if port == "" {
       port = "1337"
   }

   return port
}

func main() {
    app := fiber.New()
    port := getPort()

    app.Use(cors.New())
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Beta is running.")
    })

    SetupDB()

    SetupSyncAPIRoutes(app)
    SetupSingletonStatesAPIRoutes(app)
    SetupPuzzleAPIRoutes(app)

    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}