package server

import (
	"net"
	"golang.org/x/net/websocket"
)

func Server() {
	// Initialize WebSocket handler + server
  mux := http.NewServeMux()
      mux.Handle("/", websocket.Handler(func(conn *websocket.Conn) {
          func() {
              for {

                  // do something, receive, send, etc.
              }
          }
      // receive message
  // messageType initializes some type of message
  message := messageType{}
  if err := websocket.JSON.Receive(conn, &message); err != nil {
      // handle error
  }
  // send message
  if err := websocket.JSON.Send(conn, message); err != nil {
      // handle error
  }
}




