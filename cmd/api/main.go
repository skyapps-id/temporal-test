package main

import (
	"temporal-test/internal/handler"
	"temporal-test/internal/starter"
	"temporal-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	st := starter.NewTemporalStarter()
	uc := usecase.NewCheckoutUsecase(st)
	h := handler.NewCheckoutHandler(uc)

	app.Post("/checkout", h.Checkout)

	app.Listen(":8081")
}
