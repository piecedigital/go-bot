package commands

import (
  "fmt"
  "strings"
  "regexp"
  "math"
  "math/rand"
  "github.com/piecedigital/go-bot/bot/quote"
)

func GetCommands(channel string) map[string]interface{} {
  fmt.Printf("returning commands for channel %v\n", channel)

  var copy map[string]interface{};

  switch strings.ToLower(channel) {
    case "spawnofodd":
      copy = map[string]interface{}{
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
        "!lgbt": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!lgbt+": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!lgbtq": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!lgbtq+": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!gsd": "This streamer is YUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!sg": "When's Skullgirls?",
        "!skull": "When's Skullgirls?",
        "!skullgirls": "When's Skullgirls?",
        "!japan": "*takes a deep breath* *points with both hands flat* ...Japan...",
        // "!amorrius": "Amorrius is a webapp created by PieceDigital that makes the Twitch core experience much nicer (in his eyes)",
        // "!amorriuslink": "Hopefully I don't get timed out for this...: https://www.amorrius.net"
        // function commands
        "!ftest": func(str string, user string) string {
          return str + " -> " + user;
        },
        "!commands": func(str string, user string) string {
          return str + " -> " + getCommandsList(copy);
        },
        "!shoutout": func(str string, user string) string {
          return "Shoutout to " + str + ", you amazing strangeling!"
        },
        "!plug": func(str string, user string) string {
          fmt.Println("old str ", str)
          re1, err1 := regexp.Compile("\\s")
          if err1 != nil {
            return "Problem using !plug"
          }
          if len(str) < 1 {
            return "Usage: !plug [username]"
          }
          newStr := re1.Split(str, 2)

          re2, err2 := regexp.Compile("^(@)")
          if err2 != nil {
            return "Problem using !plug"
          }
          finalStr := re2.ReplaceAllLiteralString(newStr[0], "")

          fmt.Println("final str ", finalStr)

          return "<3 Be sure to checkout " + finalStr + " and SERIOUSLY consider dropping a follow! https://www.twitch.tv/" + finalStr + " <3"
        },
        "!hold": func(str string, user string) string {
          return "Hooooold this L, " + str
        },
        "!gentlehold": func(str string, user string) string {
          return str + ", gently grasp this L firmly"
        },
        "!welcome": func(str string, user string) string {
          return str + ", welcome... to the strange!"
        },
        "!hug": func(str string, user string) string {
          return user + " gives " + str + " a hug. <3"
        },
        "!boo": func(str string, user string) string {
          return "Boo this man! BOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO! " + str
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
    case "piecedigital":
      copy = map[string]interface{}{
        // string commands
        "!butt": "Praise booty!",
        "!test": "Yeah, I'm PieceDigital's bot. What of it?",
        "!MMM": "MMMM, B*TCH!",
        "!l": "HOLD this L, fam",
        "!L": "HOLD this L, fam",
        "!forhonor": "... and by honor they really mean ledge shoving and double teaming. Kappa",
        "strangeTry": "Yep. First try. True story Kappa",
        "!oki": "Okizeme (起き攻め) is a term used in fighting games which refers to pressuring the opponent while they are getting up after being knocked down... Okizeme is not actually a Japanese word but was popularized by fighting game enthusiasts. The term is a combination of the Japanese verbs Okiru (起きる) which means to wake-up and Semeru (攻める) which means to attack or strike. The verbs are conjugated and rendaku applies so the unvoiced consonant 's' becomes its voiced form 'a'.",
        "!okizeme": "Okizeme (起き攻め) is a term used in fighting games which refers to pressuring the opponent while they are getting up after being knocked down... Okizeme is not actually a Japanese word but was popularized by fighting game enthusiasts. The term is a combination of the Japanese verbs Okiru (起きる) which means to wake-up and Semeru (攻める) which means to attack or strike. The verbs are conjugated and rendaku applies so the unvoiced consonant 's' becomes its voiced form 'z'.",
        // "!backseat": "NO BACKSEAT GAMING! Spawn is a strong, independent, half native American man who don't need no help!",
        // "!spoilers": "NO SPOILERS OR BACKSEAT GAMING! That means no telling the streamer what to do, no telling them boss names, no letting them know that a boss is coming up, no telling him what where items are, etc. This is a BLIND and PURE playthrough! Respect it!",
        // "!blind": "This is a blind playthrough. That means that the streamer is blind and can't see what they're doing... figuratively, of course. However that is no invitation for you to tell them what to do, how to do it, where to go, etc. Unless they ask, don't tell.",
        // "!salt": "The PJSalt salt PJSalt is PJSalt real PJSalt",
        "!tea": "DROP... THE BAGS... OF TEEEEAA!",
        "!whoami": "Oh, me? I'm PieceDigital's bot. A little pet project of his, made with GoLang. He made me primarily to learn a new programming language. That's about it.",
        "!whosmymaker": "PieceDigital made me with GoLang. He's a web developer. He makes web apps. That's about it.",
        "!discord": "Ain't got one",
        "!chest": "You just got PUNCHED in yo' chest!",
        "!cash": "CASH ME OU'SIDE! HOW 'BOUT DAH?!",
        // "!rules": "Keep it positive! No talking about religion or politics. Keep use of CAPS LOCK to a minimum. Ask permission before posting links. Refrain from backseat gaming, unless otherwise stated. Don't talk about your age; you're 100, far as we're concerned. Try not to cuss too much. NO spoilers!",
        "!muggers": "Muggers... muggers, everywhere...",
        "!lgbt": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!lgbt+": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!lgbtq": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!lgbtq+": "This streamer is HUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!gsd": "This streamer is YUGE supporter of the LGBT+ community. So homophobia, transphobia, disrespecting gender pronouns, etc., will not be tolerated. You will get timed out or banned, or a stern talking to if you're lucky.",
        "!sg": "When's Skullgirls?",
        "!skull": "When's Skullgirls?",
        "!skullgirls": "When's Skullgirls?",
        "!japan": "*takes a deep breath* *points with both hands flat* ...Japan...",
        "!amorrius": "Amorrius is a webapp created by PieceDigital that makes the Twitch core experience much nicer (in his eyes) https://www.amorrius.net",
        "!twitter": "https://twitter.com/piecedigital",
        // function commands
        "!ftest": func(str string, user string) string {
          return str + " -> " + user;
        },
        "!commands": func(str string, user string) string {
          return str + " -> " + getCommandsList(copy);
        },
        // "!shoutout": func(str string, user string) string {
        //   return "Shoutout to " + str + ", you amazing strangeling!"
        // },
        "!plug": func(str string, user string) string {
          fmt.Println("old str ", str)
          re1, err1 := regexp.Compile("\\s")
          if err1 != nil {
            return "Problem using !plug"
          }
          if len(str) < 1 {
            return "Usage: !plug [username]"
          }
          newStr := re1.Split(str, 2)

          re2, err2 := regexp.Compile("^(@)")
          if err2 != nil {
            return "Problem using !plug"
          }
          finalStr := re2.ReplaceAllLiteralString(newStr[0], "")

          fmt.Println("final str ", finalStr)

          return "<3 Be sure to checkout " + finalStr + " and SERIOUSLY consider dropping a follow! https://www.amorrius.net/profile/" + finalStr + " <3"
        },
        "!hold": func(str string, user string) string {
          return "Hooooold this L, " + str
        },
        "!gentlehold": func(str string, user string) string {
          return str + ", gently grasp this L firmly"
        },
        "!welcome": func(str string, user string) string {
          return str + ", welcome... to the strange!"
        },
        "!hug": func(str string, user string) string {
          return user + " gives " + str + " a hug. <3"
        },
        "!boo": func(str string, user string) string {
          return "Boo this man! BOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO! " + str
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
          options := []string{"Consider donating to this man! I think he deserves a li'l' somethin' :) : https://streamtip.com/t/piecedigital", "Give this man all your money! You don't need it: https://streamtip.com/t/piecedigital", "Just turn your purse upside down: https://streamtip.com/t/piecedigital", "At the very least put him in your will?: https://streamtip.com/t/piecedigital"}
          random :=  rand.Intn(len(options))
          // fmt.Println(random)
          index := math.Floor(float64(random))
          // fmt.Println(index)
          return options[int(index)]
        },
        "!weeb": weebs,
        "!weebs": weebs,
      }
    default:
      copy = map[string]interface{}{
        "!test": "couldn't get commands",
      }
  }

  return copy
}

func getCommandsList(cmds map[string]interface{} ) string {
  var list = []string{};

  for command, _ := range cmds {
    list = append(list, command)
  }

  return "Commands: " + strings.Join(list, ", ")
}

func weebs(str string, user string) string {
  options := []string{"WEEEEEEEEBS! WEEEEEEEEEEEEEEEEEEEEEEEEEEBS!", "WEEEEEBS! FREAKIN' WEEEEEEEEEEEEBS!", "Gosh, so many damn WEEEEEEEEEEBS!"}
  random :=  rand.Intn(len(options))
  // fmt.Println(random)
  index := math.Floor(float64(random))
  // fmt.Println(index)
  return options[int(index)]
}


// func init()  {
//   copy = commands
// }
