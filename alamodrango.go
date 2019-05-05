// alamogrango.go
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

	response, err := http.Get("https://feeds.drafthouse.com/adcService/showtimes.svc/calendar/0801/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	pageBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(pageBytes)

	startIndex := strings.Index(pageContent, "SessionDateTime")
	if startIndex == -1 {
		fmt.Println("No element found")
		os.Exit(0)
	}

	startIndex += 18

	endIndex := strings.Index(pageContent, "SessionType")
	if endIndex == -1 {
		fmt.Println("No closing tag for element found")
		os.Exit(0)
	}

	endIndex -= 3


	sessionDateTime := []byte(pageContent[startIndex:endIndex])

	fmt.Printf("SessionDateTime: %s\n", sessionDateTime)

	}