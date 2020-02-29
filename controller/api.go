package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"webframework/models"
	"webframework/utils"

	"github.com/dgrijalva/jwt-go"
)

// Serve will serve the frontend
var Serve = func(w http.ResponseWriter, r *http.Request) {
	var prod string = os.Getenv("is_production")
	// Deal with the authentication first
	authenticated := utils.Auth(w, r)
	if !authenticated {
		return
	}
	var staticPath string
	// Change static path based on production or not
	if strings.ToLower(prod) == "true" {
		staticPath = "./build"
	} else {
		staticPath = "./website/build"
	}
	const indexPath = "index.html"
	fileServer := http.FileServer(http.Dir(staticPath))
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(staticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileServer.ServeHTTP(w, r)
}

// CreateAccount : Creates an account on the database
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Attempt to create an account detected")
	// Parse the incoming payload
	// The account has to follow this format
	// {
	// 		email:"example@acopointer.sg",
	// 		password:"password",
	// }
	// Create an account struct to hold the data
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		fmt.Println("Failed to create an account")
		// Handle a generic error
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}
	// Create the account
	message, ok := account.Create()
	if !ok {
		fmt.Println(message)
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, message)
		return
	}
	w.WriteHeader(http.StatusOK)
	addCookie(w, message["account"].(*models.Account).Token)
	utils.Respond(w, message)
}

// ChangePassword : Changes the password of the account, provided that the user knows the old password
var ChangePassword = func(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming payload
	// The account has to follow this format
	// {
	// 		email:"example@acopointer.sg",
	// 		username:"username"
	// 		password:"password",
	// 		newpassword:"password2",
	// }
	// Declare a temporary NewAccount struct
	type newAccount struct {
		Email       string `json:"email"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Newpassword string `json:"newpassword"`
	}
	temp := &newAccount{}
	// Decode the payload into temp
	err := json.NewDecoder(r.Body).Decode(temp)
	if err != nil {
		// Handle a generic error
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}
	if temp.Newpassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(false, "You have to provide a new password"))
		return
	}
	// Convert the temp object to an account object
	acc := &models.Account{
		Email:    temp.Email,
		UserName: temp.Username,
		Password: temp.Password,
	}
	message, ok := acc.ChangePassword(temp.Newpassword)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, message)
		return
	}
	w.WriteHeader(http.StatusOK)
	addCookie(w, message["account"].(*models.Account).Token)
	utils.Respond(w, message)
}

//Login : Login to the main page
var Login = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login attempt detected")
	// Parse the incoming payload
	// The account has to follow this format
	// {
	// 		email:"example@acopointer.sg",
	// 		username:"username"
	// 		password:"password",
	// }
	// Create an account struct to hold the data
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		// Handle a generic error
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}
	message, ok := account.Login()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(false, "Invalid Email/Password provided"))
		return
	}
	w.WriteHeader(http.StatusOK)
	addCookie(w, message["account"].(*models.Account).Token)
	utils.Respond(w, message)
}

func addCookie(w http.ResponseWriter, jwString string) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   jwString,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func miniAuth(writer http.ResponseWriter, request *http.Request) {
	tokenHeader := request.Header.Get("Authorization")
	response := make(map[string]interface{})
	if tokenHeader == "" {
		response = utils.Message(false, "Missing auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		utils.Respond(writer, response)
		return
	}
	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		response = utils.Message(false, "Invalid/Malformed auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		utils.Respond(writer, response)
		return
	}
	tokenPart := splitted[1] // the information that we're interested in
	tk := &models.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	//malformed token, return 403
	if err != nil {
		response = utils.Message(false, "Malformed auth token")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		utils.Respond(writer, response)
		return
	}
	//token is invalid
	if !token.Valid {
		response = utils.Message(false, "Token is invalid")
		writer.WriteHeader(http.StatusForbidden)
		writer.Header().Add("Content-Type", "application/json")
		utils.Respond(writer, response)
		return
	}
	ctx := context.WithValue(request.Context(), "user", tk.UserID)
	request = request.WithContext(ctx)
}
