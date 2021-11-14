package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//================================================================================
// Mine Endpoint [/mine]
//================================================================================

func mine(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

	// Get data to hash
	var concatenatestrings strings.Builder
	concatenatestrings.WriteString(strings.TrimSpace(bp.CertificateCatgeory))
	concatenatestrings.WriteString(strings.TrimSpace(bp.CertificateToken))
	concatenatestrings.WriteString(strings.TrimSpace(bp.CertificateUrl))
	concatenatestrings.WriteString(strings.TrimSpace(bp.DateRange))
	concatenatestrings.WriteString(strings.TrimSpace(bp.DegreeName))
	concatenatestrings.WriteString(strings.TrimSpace(bp.Description))
	concatenatestrings.WriteString(strings.TrimSpace(strconv.FormatInt(bp.Index, 10)))
	concatenatestrings.WriteString(strings.TrimSpace(bp.InstitionName))
	concatenatestrings.WriteString(strings.TrimSpace(bp.PreviousHash))
	concatenatestrings.WriteString(strings.TrimSpace(strconv.FormatInt(bp.Timestamp, 10)))
	concatenatestrings.WriteString(strings.TrimSpace(bp.UserId))
	var msg string = concatenatestrings.String()

	for {
		guessMsg := msg + strconv.FormatInt(bp.Nonce, 10)

		// Calculate hash
		h := sha256.New()
		h.Write([]byte(guessMsg))
		var tmp string = fmt.Sprintf("%x", h.Sum(nil))

		if tmp[0:5] == "00000" {
			finalHash = tmp
			break
		} // if

		bp.Nonce = bp.Nonce + 1
	} // for

	bp.Hash = finalHash
	bp.BlockType = bType
	bp.CreateOriginHash = CreateOriginHash
} // mine
