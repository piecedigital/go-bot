package socket

import (
  "fmt"
  "time"
  "errors"
  "net/http"
  "gopkg.in/sorcix/irc.v1"
  "golang.org/x/net/websocket"
  "github.com/piecedigital/go-bot/bot"
)
var channelName = "piecedigital"

func StartSockets(res http.ResponseWriter, req *http.Request, name string, stopchan chan) int {
  channelName = name
  // channelName = "piecedigital"
  s := websocket.Server{Handler: websocket.Handler(socketHandler)}
  s.ServeHTTP(res, req)
  // select {
  //   case <- stopchan:
  //     fmt.Println("quit")
  // }
  return 200
}

func socketHandler(ws *websocket.Conn) {
  _, err := bot.Connect(sendMessage, receiveMessage, ws, channelName)
  if err != nil {
    fmt.Println(errors.New("Connection struct could not be gotten"))
    return
  }
}

func receiveMessage(conn *irc.Conn, ws *websocket.Conn) {
  for {
    // fmt.Print("\r\n\r\n READING SOCKET MESSAGES \r\n\r\n")
    var in string
    websocket.Message.Receive(ws, &in)
    fmt.Printf("WS Received: %s\n", in)
    command := bot.SendChatMessage(conn, &irc.Message{
      Params: []string{"#" + channelName},
      Trailing: in,
    }, true)
    if command != nil {
      sendMessage(ws, "[PRIVMSG] : [piecedigital] : " + command.(string))
    }
    time.Sleep(time.Millisecond * 100)
  }
}

func sendMessage(ws *websocket.Conn, s string)  {
  websocket.Message.Send(ws, s)
}
