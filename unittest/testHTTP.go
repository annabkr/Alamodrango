package unittest

import (
	"testing"
	"webscrape"
	"http" 
)

func TestHTTP(t *testing.T){
	var testSite webscrape.CinemaWebsite
	response := http.GetData(webscrape.Url)
	if response == nil {
		t.Fatalf("TestHTTP: http.GetData() failed: %s")
	} 

	pageBytes := http.ReadData(response)
	if pageBytes == nil {
		t.Fatalf("TestHTTP: http.ReadData error %s")
	}

	http.DecodeData(pageBytes, &testSite)

}  
