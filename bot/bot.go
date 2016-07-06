package bot

import (
  "fmt"
  // "errors"
  // "io"
  "time"
  "gopkg.in/sorcix/irc.v1"
  "golang.org/x/net/websocket"
  "../private"
  "regexp"
)

const server = "irc.chat.twitch.tv"
const port = "6667"
var botAccount = "piecedigital"

// const (
//   millisecond = time.Duration(1 * 1000 * 1000)
//   second = time.Duration(millisecond * 1000)
//   minute = time.Duration(second * 60)
//   hour = time.Duration(minute * 60)
//   day = time.Duration(hour * 24)
// )

var retries = 0
var retryLimit = 10

func Connect( send func(ws *websocket.Conn, s string), ws *websocket.Conn ) (*irc.Conn, interface{}) {
  fmt.Println("connecting...")
  conn, connErr := irc.Dial(server + ":" + port)
  if connErr != nil {
    fmt.Println(connErr)
  }
  retries = 0
  err := initMsgs(conn)
  if err != nil {
    reconnect(send, ws)
    return nil, err
  }
  for {
    err := checkForMessags(conn, send, ws)
    if err != nil {
      reconnect(send, ws)
      return nil, err
    }
    time.Sleep(time.Millisecond)
  }
  return conn, nil
}

func reconnect( send func(ws *websocket.Conn, s string), ws *websocket.Conn ) {
  if retries < retryLimit {
    Connect(send, ws)
  } else {
    fmt.Println("Damn, can't reconnect :/")
  }
}

func initMsgs(c *irc.Conn) interface{} {
  for _, message := range messageSlice {
    outerr := c.Encode(message)
    if outerr != nil {
      fmt.Println(outerr)
      return outerr
    }
  }
  return nil
}

var messageSlice = []*irc.Message{
  &irc.Message{
    Params: []string{private.GetAuthToken()},
    Command: "PASS",
  },
  &irc.Message{
    Params: []string{"piecedigital"},
    Command: "NICK",
  },
  &irc.Message{
    Params: []string{"#piecedigital"},
    Command: "JOIN",
  },
}

func checkForMessags( c *irc.Conn, send func(ws *websocket.Conn, s string), ws *websocket.Conn ) interface{} {
  incoming, inerr := c.Decode()
  if inerr != nil {
    return inerr
  }
  fmt.Printf("[READ] - %v\n", incoming)
  if incoming.Command == "PRIVMSG" {
    send(ws, "[" + incoming.Command + "] : [" + incoming.Prefix.User + "] : " + incoming.Trailing)
    for command, value := range commands {
      match, _ := regexp.MatchString("^(" + command + ")", incoming.Trailing)
      if match == true {
        SendChatMessage(c, &irc.Message{
          Params: []string{"#" + botAccount},
          Trailing: value,
        })
        return nil
      }
    }
  }
  return nil
}

func SendChatMessage(c *irc.Conn, msg *irc.Message) {
  message := &irc.Message{
    Params: msg.Params,
    Command: "PRIVMSG",
    Trailing: msg.Trailing,
  }
  outerr := c.Encode(message)
  if outerr != nil {
    fmt.Println(outerr)
  }
}

var commands = map[string]string{
  "!butt": "Praise booty!",
}
