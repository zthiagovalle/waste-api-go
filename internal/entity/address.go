package entity

type Address struct {
	Street    string `json:"street"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

func NewAddress(street string, latitude string, longitude string) Address {
	return Address{
		Street:    street,
		Latitude:  latitude,
		Longitude: longitude,
	}
}
