package drafthouse

type Calendar struct {
	FeedGenerated string
	SessionsGenerated string
	FilterLists FilterList 
	Cinemas []Cinema
}

type Month struct {
	MonthDate string 
	Weeks []Week
}

type Week struct {
	WeekNumber string 
	Days []Day
}

type Day struct {
	DayOfWeekNumber string
	DateId string
	Date string
	DayType string
	Films []Film	
}
