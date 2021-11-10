package main

import "net/http"

// Serves image
func getImage(w http.ResponseWriter, r *http.Request) {
	// This is bad, I know, it's for demo purposes only...
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/png")

	if name := r.URL.Query().Get("name"); name == "bobby" {
		http.ServeFile(w, r, "./images/bobby-headshot.png")
	} else {
		http.Error(w, "Image doesn't exist of exists!", http.StatusNoContent)
	} // else
} // getImage
