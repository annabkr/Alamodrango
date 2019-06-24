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
  sid := "DISABLEDFORDEMO"
  auth := "DISABLEDFORDEMO"
  url := "https://api.twilio.com/2010-04-01/Accounts/" + sid + "/Messages.json"
  toField := "DISABLED"
  fromField := "DISABLEDFORDEMO"
  quote := "Hi, Anna! It looks like Alamo Drafthouse has updated their movie showtimes for this weekend."

  return TextMessage{To: toField, From: fromField, AccountSid: sid, AuthToken: auth, UrlStr: url, Text: quote}
}