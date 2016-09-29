package core

import (
	"database/sql"
	"log"
	"strings"

	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
)

// LoadEverything loads everything from postgres
func LoadEverything(c *AppContext) error {
	// connecting to database
	log.Println("*** Connecting to database")
	db, err := sql.Open("postgres", "user=postgres password='r00tme' dbname=shops sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// loading branches
	log.Println("*** Loading branches")
	rows, err := db.Query(`
    SELECT id, region, city, address, longitude, latitude, 
            phones, schedule, has_bar, has_vip, description 
    FROM branches_branch
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		b := &Branch{}
		var openDate string
		err = rows.Scan(&b.ID, &b.Region, &b.City, &b.Address, &b.Longitude, &b.Latitude, &b.Phones,
			&b.Schedule, &b.HasBar, &b.HasVIP, &b.Description) /* sequentially mapping colums to their corresponding fields in Branch struct; SELECT's field order is retained */
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(openDate) != "" {
			b.OpenDate, err = time.Parse("1982-Aug-08", openDate) /* converting string to time.Time */
			if err != nil {
				log.Fatal(err)
			}
		}

		err = c.Data.Branches.AddBranch(b) /* adding branch to a pool */
		if err != nil {
			log.Fatal(err)
		}
	}

	spew.Dump(c.Data)

	return nil
}
