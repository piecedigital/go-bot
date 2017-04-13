package bot

import (
  "fmt"
  // "errors"
  // "io"
  "time"
  "gopkg.in/sorcix/irc.v1"
  "golang.org/x/net/websocket"
  "github.com/piecedigital/go-bot/private"
  "regexp"
  "github.com/piecedigital/go-bot/bot/commands"
)

const server = "irc.chat.twitch.tv"
const port = "6667"
var botAccount = "piecebot"
var channelName = ""

var retries = 0
var retryLimit = 10
func Connect( send func(ws *websocket.Conn, s string), receive func(conn *irc.Conn, ws *websocket.Conn), ws *websocket.Conn, name string) (*irc.Conn, error) {
  channelName = name
  setCommands()
  fmt.Println("connecting...", name, channelName)
  conn, connErr := irc.Dial(server + ":" + port)
  if connErr != nil {
    fmt.Println(connErr)
    return nil, connErr
  }
  retries = 0
  err := initMsgs(conn)
  if err != nil {
    reconnect(send, receive, ws, name)
    return nil, err
  }
  go receive(conn, ws)
  for {
    // fmt.Println("still going...")
    err := checkForMessags(conn, send, ws)
    if err != nil {
      reconnect(send, receive, ws, name)
      // return nil, err
    }
    time.Sleep(time.Millisecond)
  }
  return conn, nil
}

func reconnect( send func(ws *websocket.Conn, s string), receive func(conn *irc.Conn, ws *websocket.Conn), ws *websocket.Conn, name string) {
  if retries < retryLimit {
    Connect(send, receive, ws, name)
  } else {
    fmt.Println("Damn, can't reconnect :/")
  }
}

func initMsgs(c *irc.Conn) error {
  // fmt.Println("init...", channelName)
  for _, message := range getInitMessages() {
    // fmt.Println("Message on connect", message)
    outerr := c.Encode(message)
    if outerr != nil {
      fmt.Println(outerr)
      return outerr
    }
  }
  return nil
}

func getInitMessages() []*irc.Message {
  // fmt.Println("get init messages...", channelName)
  messageSlice := []*irc.Message{
    &irc.Message{
      Params: []string{private.GetAuthToken()},
      Command: "PASS",
    },
    &irc.Message{
      Params: []string{"piecebot"},
      Command: "NICK",
    },
    &irc.Message{
      Params: []string{"#" + channelName},
      Command: "JOIN",
    },
  }
  return messageSlice
}

func checkForMessags( c *irc.Conn, botPageSend func(ws *websocket.Conn, s string), ws *websocket.Conn ) error {
  incoming, inerr := c.Decode()
  if inerr != nil {
    return inerr
  }
  fmt.Printf("[READ] - %v\n", incoming)
  if incoming.Command == "PRIVMSG" {
    // fmt.Println("[" + incoming.Command + "] : [" + incoming.Prefix.User + "] : " + incoming.Trailing)
    botPageSend(ws, "[" + incoming.Command + "] : [" + incoming.Prefix.User + "] : " + incoming.Trailing)
    // to send chat command
    match, value := checkCommand(c, incoming)
    if match == true {
      fmt.Println("Positive for command via input", incoming)
      SendChatMessage(c, &irc.Message{
        Params: []string{"#" + channelName},
        Command: "PRIVMSG",
        Trailing: value,
      }, false)
    }
  }
  // keep server alive. ping back!
  if incoming.Command == "PING" {
    fmt.Println("Ping? Pong!")
    SendChatMessage(c, &irc.Message{
      // Params: []string{"#" + channelName},
      Command: "PONG",
      // Trailing: "",
    }, false)
  }
  return nil
}

func SendChatMessage(c *irc.Conn, msg *irc.Message, fromInterface bool) interface{} {
  message := &irc.Message{
    Prefix: &irc.Prefix{
      botAccount,
      botAccount,
      botAccount + ".tmi.twitch.tv",
    },
    Params: msg.Params,
    Command: msg.Command,
    Trailing: msg.Trailing,
  }
  fmt.Println("IRC relay -", message)
  outerr := c.Encode(message)
  if outerr != nil {
    fmt.Println(outerr)
  }
  if fromInterface == true {
    match, value := checkCommand(c, msg)
    if match == true {
      fmt.Println("Positive for command via interface", msg)
      SendChatMessage(c, &irc.Message{
        Params: []string{"#" + channelName},
        Command: "PRIVMSG",
        Trailing: value,
      }, false)
      return value
    }
    return nil
  }
  return nil
}

var commandsMap map[string]interface{};

func setCommands() {
  fmt.Println("Getting commands for '" + channelName + "'")
  commandsMap = commands.GetCommands(channelName)
}

func checkCommand(c *irc.Conn, incoming *irc.Message) (bool, string) {
  for command, value := range commandsMap {
    // fmt.Println(command, incoming.Trailing)
    if sValue, ok := value.(string); ok {
      // fmt.Println(command, getCommand(incoming.Trailing))

      match, _ := regexp.MatchString("^(" + command + ")$", getCommand(incoming.Trailing))
      if(match) {
        return match, sValue
      }
    } else if fValue, ok := value.(func(str string, user string) string); ok {
      // fmt.Println(command, getCommand(incoming.Trailing))

      match, _ := regexp.MatchString("^(" + command + ")$", getCommand(incoming.Trailing))
      if(match) {
        return match, fValue(stripCommand(incoming.Trailing), incoming.User)
      }
    } else {
      return false, "";
    }
  }
  return false, ""
}

// http://stackoverflow.com/a/27160765/4107851
func typeof(v interface{}) string {
  s := fmt.Sprintf("%T", v)
  return s
}

func stripCommand(str string) string {
  // fmt.Println("old str ", str)
  re := regexp.MustCompile("^(!)[a-zA-Z0-9]*(\\s+)?")
  newStr := re.ReplaceAllLiteralString(str, "")
  // fmt.Println("new str ", newStr)
  return newStr
}

func getCommand(str string) string {
  re := regexp.MustCompile("\\s")
  newStr := re.Split(str, 2)
  // fmt.Println(newStr[0])
  return newStr[0]
}
