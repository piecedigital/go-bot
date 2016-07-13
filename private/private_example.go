package private

type Private struct {
  client, secret, oauth string
}

var tokens *Private = &Private{
  "<Twitch Client Key>", // client
  "<Twitch Secret Key>", // secret
  "<Auth Token>", // oauth : http://twitchapps.com/tmi/
}

func GetClient() string {
  t := tokens.client
  return t
}

func GetSecret() string {
  t := tokens.secret
  return t
}

func GetAuthToken() string {
  t := tokens.oauth
  return t
}
