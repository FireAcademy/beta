package main

import (
    "os"
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

//	@title			Beta API
//	@version		1.0
//	@description	The Beta FireAcademy.io API
//	@termsOfService	https://fireacademy.io/terms-and-conditions.txt

//	@contact.name	yakuhito
//	@contact.url	https://twitter.com/yakuh1t0
//	@contact.email	y@kuhi.to

//	@license.name	MIT
//	@license.url	https://github.com/FireAcademy/beta/blob/master/LICENSE

//	@host		https://kraken.fireacademy.io
//	@BasePath	/beta

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-API-Key
//	@description				The API key can also be included in the URL or as a GET parameter - please see [https://docs.fireacademy.io/developers/using-api-keys](https://docs.fireacademy.io/developers/using-api-keys) for more details.

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