package main

import (
	"fmt"
	"log"
	"net/http"
	"shopd/core"
	"shopd/handlers"

	"github.com/gorilla/mux"
)

// PORT server's listening on
const PORT = 3030

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) /* configuring default logger to show additional info */
	ac := &core.AppContext{}
	ac.Init()

	// configuring http server
	r := mux.NewRouter()

	r.HandleFunc("/api/v1.1/{lang}/region", handlers.RegionHandler).Methods("GET")
	r.HandleFunc("/api/v1.1/{lang}/city", handlers.CityHandler).Methods("GET")
	r.HandleFunc("/api/v1.1/{lang}/allshops", handlers.AllShopsHandler).Methods("GET")
	r.HandleFunc("/api/v1.1/{lang}/shopcity", handlers.ShopCityHandler).Methods("GET")
	r.HandleFunc("/api/v1.1/{lang}/shopregion", handlers.ShopRegionHandler).Methods("GET")

	log.Printf("*** Started listening on %d port\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
}
