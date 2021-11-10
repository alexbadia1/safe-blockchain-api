package main

import (
	"encoding/json"
	"net/http"
)

//================================================================================
// Chain Endpoint [/chain]
//================================================================================

// Sends JSON representation of the user's blockchain
func getChain(w http.ResponseWriter, r *http.Request) {
	// This is bad, I know, it's for demo purposes only...
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Make sure this endpoint is only accessible at "/chain".
	if r.URL.Path != "/chain" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	if userId := r.URL.Query().Get("userId"); userId != "" {
		// Search for user's blockchain
		if userChain, exist := UserChains[userId]; exist {
			if chainJson, err := json.Marshal(&userChain.Chain); err == nil {
				w.Write(chainJson)
			} else {
				http.Error(w, "Userfound, but Marshal Failed", http.StatusNoContent)
			} // if-else
		} else {
			http.Error(w, "Blockchain for the specified user ID doesn't exist", http.StatusNoContent)
		} // if-else
	} else {
		http.Error(w, "No user id was given!", http.StatusNoContent)
	} // else

	// Get input
	var tmpBloc Block = Block{}
	parseJsonToBlock(&tmpBloc, r)
} // chain
