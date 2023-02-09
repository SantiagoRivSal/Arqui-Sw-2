package dtos

type PropertyDto struct {
	Id          string `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Rooms       int    `json:"rooms"`
	Bathrooms   int    `json:"bathrooms"`
	Service     string `json:"service"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	UserId      int    `json:"userid"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

type PropertiesDto []PropertyDto
