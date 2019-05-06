//Package main 
package main

import (
	"encoding/json"  
	"log"
	"fmt"
	"io/ioutil"
	"net/http"  
	"drafthouse" 
)

type SimpleFilm struct {
	FilmTitle string
	Showtimes []Showtime
}

type Schedule struct {
	Films []SimpleFilm 
}

type Showtime struct {  
	Date 	string
	Time  	string
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

//	var s Schedule 
//	setSchedule(mainSite, s)

	friday := isItFriday(mainSite)

	fmt.Println(friday)


}

func isItFriday(mainSite drafthouse.CinemaWebsite) (bool) {
	return (mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[5].DayOfWeekNumber == "6")
}

func setSchedule(mainSite drafthouse.CinemaWebsite, s Schedule) {
for i :=  0; i < drafthouse.DaysOfTheWeek; i++ { 
	for j := 0; j < len(mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films); j++{
		s.Films = append (mainSite.Calendar.Cinemas[0].Months[0].Weeks[0].Days[i].Films[j], j)
		fmt.Println(s.Films)
		}
	}
}

func setShowtimes(mainSite drafthouse.CinemaWebsite) {

//	for _, num := range mainSite.
}

func setHasUpdated(mainSite drafthouse.CinemaWebsite) {

}