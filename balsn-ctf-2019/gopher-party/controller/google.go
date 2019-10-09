package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"balsnctf/gopherparty/event"
	"balsnctf/gopherparty/model"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("GOOGLE_OAUTH_CALLBACK_URL"),
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// OauthGoogleLogin is used for google oauth login.
func OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	if event.IsDown() {
		gogopher(w, "event has not started yet")
		return
	}

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

// OauthGoogleCallback controller the func for google callback.
func OauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	if event.IsDown() {
		gogopher(w, "event has not started yet")
		return
	}
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := getAccessToken(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	user, err := getUserDataFromGoogle(token.AccessToken)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	t, err := template.ParseFiles("template/register_form.html")
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	context := make(map[string]string)
	context["AccessToken"] = token.AccessToken
	context["Name"] = user.Name
	context["Email"] = user.Email
	context["Picture"] = user.Picture
	t.Execute(w, context)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getAccessToken(code string) (*oauth2.Token, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	return token, nil
}

func getUserDataFromGoogle(t string) (model.User, error) {
	response, err := http.Get(oauthGoogleURLAPI + t)
	if err != nil {
		log.Printf("failed getting user info: %s", err.Error())
		return model.User{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed read response: %s", err.Error())
		return model.User{}, fmt.Errorf("failed read response: %s", err.Error())
	}
	user := model.User{}
	json.Unmarshal([]byte(contents), &user)
	if err != nil {
		log.Printf("json unmarshal failed: %s", err.Error())
		return model.User{}, fmt.Errorf("json unmarshal failed: %s", err.Error())
	}
	if user.Error.Code == 401 {
		return user, fmt.Errorf("json unmarshal failed: %s", user.Error.Message)
	}
	return user, nil
}

func getUser(t string) (model.User, error) {
	return getUserDataFromGoogle(t)
}

func getUserPicture(t string) (string, error) {
	user, err := getUserDataFromGoogle(t)
	if err != nil {
		return "", err
	}
	return user.Picture, nil
}

func getUserLocale(t string) (string, error) {
	user, err := getUserDataFromGoogle(t)
	if err != nil {
		return "", err
	}
	return user.Locale, nil
}
