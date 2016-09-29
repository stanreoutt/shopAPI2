package core

// Type representing branch type i.e. shop, euroset, qiwi, etc.
type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
