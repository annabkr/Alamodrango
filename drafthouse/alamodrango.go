// alamogrango.go
package main

import (
	"encoding/json"  
	"sort"
	"log" 
	"os"
	"net/http" 
	"github.com/annabakr/Alamodrango/drafthouse"
)

func main() {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	in := json.NewDecoder(response.Body)
	out := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := in.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		var keys []string
		for k := range v {
			if k != "FilmName" {
				delete(v, k) 
			} else {
				keys = append(keys, k)
			}

			sort.Strings(keys)
			for _, k := range keys {
				log.Println("Key:", k, "Value: ", v[k])
			}
		}

		if err := out.Encode(&v); err != nil {
			log.Println(err)
		}

		log.Println(out)
	}

	//var mainSite CinemaWebsite
	//e := json.Unmarshal(pageBytes, &mainSite)
	//if e != nil {
	//	log.Println(e)
	//}

	//var numFilms := mainSite.Calendar.Cinemas[1].Films


	//for i = 0; i < 

	//log.Println("Day's films: \n", mainSite.Calendar.Cinemas.size())

	}