package client

import (
	"net"
	"golang.org/x/net/websocket"
)


func Client(){
	// create connection
  // schema can be ws:// or wss://
  // host, port – WebSocket server
  conn, err := websocket.Dial("{schema}://{host}:{port}", "", op.Origin)
  if err != nil {
      // handle error
  }
  defer conn.Close()
    // send message
      if err = websocket.JSON.Send(conn, {message}); err != nil {
       // handle error
  }
      // receive message
  // messageType initializes some type of message
  message := messageType{}
  if err := websocket.JSON.Receive(conn, &message); err != nil {
        // handle error
  }
}