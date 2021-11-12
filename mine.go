package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//================================================================================
// Mine Endpoint [/mine]
//================================================================================

func mine(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Parse JSON to Block Struct
		var bloc Block = Block{}
		var success bool = parseJsonToBlock(&bloc, r)
		if !success {
			http.Error(w, "JSON Parse Failed", http.StatusBadRequest)
			return
		} // if

		mineBlock(&bloc)

		if blockJson, err := json.Marshal(&bloc); err == nil {
			w.Write(blockJson)
		} else {
			http.Error(w, "Marshal Failed", http.StatusNoContent)
		} // if-else
	} // if
} // mine

// Mines a block
func mineBlock(bp *Block) {
	// Don't include block type and origin hash in hash
	var bType string = bp.BlockType
	bp.BlockType = ""

	var CreateOriginHash string = bp.CreateOriginHash
	bp.CreateOriginHash = ""

	bp.Nonce = 0
	var finalHash string = ""

	for {
		// Calculate hash
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("%v", bp)))
		var tmp string = fmt.Sprintf("%x", h.Sum(nil))

		if tmp[0:3] == "000" {
			finalHash = tmp
			break
		} // if

		bp.Nonce = bp.Nonce + 1

		log.Printf("%v", tmp)
	} // for

	bp.Hash = finalHash
	bp.BlockType = bType
	bp.CreateOriginHash = CreateOriginHash
} // mine
