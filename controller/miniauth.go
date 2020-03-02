package controller

import (
	"context"
	"net/http"
	"os"
	"strings"
	"webframework/models"
	"webframework/utils"

	"github.com/dgrijalva/jwt-go"
)

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
