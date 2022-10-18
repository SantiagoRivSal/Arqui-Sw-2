package dtos

type PropertyDto struct {
	Id          string  `json:"id"`
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
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}
type PropertiesDto []PropertyDto
