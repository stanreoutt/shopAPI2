package core

import "time"

// Branch represents the shop branch
type Branch struct {
	ID          int64      `json:"id"`
	Region      string     `json:"region"`
	City        string     `json:"city"`
	Address     string     `json:"address"`
	Longitude   float64    `json:"longitude"`
	Latitude    float64    `json:"latitude"`
	Schedule    string     `json:"schedule"`
	HasBar      bool       `json:"has_bar"`
	HasVIP      bool       `json:"has_vip"`
	Description string     `json:"description"`
	Phones      string     `json:"phones"`
	OpenDate    *time.Time `json:"open_date,omitempty"`
	ShopType    int64      `json:"shop_type"`
	Metro       string     `json:"metro"`
	Dist        float64    `json:"dist,omitempty"`
}

// BranchGeoMap mapping geo locations to shops for easier in-memory search
type BranchGeoMap struct {
	Longitude float64
	Latitude  float64
	Branch    *Branch
}

// BranchPool is a branch repository
type BranchPool struct {
	pool   map[int64]*Branch
	geomap []*BranchGeoMap
}

// AddBranch adds a branch to the repository
func (p *BranchPool) AddBranch(b *Branch) error {
	if p.pool == nil {
		p.pool = make(map[int64]*Branch)
	}

	// todo perhaps some actions on the new branch
	//log.Printf("+++ Added branch [%s]\n", b.Address.String)
	p.pool[b.ID] = b

	return nil
}
