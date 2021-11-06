package main

import (
	"encoding/json"
	"net/http"
)

func getChain(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/chain".
	if r.URL.Path != "/chain" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	// Get input
	var tmpBloc Block = Block{}
	parseJson(&tmpBloc, r)

	// Search for blo
	if userChain, exist := UserChains[tmpBloc.UserId]; exist {
		if blockJson, err := json.Marshal(&userChain.chain); err == nil {
			w.Write(blockJson)
		} else {
			http.Error(w, "Marshal Failed", http.StatusNoContent)
		} // if-else
	} else {
		http.Error(w, "Blockchain not found", http.StatusNoContent)
	} // if-else
} // chain
