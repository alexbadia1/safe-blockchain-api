package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//================================================================================
// Delete Endpoint [/delete]
//================================================================================

func delete_block(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/create".
	if r.URL.Path != "/delete" {
		log.Println(r.URL.Path)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	if r.Method == "DELETE" {
		// Parse JSON to Block Struct
		var blockToDelete Block = Block{}
		var success bool = parseJsonToBlock(&blockToDelete, r)
		if !success {
			http.Error(w, "JSON Parse Failed", http.StatusBadRequest)
			return
		} // if

		// Can't delete a block from an empty chain
		if _, exist := UserChains[blockToDelete.UserId]; !exist {
			http.Error(w, "Blockchain is empty", http.StatusNoContent)
		} // if

		// Find block in chain
		var createOriginHash string = ""
		for i := len(UserChains[blockToDelete.UserId].Chain) - 1; i >= 0; i-- {
			var currBlock Block = UserChains[blockToDelete.UserId].Chain[i]

			// Block to delete exists
			if blockToDelete.Hash == currBlock.CreateOriginHash {
				createOriginHash = blockToDelete.Hash
			} // if
		} // for

		if createOriginHash == "" {
			http.Error(w, "Block does not exist in user's blockchain", http.StatusNoContent)
		} // if

		// Append a new DELETE block to user's blockchain
		if storedUserChain, exist := UserChains[blockToDelete.UserId]; exist {
			// Store origin hash of the deleted block
			blockToDelete.BlockType = Delete
			blockToDelete.CreateOriginHash = createOriginHash
			calcBlockMetadata(&blockToDelete, storedUserChain.Chain)

			storedUserChain.Chain = append(storedUserChain.Chain, blockToDelete)
			UserChains[blockToDelete.UserId] = storedUserChain
			if blockJson, err := json.Marshal(&UserChains[blockToDelete.UserId].Chain[blockToDelete.Index]); err == nil {
				w.Write(blockJson)
			} else {
				http.Error(w, "Marshal Failed", http.StatusNoContent)
			} // if-else
		} // if
	} // if
} // delete_block
