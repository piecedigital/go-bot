package bot

import (
  "fmt"
  // "io"
  "time"
  "../private"
  "gopkg.in/sorcix/irc.v1"
)

var server = "irc.chat.twitch.tv"
var port = "6667"

type ConnData struct {
  username, authToken string
}

// const (
//   millisecond = time.Duration(1 * 1000 * 1000)
//   second = time.Duration(millisecond * 1000)
//   minute = time.Duration(second * 60)
//   hour = time.Duration(minute * 60)
//   day = time.Duration(hour * 24)
// )

var retries = 0
var retryLimit = 10

func Connect() {
  fmt.Println("connecting...")
  conn, connErr := irc.Dial(server + ":" + port)
  if connErr != nil {
    fmt.Println(connErr)
  } else {
    retries = 0
    err := initMsgs(conn)
    if err != nil {
      reconnect()
    } else {
      for {
        err := checkForMessags(conn)
        if err != nil {
          reconnect()
        } else {
          time.Sleep(time.Millisecond)
        }
      }
    }
  }
}

func reconnect() {
  if retries < retryLimit {
    Connect()
  } else {
    fmt.Println("Damn, can't reconnect :/")
  }
}

func checkForMessags(c *irc.Conn) interface{} {
  fmt.Println("This happened...")
  incoming, inerr := c.Decode()
  if inerr != nil {
    return inerr
  }
  fmt.Println(incoming)
  return nil
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


func sendChatMessage(c *irc.Conn, msg string) {
  message := &irc.Message{
    Params: []string{"#piecedigital"},
    Command: "PRIVMSG",
    Trailing: "piecedigital the bot has arrived",
  }
  outerr := c.Encode(message)
  if outerr != nil {
    fmt.Println(outerr)
  }
}
