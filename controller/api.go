package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"webframework/utils"
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
