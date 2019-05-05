package drafthouse 

type CinemaWebsite struct {
	Calendar 		 Calendar
}

type Cinema struct {
	CinemaId          string
	CinemaName        string
	CinemaTimeZoneATE string   
	Films             []Film
	Months			  []Month
}


