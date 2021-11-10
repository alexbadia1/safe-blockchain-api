package main

import "net/http"

// Serves image
func bobbyImage(w http.ResponseWriter, r *http.Request) {
	// This is bad, I know, it's for demo purposes only...
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, "images/bobby-headshot.png")
} // getImage
