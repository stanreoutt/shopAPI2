package core

import (
	"database/sql"
	"log"
	"strings"

	"time"

	"fmt"

	_ "github.com/lib/pq"
)

const sqlConnectionString = "user=postgres password='r00tme' dbname=shops sslmode=disable"

// LoadEverything loads everything from postgres
func LoadEverything(c *AppContext) error {
	// connecting to database
	log.Println("*** Connecting to database")
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// preloading all branches
	branches, err := GetAllBranches()
	if err != nil {
		log.Fatal(err)
	}
	for _, b := range branches {
		c.Data.Branches.AddBranch(b)                                                                       /* adding branch to a pool */
		c.Data.Branches.geomap = append(c.Data.Branches.geomap, &BranchGeoMap{b.Longitude, b.Latitude, b}) /* adding a geomap to a geoloc map */
	}

	// preloading cities
	cities, err := GetAllCities()
	if err != nil {
		log.Fatal(err)
	}
	c.Data.Cities = cities

	// preloading regions
	regions, err := GetAllRegions()
	if err != nil {
		log.Fatal(err)
	}
	c.Data.Regions = regions

	return nil
}

// GetAllBranches fetching all branch records from database
func GetAllBranches() ([]*Branch, error) {
	// connecting to database
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
    SELECT id, region, city, address, longitude, latitude, 
            phones, schedule, has_bar, has_vip, description, open_date, shop_type, metro, 0
    FROM branches_branch
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return transformBranches(rows)
}

// GetBranchesByCity fetching branches by city name
func GetBranchesByCity(cityName string) ([]*Branch, error) {
	// connecting to database
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
    SELECT id, region, city, address, longitude, latitude, 
            phones, schedule, has_bar, has_vip, description, open_date, shop_type, metro, 0
    FROM branches_branch
	WHERE city=$1
    `, cityName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return transformBranches(rows)
}

// GetBranchesByRegion fetching branches by region name
func GetBranchesByRegion(regionName string) ([]*Branch, error) {
	// connecting to database
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
    SELECT id, region, city, address, longitude, latitude, 
            phones, schedule, has_bar, has_vip, description, open_date, shop_type, metro, 0
    FROM branches_branch
	WHERE region LIKE $1
    `, fmt.Sprintf("%s%s", regionName, "%")) /* zomfg hackode to get a simple percent char */

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return transformBranches(rows)
}

// GetBranchesByPoint fetching branches by geospatial coordinates with a radius in miles
func GetBranchesByPoint(lon, lat, radius float64) ([]*Branch, error) {
	// connecting to database
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
    SELECT id, region, city, address, longitude, latitude, 
            phones, schedule, has_bar, has_vip, description, open_date, shop_type, metro, point($1, $1) <@> point(longitude, latitude)::point AS shop_distance
    FROM branches_branch
	WHERE (point($1, $2) <@> point(longitude, latitude)) <= $3
    `, lon, lat, radius) /* zomfg hackode to get a simple percent char */

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return transformBranches(rows)
}

// GetBranchesByPolygon fetching branches by geospatial coordinates within a box
func GetBranchesByPolygon(lon, lat, lon2, lat2 float64) ([]*Branch, error) {
	// connecting to database
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
    SELECT id, region, city, address, longitude, latitude, 
            phones, schedule, has_bar, has_vip, description, open_date, shop_type, metro, 0
    FROM branches_branch
	WHERE longitude BETWEEN $1 AND $3 AND latitude BETWEEN $2 AND $4
    `, lon, lat, lon2, lat2) /* zomfg hackode to get a simple percent char */

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return transformBranches(rows)
}

func transformBranches(rows *sql.Rows) ([]*Branch, error) {
	var branches []*Branch
	var id, shopType sql.NullInt64
	var lon, lat, dist sql.NullFloat64
	var region, city, address, phones, schedule, description, openDate, metro sql.NullString
	var hasBar, hasVIP sql.NullBool

	for rows.Next() {
		err := rows.Scan(&id, &region, &city, &address, &lon, &lat, &phones,
			&schedule, &hasBar, &hasVIP, &description, &openDate, &shopType, &metro, &dist) /* sequentially mapping colums to their corresponding fields in Branch struct; SELECT's field order is retained */

		if err != nil {
			log.Fatal(err)
		}

		var openDateTime time.Time
		if strings.TrimSpace(openDate.String) != "" {
			openDateTime, err = time.Parse("2006-1-2T15:04:05Z", openDate.String) /* converting string to time.Time */
			if err != nil {
				log.Fatal(err)
			}
		}

		b := &Branch{
			ID:          id.Int64,
			Region:      region.String,
			City:        city.String,
			Address:     address.String,
			Longitude:   lon.Float64,
			Latitude:    lat.Float64,
			Phones:      phones.String,
			Schedule:    schedule.String,
			HasBar:      hasBar.Bool,
			HasVIP:      hasVIP.Bool,
			Description: description.String,
			OpenDate:    &openDateTime,
			ShopType:    shopType.Int64,
			Metro:       metro.String,
			Dist:        dist.Float64,
		}

		branches = append(branches, b)
	}

	return branches, nil
}

// GetAllCities retrieves all cities from database
func GetAllCities() ([]*City, error) {
	// connecting to database
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var cities []*City
	var id sql.NullInt64
	var name sql.NullString

	rows, err := db.Query("SELECT id, name FROM branches_city")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		c := &City{
			ID:   id.Int64,
			Name: name.String,
		}

		cities = append(cities, c)
	}

	return cities, nil
}

// GetAllRegions retrieves all regions from database
func GetAllRegions() ([]*Region, error) {
	// connecting to database
	db, err := sql.Open("postgres", sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var regions []*Region
	var id sql.NullInt64
	var name sql.NullString

	rows, err := db.Query("SELECT id, name FROM branches_region")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		r := &Region{
			ID:   id.Int64,
			Name: name.String,
		}

		regions = append(regions, r)
	}

	return regions, nil
}
