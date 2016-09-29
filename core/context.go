package core

import "log"

// AppContext application's main context
type AppContext struct {
	Data Data
}

// Data data container for the main context
type Data struct {
	Regions  map[int]*Region
	Branches *BranchPool
}

// Init main initialization function
func (c *AppContext) Init() error {
	log.Println("*** Initializing Shop API daemon")

	// initializing main context
	c.Data = Data{
		Regions:  make(map[int]*Region),
		Branches: &BranchPool{},
	}

	err := LoadEverything(c)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
