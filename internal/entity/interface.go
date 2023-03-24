package entity

type WasteRepositoryInterface interface {
	GetWastes() ([]Waste, error)
}
