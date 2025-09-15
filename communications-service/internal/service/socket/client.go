package socket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Message struct {
	Type             string           `json:"type"`
	MessageData      MessageData      `json:"message_data"`
	NotificationData NotificationData `json:"noti_data"`
}

type MessageData struct {
	ConversationID int64  `json:"conversation_id"`
	UserID         int64  `json:"user_id"`
	Status         bool   `json:"status"`
	Message        string `json:"message"`
	Type           string `json:"type"`
}
type NotificationData struct {
	From    int64 `json:"from"`
	To      int64 `json:"to,omitempty"`
	Content any   `json:"content"`
}

type Client struct {
	Hub *Hub

	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte

	Id string

	Client_Id int64
}

func InitClient(hub *Hub, conn *websocket.Conn, id string, client_id int64) *Client {
	return &Client{
		Hub:       hub,
		Conn:      conn,
		Send:      make(chan []byte, 256),
		Id:        id,
		Client_Id: client_id,
	}
}

func (c *Client) ReadPump(rdb *redis.Client) {
	defer func() {
		c.Hub.UnRegister <- c
		c.Conn.Close()

	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.ReplaceAll(message, newline, space))
		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Invalid message format: %v", err)
			continue
		}

		rdb.Publish(context.Background(), "notification", message)
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump(rdb *redis.Client) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			fmt.Printf("Sending message: %s\n", message)
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for range n {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
