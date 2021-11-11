package main

import "net/http"

//================================================================================
// Images Endpoint [/images]
//================================================================================

// Serves image
func getImage(w http.ResponseWriter, r *http.Request) {
	if name := r.URL.Query().Get("name"); name == "bobby" {
		http.ServeFile(w, r, "images/bobby.png")
	} else if name := r.URL.Query().Get("name"); name == "certified" {
		http.ServeFile(w, r, "images/certified.png")
	} else {
		http.Error(w, "Image doesn't exist of exists!", http.StatusNoContent)
	} // else
} // getImage
