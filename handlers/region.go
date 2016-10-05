package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shopd/core"
)

// RegionHandler exports all regions
func RegionHandler(w http.ResponseWriter, r *http.Request) {
	var output []byte

	branches, err := core.GetAllRegions(Context)
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
