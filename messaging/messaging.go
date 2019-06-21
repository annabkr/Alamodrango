package messaging

import ( 
  "strings" 
  "http"
  "net/url" 
)

type TextMessage struct{
  To string
  From string
  AccountSid string
  AuthToken string
  UrlStr string
  Text string 
}

func RunMessenger() { 
  var msg TextMessage = NewTextMessage()
  
  msgData := url.Values{}
  msgData.Set("To", msg.To)
  msgData.Set("From", msg.From)
  msgData.Set("Body", msg.Text)
  msgDataReader := *strings.NewReader(msgData.Encode())

  client, req := http.RequestClient(&msgDataReader, msg.UrlStr, msg.AccountSid, msg.AuthToken)
  http.PostRequest(client, req) 
} 


func NewTextMessage() TextMessage {
  sid := "ACd93ff7e1ea248a23e77e273f0e0a9315"
  auth := "c2659fe77318e206b34fa5485d1dc062"
  url := "https://api.twilio.com/2010-04-01/Accounts/" + sid + "/Messages.json"
  toField := "14158280722"
  fromField := "14157669491"
  quote := "Hi, Anna! It looks like Alamo Drafthouse has updated their movie showtimes for this weekend."

  return TextMessage{To: toField, From: fromField, AccountSid: sid, AuthToken: auth, UrlStr: url, Text: quote}
}