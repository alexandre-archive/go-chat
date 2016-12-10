package chat

import (
    "fmt"
    "log"

    "golang.org/x/net/websocket"

)

type Server struct {
    clients map[*Client]bool
}

func NewServer() *Server {
    c := &Server{
        make(map[*Client]bool),
    }

    c.dispatch()

    return c
}

func (s *Server) AddClient(cli *Client) {
    s.clients[cli] = true
    log.Printf("User '%s' connected.\n", cli.Nickname())
}

func (s *Server) RemoveClient(cli *Client) {
    if _, ok := s.clients[cli]; ok {
        s.clients[cli] = false
        delete(s.clients, cli)

        log.Printf("User '%s' disconnected.\n", cli.Nickname())
    }
}

func (s *Server) Handler() websocket.Handler {
    return websocket.Handler(func(ws *websocket.Conn) {
        client := NewClient(ws, s)

        defer func() {
            s.RemoveClient(client)
            err := ws.Close()

            if err != nil {
                log.Println(err)
            }

            s.BroadcastMessage(NewMessage("admin", fmt.Sprintf("User <strong>%s</strong> exited the room.", client.Nickname())))
        }()

        s.BroadcastMessage(NewMessage("admin", fmt.Sprintf("User <strong>%s</strong> joined the room.", client.Nickname())))

        s.AddClient(client)
        client.SendGreetings()
        client.Listen()
    })
}

func (s *Server) BroadcastMessage(msg *Message) {
    log.Printf("broadcasting message from user '%s' to '%d' users.\n", msg.Nickname, len(s.clients))

    for cli, active := range s.clients {
        if active {
            cli.Send(msg)
        }
    }
}

func (s *Server) Shutdown() {

}

func (s *Server) dispatch() {

}
