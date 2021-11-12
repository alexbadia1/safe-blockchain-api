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

	// Make sure this endpoint is only accessible at "/put".
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

		// Prevent user from trying to update the block's userId
		// by checking the map of blockchains for the userId first.
		if _, exist := UserChains[blockToUpdate.UserId]; !exist {
			http.Error(w, "Blockchain for user doesn't exist", http.StatusNoContent)
			return
		} // if

		// Find block in chain
		var createOriginHash string = ""
		for i := len(UserChains[blockToUpdate.UserId].Chain) - 1; i >= 0; i-- {
			var currBlock Block = UserChains[blockToUpdate.UserId].Chain[i]

			// Can't update a deleted block
			if currBlock.BlockType == Delete && currBlock.CreateOriginHash == blockToUpdate.CreateOriginHash {
				http.Error(w, "Can't update a block that was already deleted", http.StatusNoContent)
				return
			} // if

			// Block to update exists
			if blockToUpdate.CreateOriginHash == currBlock.CreateOriginHash {
				createOriginHash = blockToUpdate.CreateOriginHash
			} // if
		} // for

		if createOriginHash == "" {
			http.Error(w, "Block does not exist in user's blockchain", http.StatusNoContent)
		} // if

		// Append a new block to user's blockchain
		if storedUserChain, exist := UserChains[blockToUpdate.UserId]; exist {
			calcBlockMetadata(&blockToUpdate, storedUserChain.Chain)

			// Set block type and originHash
			blockToUpdate.BlockType = Update
			blockToUpdate.CreateOriginHash = createOriginHash

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
