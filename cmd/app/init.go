package app

import (
	websocketpk "websocketchat/cmd/websocket"

	"github.com/gofiber/fiber/v2"
)

type MainApp struct {
	App *fiber.App
}

func NewMainApp() *MainApp {
	app := fiber.New()

	// Sử dụng middleware WebSocket
	app.Use("/ws", websocketpk.WebSocketMiddleware())
	websocketInstance := websocketpk.NewWebSocketServer()

	app.Get("/ws/:id", websocketpk.WebSocketChatController(websocketInstance.Clients))
	return &MainApp{
		App: app,
	}
}
