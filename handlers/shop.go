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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var output []byte
	var branches []*core.Branch
	var err error

	var lon, lat, lon2, lat2 = r.URL.Query().Get("lon"), r.URL.Query().Get("lat"),
		r.URL.Query().Get("lon2"), r.URL.Query().Get("lat2")
	var exprFindByPoint = lon != "" && lat != "" && lon2 == "" && lat2 == ""
	var exprFindByPolygon = lon != "" && lat != "" && lon2 != "" && lat2 != ""
	var exprInvalidGeo = (lon != "" || lat != "") && (lon2 == "" || lat2 == "")

	if !exprFindByPoint && !exprFindByPolygon && exprInvalidGeo {
		w.Write([]byte("an error has occured, given geo criteria is inconsistent"))
		log.Println(err)
		return
	}

	var offsetLen, offsetPage int64 /* items per page; page offset */
	offsetLen, offsetPage = 100000, 0

	if olen := r.URL.Query().Get("olen"); olen != "" {
		offsetLen, err = strconv.ParseInt(olen, 10, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	}

	if opage := r.URL.Query().Get("opage"); opage != "" {
		offsetPage, err = strconv.ParseInt(opage, 10, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	}

	if city := r.URL.Query().Get("city"); strings.TrimSpace(city) != "" { /* by city */
		branches, err = core.GetBranchesByCity(Context, city, offsetLen, offsetPage)
	} else if region := r.URL.Query().Get("region"); strings.TrimSpace(region) != "" { /* by region */
		branches, err = core.GetBranchesByRegion(Context, region, offsetLen, offsetPage)
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

		branches, err = core.GetBranchesByPoint(Context, flon, flat, Context.Misc.DefaultPointRadiusOnMap, offsetLen, offsetPage)
	} else if exprFindByPolygon {
		// todo find by polygon/box
	} else { /* default case, exporting all branches */
		branches, err = core.GetAllBranches(Context, offsetLen, offsetPage)
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
