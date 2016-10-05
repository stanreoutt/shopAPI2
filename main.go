package main

import (
	"fmt"
	"log"
	"net/http"
	"shopd/core"
	"shopd/handlers"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/gorilla/mux"
)

// PORT server's listening on
const PORT = 3030

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile) /* configuring default logger to show additional info */
	ac := &core.AppContext{}

	// requiring config path from command line
	confFlag := kingpin.Flag("config", "Path to configuration file").Short('c').Required().String()
	kingpin.Parse()
	configFilepath := *confFlag

	log.Println("*** Initializing daemon")
	ac.Init(configFilepath) /* initializing main context, etc */

	// configuring http server
	r := mux.NewRouter()

	handlers.Context = ac /* sharing app context with handlers package */

	// defining routes
	r.HandleFunc("/api/v1.1/{lang}/region", handlers.RegionHandler).Methods("GET")
	r.HandleFunc("/api/v1.1/{lang}/city", handlers.CityHandler).Methods("GET")
	r.HandleFunc("/api/v1.1/{lang}/shop", handlers.ShopHandler).Methods("GET")

	log.Printf("*** Started listening on %d port\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
}
