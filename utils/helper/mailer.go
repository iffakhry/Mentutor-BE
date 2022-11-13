package helper

import (
	"be12/mentutor/features/mentee"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// Get link for auth code
func GetUrl() (string, error){
	// Reads in our credentials
	secret, err := ioutil.ReadFile("client_secret_web.json")
	if err != nil {
		log.Printf("Error: %v", err)
		return "", err
	}

	// Creates a oauth2.Config using the secret
	// The second parameter is the scope, in this case we only want to send email
	conf, err := google.ConfigFromJSON(secret, gmail.GmailSendScope)
	if err != nil {
		log.Printf("Error: %v", err)
		return "", err
	}

	// Creates a URL for the user to follow
	url := conf.AuthCodeURL("CSRF", oauth2.AccessTypeOffline)
	// Prints the URL to the terminal
	return url, nil
}

// Get token from access code
func GetToken(code string) (*oauth2.Token, error) {

	// Reads in our credentials
	secret, err := ioutil.ReadFile("client_secret_web.json")
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Creates a oauth2.Config using the secret
	// The second parameter is the scope, in this case we only want to send email
	conf, err := google.ConfigFromJSON(secret, gmail.GmailSettingsBasicScope)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Exchange the auth code for an access token
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	staticTok := conf.TokenSource(oauth2.NoContext, tok)
	token , err := staticTok.Token()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return token, nil
}

// Send gmail message
func BuildMessage(tok *oauth2.Token, mentee mentee.MenteeCore, mentor mentee.MentorCore, task mentee.Task) error {

	// Reads in our credentials
	secret, err := ioutil.ReadFile("client_secret_web.json")
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Creates a oauth2.Config using the secret
	// The second parameter is the scope, in this case we only want to send email
	conf, err := google.ConfigFromJSON(secret, gmail.MailGoogleComScope)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Create the *http.Client using the access token
	client := conf.Client(oauth2.NoContext, tok)

	// Create a new gmail service using the client
	gmailService, err := gmail.New(client)
	if err != nil {
		log.Printf("Eror: %v", err)
	}

	// New message for our gmail service to send
	var message gmail.Message

	from := "From: " + mentee.Email + "\r\n"
	to := "To: " + mentor.Email + "\r\n"
	subject := "Subject: [ASSIGNMENT] " + mentee.Name + " telah mengerjakan tugas " + task.Title  + "\r\n\r\n"
	body := "Dear " + mentor.Name + "\r\n" + mentee.Name + " telah submit tugas, mohon untuk periksa tugas dan berikan penialaian. \r\n Terima kasih" 
	// Compose the message
	messageStr := []byte(
		from +
			to +
			subject +
			body)

	// Place messageStr into message.Raw in base64 encoded format
	message.Raw = base64.URLEncoding.EncodeToString(messageStr)

	// Send the message
	_, err = gmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	} else {
		fmt.Println("Message sent!")
		return nil
	}
}