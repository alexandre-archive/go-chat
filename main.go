package main

import (
    "net/http"
    "log"

    "github.com/alexandrevicenzi/go-chat/chat"
)

func main() {
    srv := chat.NewServer()
    defer srv.Shutdown()

    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.Handle("/ws", srv.Handler())

    log.Printf("server running at http://127.0.0.1:8000")

    http.ListenAndServe(":8000", nil)
}
