package socket

import (
  "fmt"
  "errors"
  "net/http"
  "gopkg.in/sorcix/irc.v1"
  "golang.org/x/net/websocket"
  "../bot"
)

func StartSockets(res http.ResponseWriter, req *http.Request) int {
  s := websocket.Server{Handler: websocket.Handler(socketHandler)}
  s.ServeHTTP(res, req)
  return 200
}

func socketHandler(ws *websocket.Conn) {
  conn, err := bot.Connect(sendMessge, ws)
  if err != nil {
    fmt.Println(errors.New("Connection struct could not be gotten"))
    return
  }
  receiveMessage(conn, ws)
}

func receiveMessage(conn *irc.Conn, ws *websocket.Conn) {
  var in string
  websocket.Message.Receive(ws, &in)
  // if err != nil {
  //   fmt.Println(errors.New("Connection struct could not be gotten"))
  //   // return
  // }
  fmt.Printf("Received: %s\n", in)
  bot.SendChatMessage(conn, &irc.Message{
    Params: []string{"#" + "piecedigital"},
    Trailing: in,
  })
}

func sendMessge(ws *websocket.Conn, s string)  {
  websocket.Message.Send(ws, s)
}
