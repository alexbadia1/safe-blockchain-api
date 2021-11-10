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

	// Get input
	var tmpBloc Block = Block{}
	parseJsonToBlock(&tmpBloc, r)

	// Search for user's blockchain
	if userChain, exist := UserChains[tmpBloc.UserId]; exist {
		if chainJson, err := json.Marshal(&userChain.Chain); err == nil {
			w.Write(chainJson)
		} else {
			http.Error(w, "Marshal Failed", http.StatusNoContent)
		} // if-else
	} else {
		http.Error(w, "Blockchain not found", http.StatusNoContent)
	} // if-else
} // chain
