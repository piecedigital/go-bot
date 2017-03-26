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
  "math/rand"
  "math"
  "strings"
  "github.com/piecedigital/go-bot/bot/quote"
)

const server = "irc.chat.twitch.tv"
const port = "6667"
var botAccount = "piecebot"
var channelName = ""

var retries = 0
var retryLimit = 10

func Connect( send func(ws *websocket.Conn, s string), receive func(conn *irc.Conn, ws *websocket.Conn), ws *websocket.Conn, name string) (*irc.Conn, error) {
  channelName = name
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

var commands = map[string]interface{}{
  // string commands
  "!butt": "Praise booty!",
  "!test": "Yeah, I'm PieceDigital's bot. What of it?",
  "!MMM": "MMMM, B*TCH!",
  "!huge": "That's a HUGE bit!",
  "!yuge": "That's a HUGE bit...ch!",
  "!l": "HOLD this L, fam",
  "!L": "HOLD this L, fam",
  "!forhonor": "... and by honor they really mean ledge shoving and double teaming. Kappa",
  "strangeTry": "Yep. First try. True story Kappa",
  "!oki": "Okizeme (起き攻め) is a term used in fighting games which refers to pressuring the opponent while they are getting up after being knocked down... Okizeme is not actually a Japanese word but was popularized by fighting game enthusiasts. The term is a combination of the Japanese verbs Okiru (起きる) which means to wake-up and Semeru (攻める) which means to attack or strike. The verbs are conjugated and rendaku applies so the unvoiced consonant 's' becomes its voiced form 'a'.",
  "!okizeme": "Okizeme (起き攻め) is a term used in fighting games which refers to pressuring the opponent while they are getting up after being knocked down... Okizeme is not actually a Japanese word but was popularized by fighting game enthusiasts. The term is a combination of the Japanese verbs Okiru (起きる) which means to wake-up and Semeru (攻める) which means to attack or strike. The verbs are conjugated and rendaku applies so the unvoiced consonant 's' becomes its voiced form 'z'.",
  "!backseat": "NO BACKSEAT GAMING! Spawn is a strong, independent, half native American man who don't need no help!",
  "!spoilers": "NO SPOILERS OR BACKSEAT GAMING! That means no telling the streamer what to do, no telling them boss names, no letting them know that a boss is coming up, no telling him what where items are, etc. This is a BLIND and PURE playthrough! Respect it!",
  "!blind": "This is a blind playthrough. That means that the streamer is blind and can't see what they're doing... figuratively, of course. However that is no invitation for you to tell them what to do, how to do it, where to go, etc. Unless they ask, don't tell.",
  "!salt": "The PJSalt salt PJSalt is PJSalt real PJSalt",
  "!tea": "DROP... THE BAGS... OF TEEEEAA!",
  "!whoami": "Oh, me? I'm PieceDigital's bot. A little pet project of his, made with GoLang. He made me primarily to learn a new programming language. That's about it.",
  "!whosmymaker": "PieceDigital made me with GoLang. He's a web developer. He makes web apps. That's about it.",
  "!discord": "https://discord.gg/0gD1dnzC4JZWylTV",
  "!chest": "You just got PUNCHED in yo' chest!",
  "!cash": "CASH ME OU'SIDE! HOW 'BOUT DAH?!",
  "!rules": "Keep it positive! No talking about religion or politics. Keep use of CAPS LOCK to a minimum. Ask permission before posting links. Refrain from backseat gaming, unless otherwise stated. Don't talk about your age; you're 100, far as we're concerned. Try not to cuss too much. NO spoilers!",
  "!muggers": "Muggers... muggers, everywhere...",
  // "!amorrius": "Amorrius is a webapp created by PieceDigital that makes the Twitch core experience much nicer (in his eyes)",
  // "!amorriuslink": "Hopefully I don't get timed out for this...: https://www.amorrius.net"
  // function commands
  "!ftest": func(str string, user string) string {
    return str + " -> " + user;
  },
  "!commands": func(str string, user string) string {
    return str + " -> " + getCommands();
  },
  "!shoutout": func(str string, user string) string {
    return "Shoutout to " + str + ", you amazing strangeling!"
  },
  "!plug": func(str string, user string) string {
    return "<3 Be sure to checkout " + str + " and SERIOUSLY consider dropping a follow! https://www.twitch.tv/" + str + " <3"
  },
  "!hold": func(str string, user string) string {
    return "Hooooold this L, " + str
  },
  "!gentlehold": func(str string, user string) string {
    return str + ", gently grasp this L firmly"
  },
  "!quote": func(str string, user string) string {
    // fmt.Printf("in command: %q\n", str)
    re, err := regexp.Compile("\\s")
    if err != nil {
      return "Problem using !quote"
    }
    if len(str) < 1 {
      return "Usage: !quote [operation] [value]. Available operations: get, who, set"
    }
    newStr := re.Split(str, 2)
    fmt.Println(newStr)
    if len(newStr) < 2 {
      newStr = append(newStr, "")
    }
    fmt.Println(newStr)
    fmt.Println("in command", newStr[1])
    var returnVal string
    returnVal = quote.Quote(newStr[0], newStr[1], user)
    return returnVal
  },
  "!salty": func(str string, user string) string {
    options := []string{str + ", is you salty? Or is you salty?", str + ", is you is or is you ain't super salty?", "Mad or naw, " + str + "?", "PJSalt " + str, "Stay salty, " + str}
    random :=  rand.Intn(len(options))
    // fmt.Println(random)
    index := math.Floor(float64(random))
    // fmt.Println(index)
    return options[int(index)]
  },
  "!donate": func(str string, user string) string {
    options := []string{"Consider donating to this man! I think he deserves a li'l' somethin' :) : https://www.twitchalerts.com/donate/spawnofodd", "Give this man all your money! You don't need it: https://www.twitchalerts.com/donate/spawnofodd", "Just turn your purse upside down: https://www.twitchalerts.com/donate/spawnofodd", "At the very least put him in your will?: https://www.twitchalerts.com/donate/spawnofodd"}
    random :=  rand.Intn(len(options))
    // fmt.Println(random)
    index := math.Floor(float64(random))
    // fmt.Println(index)
    return options[int(index)]
  },
  "!weeb": weebs,
  "!weebs": weebs,
}

func weebs(str string, user string) string {
  options := []string{"WEEEEEEEEBS! WEEEEEEEEEEEEEEEEEEEEEEEEEEBS!", "WEEEEEBS! FREAKIN' WEEEEEEEEEEEEBS!", "Gosh, so many damn WEEEEEEEEEEBS!"}
  random :=  rand.Intn(len(options))
  // fmt.Println(random)
  index := math.Floor(float64(random))
  // fmt.Println(index)
  return options[int(index)]
}

var copy map[string]interface{};

func init()  {
  copy = commands
}

func getCommands() string {
  var list = []string{};

  for command, _ := range copy {
    list = append(list, command)
  }

  return "Commands: " + strings.Join(list, ", ")
}

func checkCommand(c *irc.Conn, incoming *irc.Message) (bool, string) {
  for command, value := range commands {
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
