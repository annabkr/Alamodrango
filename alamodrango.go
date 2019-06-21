package main

import (
	"webscrape"   
	"drafthouse"
	"http"
)

func main() {   
	http.HTTPHandler(webscrape.Url) 

	var schedule drafthouse.Schedule
	drafthouse.SetSchedule(&schedule) 
	drafthouse.MonitorForUpdates(schedule)
}
