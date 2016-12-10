package chat

import (
    "encoding/json"
    "fmt"
    "io"
    "log"

    "golang.org/x/net/websocket"
)

type Client struct {
    ws *websocket.Conn
    srv *Server
    nickname string
}

func NewClient(ws *websocket.Conn, srv *Server) *Client {
    return &Client{
        ws,
        srv,
        ws.Request().URL.Query().Get("nickname"),
    }
}

func (c *Client) Nickname() string {
    return c.nickname
}

func (c *Client) SendGreetings() {
    m := NewMessage("admin", fmt.Sprintf("Hello %s", c.nickname))
    c.Send(m)
}

func (c *Client) Send(m *Message) {
    data, err := json.Marshal(m)

    if err != nil {
        log.Printf("can't marshal message: %s\n", err)
    } else {
        _, err := c.ws.Write(data)

        if err != nil {
            log.Printf("can't send message to user '%s': %s\n", c.nickname, err)
        } else {
            log.Printf("message sent to user: %s\n", c.nickname)
        }
    }
}

func (c *Client) Listen() {
    var msg Message

    for {
        err := websocket.JSON.Receive(c.ws, &msg)

        if err == io.EOF {
            // client disconnected
            break
        } else if err != nil {
            log.Printf("error receiving message from user '%s': %s\n", c.nickname, err)
        } else {
            // broadcast message
            log.Println(fmt.Sprintf("new message received from user '%s': '%s'", msg.Nickname, msg.Message))
            c.srv.BroadcastMessage(&msg)
        }
    }
}
