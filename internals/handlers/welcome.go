package handlers

import "github.com/gofiber/fiber/v2"

func Welcome(ctx *fiber.Ctx) error {
	return ctx.Render("welcome", nil, "layouts/main")
}