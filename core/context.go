package core

import (
	"log"

	ini "github.com/go-ini/ini"
)

// AppContext application's main context
type AppContext struct {
	ListenAt string
	Database struct {
		Hostname string
		Username string
		Password string
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

	err = LoadEverything(c)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
