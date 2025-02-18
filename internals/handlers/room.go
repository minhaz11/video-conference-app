package handlers

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateRoom(ctx *fiber.Ctx) error {
	return ctx.Redirect(fmt.Sprintf("/room/%s", uuid.New().String()))
}

func Room(ctx *fiber.Ctx) error {
	id := ctx.Params("uuid")

	if id == "" {
		ctx.Status(400)
		return nil
	}

	uuid, suuid, _ := createOrGetRoom(id)
}

func WebSocketRoom(ctx *websocket.Conn) {
	id := ctx.Params("uuid")

	if id == "" {
		return
	}

	_, _, room := createOrGetRoom(id)
}

func createOrGetRoom(id string) (string, string, Room)  {
	
}
