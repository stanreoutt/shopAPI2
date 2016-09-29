package core

import (
	"log"
	"time"

	"database/sql"

	geo "github.com/kellydunn/golang-geo"
)

// Branch represents the shop branch
type Branch struct {
	ID          int            `json:"id"`
	Region      sql.NullString `json:"region"`
	City        sql.NullString `json:"city"`
	Address     sql.NullString `json:"address"`
	Longitude   float64        `json:"longitude"`
	Latitude    float64        `json:"latitude"`
	Point       geo.Point      `json:"-"`
	Schedule    sql.NullString `json:"schedule"`
	HasBar      bool           `json:"has_bar"`
	HasVIP      bool           `json:"has_vip"`
	Description sql.NullString `json:"description"`
	Phones      sql.NullString `json:"phones"`
	OpenDate    time.Time      `json:"open_date"`
	ShopType    int            `json:"shop_type"`
	Metro       sql.NullString `json:"metro"`
}

// BranchPool is a branch repository
type BranchPool struct {
	pool map[int]*Branch
}

// AddBranch adds a branch to the repository
func (p *BranchPool) AddBranch(b *Branch) error {
	if p.pool == nil {
		p.pool = make(map[int]*Branch)
	}

	// todo perhaps some actions on the new branch
	log.Printf("+++ Added branch [%s]\n", b.Address.String)
	p.pool[b.ID] = b

	return nil
}
