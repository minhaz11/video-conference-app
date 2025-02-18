package server

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/minhaz11/video-conference/internals/handlers"
)

var (
	addr = flag.String("addr", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *addr == "" {
		*addr = ":8080"
	}

	viewEngine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: viewEngine,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.CreateRoom)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.WebSocketRoom, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebSocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebSocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket", handlers.Stream)
	app.Get("/stream/:ssuid/chat/websocket", handlers.Stream)
	app.Get("/stream/:ssuid/viewer/websocket", handlers.Stream)

	err := app.Listen(*addr)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
