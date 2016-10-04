package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shopd/core"
	"strconv"
	"strings"
)

// ShopHandler handles the export of branches(shops) depending on what criterias it gets
// ie. query GET params like city, region, lat, lon, lat2, lon2
// if only lat and lon are given then searching for branches around a specified point with a radius given in configuration
// if lat2, lon2 are applied then looking within a given rectangle area
// city, region params are self-explanatory
func ShopHandler(w http.ResponseWriter, r *http.Request) {
	var output []byte
	var branches []*core.Branch
	var err error

	var lon, lat, lon2, lat2 = r.URL.Query().Get("lon"), r.URL.Query().Get("lat"), r.URL.Query().Get("lon2"), r.URL.Query().Get("lat2")
	var exprFindByPoint = lon != "" && lat != "" && lon2 == "" && lat2 == ""
	var exprFindByPolygon = lon != "" && lat != "" && lon2 != "" && lat2 != ""
	var exprInvalidGeo = (lon != "" || lat != "") && (lon2 == "" || lat2 == "")

	if !exprFindByPoint && !exprFindByPolygon && exprInvalidGeo {
		w.Write([]byte("an error has occured, given geo criteria is inconsistent"))
		log.Println(err)
		return
	}

	if city := r.URL.Query().Get("city"); strings.TrimSpace(city) != "" { /* by city */
		branches, err = core.GetBranchesByCity(city)
	} else if region := r.URL.Query().Get("region"); strings.TrimSpace(region) != "" { /* by region */
		branches, err = core.GetBranchesByRegion(region)
	} else if exprFindByPoint {
		var flon, flat float64
		flon, err = strconv.ParseFloat(lon, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		flat, err = strconv.ParseFloat(lat, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		branches, err = core.GetBranchesByPoint(flon, flat, 50)
	} else if exprFindByPolygon {

	} else { /* default case, exporting all branches */
		branches, err = core.GetAllBranches()
	}

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
