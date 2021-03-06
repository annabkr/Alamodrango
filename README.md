# go-alamodrango

Web scraper bot that sends text messages when Alamo Drafthouse posts new movie showtimes.

I'm an Alamo Drafthouse regular. They update the showtimes manually, sometime between Monday-Wednesday. I got tired of refreshing the page every hour, so I decided to make a bot to do it for me. 

Gotta get those good seats! Continuing to expand functionality and support.

<p align="center"><img width="502" alt="Screen Shot 2019-06-19 at 5 45 34 PM" src="https://user-images.githubusercontent.com/44475953/59811201-a66ac480-92bd-11e9-967d-c54b5d8c6b4e.png"></p> 
 
## Getting Started
These instructions will get a copy of the project up and running on your local machine.

Note that some features, including Twilio functionality, are disabled for the demo of this project for security purposes.

## Installing

Download or clone the repository. Ensure that you have Golang installed. In the console, navigate to the project directory. 

Build using the following command:
```
go build alamodrango.go
```
Run with:
```
go run alamodrango.go
```

## Running the Tests

```
go test unittest
```

## Built With

* [Twilio's REST API](https://www.twilio.com/) - Text messaging
* [Cron](https://github.com/robfig/cron) - Automation 
