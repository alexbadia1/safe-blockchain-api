package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//================================================================================
// Create Endpoint [/create]
//================================================================================

// Creates a new block and adds it to the user's blockchain
func new_block(w http.ResponseWriter, r *http.Request) {

	// Make sure this endpoint is only accessible at "/create".
	if r.URL.Path != "/create" {
		log.Println(r.URL.Path)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} // if

	if r.Method == "POST" {
		// Parse JSON to Block Struct
		var bloc Block = Block{}
		var success bool = parseJsonToBlock(&bloc, r)
		if !success {
			http.Error(w, "JSON Parse Failed", http.StatusBadRequest)
			return
		} // if

		// If needed, create a new blockchain
		if _, exist := UserChains[bloc.UserId]; !exist {
			// New blochain
			var newBlockchain Blockchain = Blockchain{}

			// New genesis block
			var genesisBlock Block = Block{}
			genesisBlock.BlockType = Genesis
			calcBlockMetadata(&genesisBlock, UserChains[bloc.UserId].Chain)
			newBlockchain.Chain = append(newBlockchain.Chain, genesisBlock)

			// Update the reference in the map
			UserChains[bloc.UserId] = newBlockchain
		} // if

		// Append a new block to user's blockchain
		if storedUserChain, exist := UserChains[bloc.UserId]; exist {
			calcBlockMetadata(&bloc, storedUserChain.Chain)

			// Set block type and originHash
			bloc.BlockType = Create
			bloc.CreateOriginHash = bloc.Hash

			// Add to map of blockchains
			storedUserChain.Chain = append(storedUserChain.Chain, bloc)
			UserChains[bloc.UserId] = storedUserChain
			if blockJson, err := json.Marshal(&UserChains[bloc.UserId].Chain[bloc.Index]); err == nil {
				w.Write(blockJson)
			} else {
				http.Error(w, "Marshal Failed", http.StatusNoContent)
			} // if-else
		} // if
	} // if
} // new_block

// Calculates Block Metadata:
//  - index
//  - timestamp
//  - previous hash
//  - current hash
func calcBlockMetadata(bp *Block, chain []Block) {
	var size int = len(chain)
	bp.Index = int64(size)
	bp.Timestamp = time.Now().Unix()
	if size > 0 {
		bp.PreviousHash = chain[size-1].Hash
	} // if

	// Calculate hash
	mineBlock(bp)

	// Previous hash of first block is it's own hash
	if size == 0 {
		bp.PreviousHash = bp.Hash
	} // if
} // calcBlockMetadata
