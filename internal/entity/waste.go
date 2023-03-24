package entity

type Waste struct {
	Name      string    `json:"name"`
	Addresses []Address `json:"addresses"`
	ImageUrl  string    `json:"imageUrl"`
}

func NewWaste(name string, addresses []Address, imageUrl string) Waste {
	return Waste{
		Name:      name,
		Addresses: addresses,
		ImageUrl:  imageUrl,
	}
}
