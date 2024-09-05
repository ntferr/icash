package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ntferr/icash/api/controller"
	"github.com/ntferr/icash/api/http"
	"github.com/ntferr/icash/drivers"
	"github.com/ntferr/icash/settings"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           settings.GetSettings().Service.Name,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		EnablePrintRoutes: true,
		ErrorHandler:      fiber.DefaultErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: http.Cors.AllowOrigins,
		AllowMethods: http.Cors.AllowMethods,
		AllowHeaders: http.Cors.AllowHeaders,
	}))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	serverShutdown := make(chan struct{})

	go func() {
		_ = <-sigs
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
		serverShutdown <- struct{}{}
	}()

	drv := drivers.Init()
	controllers := controller.Init(&drv)
	http.SetupRouter(app, controllers)

	serviceAddress := fmt.Sprintf("%s:%s",
		settings.GetSettings().Service.Host,
		settings.GetSettings().Service.Port,
	)

	if err := app.Listen(serviceAddress); err != nil {
		log.Panic(err)
	}

	<-serverShutdown
}
