package http

import (
	"io"
	"io/ioutil"
	"net/http"  
	"webscrape" 
	"log"
	"encoding/json"
	"fmt"  
)

func HTTPHandler(url string){
	response := GetData(url)
	pageBytes := ReadData(response)
	DecodeData(pageBytes, &webscrape.MainSite)
}

func GetData(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		log.Println("GetData error")
	}
	return response
}

func ReadData(response *http.Response) [] byte{
	pageBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ReadData error")
	} 

	response.Body.Close()

	return pageBytes 
}

func DecodeData(pageBytes [] byte, site *webscrape.CinemaWebsite){ 
	e := json.Unmarshal(pageBytes, site) 
	if e != nil {
		log.Println("DecodeData error")
	}
}

func RequestClient(msgDataReader io.Reader, urlStr string, accountSid string, authToken string) (*http.Client, *http.Request){ 
  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, msgDataReader)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  return client, req
}

func PostRequest(client *http.Client, req *http.Request){ 
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