package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

var clients = make(map[string]*Client)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file")
	}
	app := fiber.New()

	// Sử dụng middleware WebSocket
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		clientID := c.Params("id")
		log.Println("Client connected:", clientID)

		client := &Client{
			ID:   clientID,
			Conn: c,
		}

		clients[clientID] = client

		defer func() {
			delete(clients, clientID)
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
			for id, client := range clients {
				if id != clientID {
					if err := client.Conn.WriteMessage(messageType, msg); err != nil {
						log.Println("Error sending message:", err)
					}
				}
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}
