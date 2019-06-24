package unittest

import (
	"testing"
	"messaging" 
)

func TestMessaging(t * testing.T){
	var testTextMessage = messaging.NewTextMessage()

	if testTextMessage.To != "DISABLEDFORDEMO" {
		t.Errorf("TestMessaging failed at toField%s")
	}

	if testTextMessage.From != "DISABLEDFORDEMO" {
		t.Errorf("TestMessaging failed at fromField%s")
	}

	if testTextMessage.AccountSid != "DISABLEDFORDEMO" {
		t.Errorf("TestMessaging failed at accountSid %s")
	}

	if testTextMessage.AuthToken != "DISABLEDFORDEMO" {
		t.Fatalf("TestMessaging failed at authToken %s")
	}

	if testTextMessage.UrlStr != ("https://api.twilio.com/2010-04-01/Accounts/" + testTextMessage.AccountSid + "/Messages.json"){
		t.Errorf("TestMessaging failed at urlStr %s")
	}

	if testTextMessage.Text != "Hi, Anna! It looks like Alamo Drafthouse has updated their movie showtimes for this weekend." {
		t.Errorf("TestMessaging failed at textField%s")
	}
}
