package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shopd/core"
)

// AllShopsHandler this handler exports a full shop listing
func AllShopsHandler(w http.ResponseWriter, r *http.Request) {
	var output []byte

	branches, err := core.GetAllBranches()
	if err != nil {
		w.Write([]byte("an error has occured, please contact a service administrator"))
		log.Println(err)
	}

	output, err = json.Marshal(branches)
	if err != nil {
		w.Write([]byte("an error has occured, please contact a service administrator"))
		log.Println(err)
	}

	w.Write(output)
}
