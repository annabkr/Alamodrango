package webscrape

type Film struct{
	FilmId string
	FilmName string  
	FilmYear string 
	FilmRating string 
	FilmRuntime string 
	FilmAgePolicy string 
	FilmSlug string 
	Series []Series 
}