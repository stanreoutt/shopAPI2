package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shopd/core"
)

// ShopCityHandler exporting shop list by city; city comes in GET
func ShopCityHandler(w http.ResponseWriter, r *http.Request) {
	var output []byte

	cityName := r.URL.Query().Get("city")
	branches, err := core.GetBranchesByCity(cityName)
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
