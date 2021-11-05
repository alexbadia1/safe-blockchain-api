package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Map to store all users blockchains
var UserChains = make(map[string]Blockchain)

// Returns the current status of the server
func status(w http.ResponseWriter, r *http.Request) {
	log.Println("\"alive\":true")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
} // status

// Create endpoint
func create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	} // if

	switch r.Method {
	// Serve create.html
	case "GET":
		http.ServeFile(w, r, "templates/create.html")
	// Parse the form
	case "POST":
		// Parse JSON to Block Struct
		var bloc Block
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		} // if
		if err := json.Unmarshal(body, &bloc); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		} // if

		// Check for user's chain
		userChain, exist := UserChains[bloc.UserId]
		if !exist {
			// Add new User Blockchain
			var newBlockchain *Blockchain = new(Blockchain)
			UserChains[bloc.UserId] = *newBlockchain
		} // if

		// Add to chain
		userChain.chain = append(userChain.chain, bloc)

		fmt.Printf("%v\n", UserChains)
		// // Parse Form
		//
		// if err := r.ParseForm(); err != nil {
		// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
		// 	return
		// } // if
		//
		// block := Block{
		// 	InstitionName: r.FormValue("institution_name"),
		// 	DegreeName:    r.FormValue("degree_name"),
		// 	DateRange:     r.FormValue("dates"),
		// 	Description:   r.FormValue("description"),
		// } // Block
		//
		// if b, exist := UserChains[bloc.UserId]; exist {
		// 	w.Write(blockJson)
		// } // if
		if storedUserChain, exist := UserChains[bloc.UserId]; exist {
			var lastBlock Block = storedUserChain.chain[len(storedUserChain.chain)-1]
			blockJson, err := json.Marshal(&lastBlock)
			if err != nil {
				http.Error(w, "Marshal Failed", http.StatusBadRequest)
			} else {
				w.Write(blockJson)
			} // if-else
		} // if
	default:
		w.WriteHeader(http.StatusNotImplemented)
	} // switch
} // newBlock

func main() {
	// TODO: Run locally
	// port := "8000"

	// TODO: For production
	port := os.Getenv("PORT")
	http.HandleFunc("/", status)
	http.HandleFunc("/create", create)
	log.Print("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} // main
