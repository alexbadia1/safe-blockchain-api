package main

import "net/http"

// Serves image
func getImage(w http.ResponseWriter, r *http.Request) {
	// This is bad, I know, it's for demo purposes only...
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Make sure this endpoint is only accessible at "/image"
	if r.URL.Path != "/image" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	if name := r.URL.Query().Get("name"); name == "bobby" {
		http.ServeFile(w, r, "images/bobby-headshot.png")
	} else {
		http.Error(w, "No Image of "+name+" exists!", http.StatusNoContent)
	} // else
} // getImage
