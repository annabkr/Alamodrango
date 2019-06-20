package main

import (
	"encoding/json"  
	"log" 
	"fmt"
	"io/ioutil"
	"net/http"  
	"webscrape"
	"unittest"
	"messaging"  
)

type Schedule struct {
	Dates [7]Days
}

type Days struct { 
	Films []Showing
}

type Showing struct {
	FilmName string
	Showtimes []string
}


func main() {

	response, err := http.Get(drafthouse.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	pageBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("error")
	} 

	var mainSite drafthouse.CinemaWebsite
	e := json.Unmarshal(pageBytes, &mainSite)
	if e != nil {
		log.Println("error")
	}

	var schedule Schedule
	getFilmSessions(mainSite, &schedule)

	for i := 0; i < len(schedule.Dates); i++{
		log.Println("Day: ", i)
		for j := 0; j < len(schedule.Dates[i].Films); j++ {
			fmt.Printf(" Film Name:%s, Showtimes:%s \n ", schedule.Dates[i].Films[j].FilmName, schedule.Dates[i].Films[j].Showtimes)
		}
	}
	 
}



func getFilmSessions(mainSite drafthouse.CinemaWebsite, schedule *Schedule){
	for i := 0; i < len(schedule.Dates); i++ {
	
		var dailyFilms []Showing

		for j := 0; j < len(mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films); j++ {

			var filmName string = mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films[j].FilmName
			var showtimes []string

			for k := 0; k < len(mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films[j].Series); k++ {
				for l := 0; l < len(mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films[j].Series[k].Formats); l++ {
					for m := 0; m < len(mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films[j].Series[k].Formats[l].Sessions); m++ {
						showtimes = append(showtimes, mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films[j].Series[k].Formats[l].Sessions[m].SessionTime)
					}
				}
			}
			
			var film = Showing {filmName, showtimes} 
			dailyFilms = append(dailyFilms, film)
		}

		schedule.Dates[i].Films = append(dailyFilms)
	} 

}

