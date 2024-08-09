package websocketpk

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type WebSocketServer struct {
	Clients map[string]*Client
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{}
}

func WebSocketMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}

func WebSocketChatController(Clients map[string]*Client) func(c *fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		clientID := c.Params("id")
		log.Println("Client connected:", clientID)

		client := &Client{
			ID:   clientID,
			Conn: c,
		}

		Clients[clientID] = client

		defer func() {
			delete(Clients, clientID)
			log.Println("Client disconnected:", clientID)
			_ = c.Close()
		}()

		for {
			messageType, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			log.Printf("Received message from %s: %s\n", clientID, string(msg))

			// Gửi tin nhắn cho tất cả các client khác
			for id, client := range Clients {
				if id != clientID {
					if err := client.Conn.WriteMessage(messageType, msg); err != nil {
						log.Println("Error sending message:", err)
					}
				}
			}
		}
	})
}
