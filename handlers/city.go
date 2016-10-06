package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shopd/core"
)

// CityHandler returns a full list of cities
func CityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var output []byte
	branches, err := core.GetAllCities(Context)
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
