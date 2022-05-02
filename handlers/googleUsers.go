package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lazyspell/Ecommerce_Backend/config"
	"github.com/lazyspell/Ecommerce_Backend/utils"
)

func (m *Repository) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := utils.GenerateStateOauthCookie(w)

	u := config.Config.GoogleLoginConfig.AuthCodeURL(oauthState)

	http.Redirect(w, r, u, http.StatusTemporaryRedirect)

}

func (m *Repository) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// get oauth state from cookie for this user
	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")

	w.Header().Add("content-type", "application/json")

	// ERROR : Invalid OAuth State
	if state != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "invalid oauth google state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := config.Config.GoogleLoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(w, "failed code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from google server
	response, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		fmt.Fprintf(w, "failed getting user info: %s", err.Error())
		return
	}

	// Parse user data JSON Object
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, "failed read response: %s", err.Error())
		return
	}

	// // send back response to browser
	fmt.Fprintln(w, string(contents))

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contents)

	type GoogleObject struct {
		ID        string
		Email     string
		Verified  bool
		Name      string
		GivenName string
		Picture   string
		Locale    string
	}

	var googleObject GoogleObject

	if err := json.Unmarshal(contents, &googleObject); err != nil {
		log.Println(err)
	}

}
