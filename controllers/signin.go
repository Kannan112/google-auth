package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var ssogolang *oauth2.Config
var RandomString = "random-string"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Clien := os.Getenv("CLIENT_ID")
	ClientSrc := os.Getenv("CLIENT_SECRET")
	RedirUrl := os.Getenv("REDIRECT_URL")

	ssogolang = &oauth2.Config{
		RedirectURL:  RedirUrl,
		ClientID:     Clien,
		ClientSecret: ClientSrc,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	url := ssogolang.AuthCodeURL(RandomString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
