package messaging

import (
  "fmt"
  "strings"
  "net/http"
  "net/url"
  "encoding/json"
  "webscrape"
  "unittest"
)

func messanger() { 
  accountSid := "ACd93ff7e1ea248a23e77e273f0e0a9315"
  authToken := "c2659fe77318e206b34fa5485d1dc062"
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
 
  var quote string = "Hi, Anna! It looks like Alamo Drafthouse has updated their movie showtimes for this weekend."
  
  msgData := url.Values{}
  msgData.Set("To","14158280722")
  msgData.Set("From","14157669491")
  msgData.Set("Body", quote)
  msgDataReader := *strings.NewReader(msgData.Encode())

  //HTTP request client
  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  //HTTP POST request, return message SID
  resp, _ := client.Do(req)
  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode(&data)
    if (err == nil) {
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }

}