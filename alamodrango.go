package main

import (
	"log" 
	"fmt"
	"webscrape"
	"github.com/robfig/cron" 
	"messaging"
	"http"
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

var schedule Schedule
var mainSite = &webscrape.MainSite

func main() {  
	http.HTTPHandler(webscrape.Url) 
	setSchedule(&schedule) 
	monitorForUpdates()
}

func printFilmSessions(mainSite webscrape.CinemaWebsite){
	for i := 0; i < len(schedule.Dates); i++{
		log.Println("Day: ", i)
		for j := 0; j < len(schedule.Dates[i].Films); j++ {
			fmt.Printf(" Film Name:%s, Showtimes:%s \n ", schedule.Dates[i].Films[j].FilmName, schedule.Dates[i].Films[j].Showtimes)
		}
	}
}

func setSchedule(schedule *Schedule){
	log.Println("Set Schedule")
	if len(mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days) != 0{
		for i := 0; i < len(mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days); i++ {
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
}

func monitorForUpdates(){
	c := cron.New()
	c.AddFunc("@hourly", func() {
		log.Println("Cron open") 
		var newSched Schedule
		http.GetData(webscrape.Url)
		setSchedule(&newSched)

		if len(newSched.Dates) > 0{
			for i := 5; i < 7; i++ {
				if len(newSched.Dates[i].Films) > len(schedule.Dates[i].Films){ 
					messaging.RunMessenger()
				}
			}
		}
		 
	})
	c.Start() 
}

