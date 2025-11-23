package handler

import (
	"temporal-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type CheckoutHandler struct {
	UC *usecase.CheckoutUsecase
}

func NewCheckoutHandler(uc *usecase.CheckoutUsecase) *CheckoutHandler {
	return &CheckoutHandler{UC: uc}
}

func (h *CheckoutHandler) Checkout(c *fiber.Ctx) error {
	req := struct {
		OrderID string `json:"order_id"`
		Amount  int64  `json:"amount"`
	}{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err)
	}

	err := h.UC.Checkout(req.OrderID, req.Amount)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
