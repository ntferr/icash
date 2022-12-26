package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ntferr/icash/http"
	"github.com/ntferr/icash/settings"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:     settings.GetSettings().Service.Name,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: http.Cors.AllowOrigins,
		AllowMethods: http.Cors.AllowMethods,
		AllowHeaders: http.Cors.AllowHeaders,
	}))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	serverShutdown := make(chan struct{})

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
		serverShutdown <- struct{}{}
	}()

	http.SetupRouter(app)

	serviceAddress := fmt.Sprintf("%s:%s", settings.GetSettings().Service.Host, settings.GetSettings().Service.Port)
	if err := app.Listen(serviceAddress); err != nil {
		log.Panic(err)
	}

	<-serverShutdown
}
