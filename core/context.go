package core

import (
	"fmt"
	"log"
	"strconv"

	ini "github.com/go-ini/ini"
)

// AppContext application's main context
type AppContext struct {
	ListenAt string
	Database struct {
		Hostname         string
		Username         string
		Password         string
		Database         string
		ConnectionString string
	}
	Misc struct {
		DefaultPointRadiusOnMap float64
	}
	Data Data
}

// Data data container for the main context
type Data struct {
	Regions  []*Region
	Branches *BranchPool
	Cities   []*City
}

// Init main initialization function; not used atm
func (c *AppContext) Init(configPath string) error {
	// initializing main context
	c.Data = Data{
		Regions:  make([]*Region, 0),
		Branches: &BranchPool{},
		Cities:   make([]*City, 0),
	}

	log.Println("*** Loading main configuration file")
	mainConfig, err := ini.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	genSec := mainConfig.Section("general")
	c.ListenAt = genSec.Key("host").Value()

	dbSec := mainConfig.Section("database")
	c.Database.Hostname = dbSec.Key("hostname").Value()
	c.Database.Username = dbSec.Key("username").Value()
	c.Database.Password = dbSec.Key("password").Value()
	c.Database.Database = dbSec.Key("database").Value()
	c.Database.ConnectionString = fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable",
		c.Database.Username,
		c.Database.Password,
		c.Database.Database)

	miscSec := mainConfig.Section("misc")
	c.Misc.DefaultPointRadiusOnMap, err = strconv.ParseFloat(miscSec.Key("default-point-radius").Value(), 64)
	if err != nil {
		log.Fatal(err)
	}

	err = LoadEverything(c)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
