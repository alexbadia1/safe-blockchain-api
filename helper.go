package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Parses JSON body from HTTP Request to Block Struct
func parseJsonToBlock(bp *Block, r *http.Request) bool {
	// Parse JSON to Block Struct
	if bodyAsByteArray, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(bodyAsByteArray, &bp); err == nil {
			return true
		} // if
	} // if
	return false
} // parseJson
