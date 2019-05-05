// alamodrango.go

// Add support for: https://feeds.drafthouse.com/adcService/showtimes.svc/calendar/0801/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
func main() {

	// Create and modify HTTP request
	response, err := http.Get("https://feeds.drafthouse.com/adcService/showtimes.svc/calendar/0801/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	
	//Get the body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	sessionDateTime := strings.Index(pageContent, "SessionDateTime")
	if sessionDateTime == -1 {
		fmt.Println("No sessions found")
		os.Exit(0)
	}

	fmt.Printf("Session: %s\n", sessionDateTime)

}