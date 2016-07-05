package bot

import (
  "fmt"
  // "io"
  "../private"
  "gopkg.in/sorcix/irc.v1"
)

var server = "irc.chat.twitch.tv"
var port = "6667"

type ConnData struct {
  username, authToken string
}

func Connect() {
  fmt.Println("connecting...")
  conn, connErr := irc.Dial(server + ":" + port)
  if connErr != nil {
    fmt.Println(connErr)
  }

  message1 := &irc.Message{
    // Prefix: &irc.Prefix{
    //   Name: server,
    //   User: "piecedigital",
    //   Host: "piecedigital.net",
    // },
    Params: []string{private.GetAuthToken()},
    Command: "PASS",
    // Trailing: private.GetAuthToken(),
    // EmptyTrailing: true,
  }
  message2 := &irc.Message{
    // Prefix: &irc.Prefix{
    //   Name: server,
    //   User: "piecedigital",
    //   Host: "piecedigital.net",
    // },
    Params: []string{"piecedigital"},
    Command: "NICK",
    // Trailing: private.GetAuthToken(),
    // EmptyTrailing: true,
  }
  // fmt.Println(message.String())
  outerr1 := conn.Encode(message1)
  if outerr1 != nil {
    fmt.Println(outerr1)
  }
  outerr2 := conn.Encode(message2)
  if outerr2 != nil {
    fmt.Println(outerr2)
  }
  incoming, inerr := conn.Decode()
  if inerr != nil {
    fmt.Println(inerr)
  }
  fmt.Println(incoming)
}
