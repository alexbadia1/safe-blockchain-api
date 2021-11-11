package main

import (
	"encoding/json"
	"net/http"
)

const (
	raw      string = "raw"
	filtered string = "filtered"
)

//================================================================================
// Chain Endpoint [/chain]
//================================================================================

// Sends JSON representation of the user's blockchain
func getChain(w http.ResponseWriter, r *http.Request) {
	// Make sure this endpoint is only accessible at "/chain".
	if r.URL.Path != "/chain" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	// Parse URL params
	userId := getUserIdFromUrlQueryParams(w, r)
	view := getViewFromUrlQueryParams(w, r)

	// Find user's chain
	userChain, exist := UserChains[userId]
	if !exist {
		http.Error(w, "Blockchain for the specified user ID doesn't exist", http.StatusNoContent)
	} // if

	// Filter chain
	filteredBlocks := make([]Block, 0)
	if view == filtered {
		createBlocks := getCreateBlocks(userChain.Chain)

	createBlockLoop:
		for _, createBlock := range createBlocks {
			for i := len(userChain.Chain) - 1; i >= 0; i-- {
				var currBlock Block = userChain.Chain[i]

				// Found a block that either updated, or deleted
				// the original block or is the original block itself.
				if createBlock.CreateOriginHash == userChain.Chain[i].CreateOriginHash {

					// Ignore genesis or deleted blocks
					if currBlock.BlockType == Genesis || currBlock.BlockType == Delete {
						continue createBlockLoop
					} // if

					// Add most recent version of the block
					if currBlock.BlockType == Update || currBlock.BlockType == Create {
						filteredBlocks = append(filteredBlocks, currBlock)
						continue createBlockLoop
					} // if
				} // if
			} // for
		} // for
	} else {
		filteredBlocks = userChain.Chain
	} // if-else

	// Respond with chain as JSON
	if chainJson, err := json.Marshal(&filteredBlocks); err == nil {
		w.Write(chainJson)
	} else {
		http.Error(w, "User found, but Marshal Failed", http.StatusNoContent)
	} // if-else
} // chain

//================================================================================
// Helper Methods
//================================================================================

// Gets userId query param from url. Fails, if no userId is given.
func getUserIdFromUrlQueryParams(w http.ResponseWriter, r *http.Request) string {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "No user id was given!", http.StatusNoContent)
	} // if

	return userId
} // getUserIdFromUrlQueryParams

// Gets view query param from url. Defaults to "filtered".
func getViewFromUrlQueryParams(w http.ResponseWriter, r *http.Request) string {
	view := r.URL.Query().Get("view")
	if view != raw && view != filtered {
		return filtered
	} // if

	return view
} // getViewFromUrlQueryParams

func getCreateBlocks(userChain []Block) []Block {
	createBlocks := make([]Block, 0)
	for _, block := range userChain {
		if block.BlockType == Create {
			createBlocks = append(createBlocks, block)
		} // if
	} // for
	return createBlocks
} // getCreateBlocks
