package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//================================================================================
// Create Endpoint [/create]
//================================================================================

// Creates a new block and adds it to the user's blockchain
func update_block(w http.ResponseWriter, r *http.Request) {

	// Make sure this endpoint is only accessible at "/create".
	if r.URL.Path != "/update" {
		log.Println(r.URL.Path)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	if r.Method == "PUT" {
		// Parse JSON to Block Struct
		var blockToUpdate Block = Block{}
		var success bool = parseJsonToBlock(&blockToUpdate, r)
		if !success {
			http.Error(w, "JSON Parse Failed", http.StatusBadRequest)
			return
		} // if

		// Can't update a block from an empty chain
		if _, exist := UserChains[blockToUpdate.UserId]; !exist {
			http.Error(w, "Blockchain is empty", http.StatusNoContent)
		} // if

		// Find block in chain
		var createOriginHash string = ""
		for _, block := range UserChains[blockToUpdate.UserId].Chain {

			// Block to delete exists
			if blockToUpdate.Hash == block.Hash {
				createOriginHash = blockToUpdate.Hash
			} // if
		} // for

		if createOriginHash == "" {
			http.Error(w, "Block does not exist in user's blockchain", http.StatusNoContent)
		} // if

		// Append a new block to user's blockchain
		if storedUserChain, exist := UserChains[blockToUpdate.UserId]; exist {
			// Set block type and originHash
			blockToUpdate.BlockType = Update
			blockToUpdate.CreateOriginHash = createOriginHash

			calcBlockMetadata(&blockToUpdate, storedUserChain.Chain)

			// Add to map of blockchains
			storedUserChain.Chain = append(storedUserChain.Chain, blockToUpdate)
			UserChains[blockToUpdate.UserId] = storedUserChain
			if blockJson, err := json.Marshal(&UserChains[blockToUpdate.UserId].Chain[blockToUpdate.Index]); err == nil {
				w.Write(blockJson)
			} else {
				http.Error(w, "Marshal Failed", http.StatusNoContent)
			} // if-else
		} // if
	} // if
} // update_block
