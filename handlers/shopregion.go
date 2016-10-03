package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shopd/core"
)

// ShopRegionHandler exporting shop list by region; region comes in GET
func ShopRegionHandler(w http.ResponseWriter, r *http.Request) {
	var output []byte

	cityName := r.URL.Query().Get("region")
	branches, err := core.GetBranchesByRegion(cityName)
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
