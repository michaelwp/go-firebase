package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
)

/**
Firebase connection
by MPutong, 23082020
 */

func main() {
	opt := option.WithCredentialsFile("email-verification-key-firebase-adminsdk-vqikz-489f25dd4e.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	errHandler("Error firebase new app", err)

	client, err := app.Auth(context.Background())
	errHandler("Error firebase auth", err)

	params := (&auth.UserToCreate{}).
		Email("michael.wenceslaus@gmail.com").
		EmailVerified(false).
		PhoneNumber("+6287778313111").
		Password("secretPassword").
		DisplayName("Michael Putong").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.CreateUser(context.Background(), params)
	errHandler("error create user", err)

	s, err := client.EmailVerificationLink(context.Background(), "michael.wenceslaus@gmail.com")
	errHandler("error send sign in link", err)

	log.Printf("send email verification: %s", s)
	log.Printf("Successfully created user: %v\n", u)
}

func errHandler(msg string, err error){
	if err != nil { log.Fatalf("%s: %s", msg, err)}
}
