package dtos

type PropertyDto struct {
	Id          int     `json:"id"`
	Tittle      string  `json:"tittle"`
	Description string  `json:"description"`
	Address     Address `json:"address"`
	Size        int     `json:"size"`
	Rooms       int     `json:"rooms"`
	Bathrooms   int     `json:"bathrooms"`
	Service     string  `json:"service"`
	Price       int     `json:"price"`
}
type Address struct {
	Street  string
	City    string
	State   string
	Country string
}
type PropertiesDto []PropertyDto
