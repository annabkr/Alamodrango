package unittest

import (
	"testing"
	"messaging" 
)

func TestMessaging(t * testing.T){
	var testTextMessage = messaging.NewTextMessage()

	if testTextMessage.To != "14158280722" {
		t.Errorf("TestMessaging failed at toField%s")
	}

	if testTextMessage.From != "14157669491" {
		t.Errorf("TestMessaging failed at fromField%s")
	}

	if testTextMessage.AccountSid != "ACd93ff7e1ea248a23e77e273f0e0a9315" {
		t.Errorf("TestMessaging failed at accountSid %s")
	}

	if testTextMessage.AuthToken != "c2659fe77318e206b34fa5485d1dc062" {
		t.Fatalf("TestMessaging failed at authToken %s")
	}

	if testTextMessage.UrlStr != ("https://api.twilio.com/2010-04-01/Accounts/" + testTextMessage.AccountSid + "/Messages.json"){
		t.Errorf("TestMessaging failed at urlStr %s")
	}

	if testTextMessage.Text != "Hi, Anna! It looks like Alamo Drafthouse has updated their movie showtimes for this weekend." {
		t.Errorf("TestMessaging failed at textField%s")
	}
}
