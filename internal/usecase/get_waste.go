package usecase

import "github.com/zthiagovalle/waste-api-go/internal/entity"

type WasteOutputDTO struct {
	Name      string             `json:"name"`
	Addresses []AddressOutputDTO `json:"addresses"`
	ImageUrl  string             `json:"imageUrl"`
}

type AddressOutputDTO struct {
	Street    string `json:"street"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

type GetWaste struct {
	WasteRepository entity.WasteRepositoryInterface
}

func (c GetWaste) Execute() ([]WasteOutputDTO, error) {
	wastes, err := c.WasteRepository.GetWastes()
	if err != nil {
		return nil, err
	}

	var wastesDTO []WasteOutputDTO
	for _, waste := range wastes {
		var addressesOutputDTO []AddressOutputDTO
		for _, address := range waste.Addresses {
			addressOutputDTO := AddressOutputDTO{
				Street:    address.Street,
				Latitude:  address.Latitude,
				Longitude: address.Longitude,
			}

			addressesOutputDTO = append(addressesOutputDTO, addressOutputDTO)
		}

		wasteDTO := WasteOutputDTO{
			Name:      waste.Name,
			ImageUrl:  waste.ImageUrl,
			Addresses: addressesOutputDTO,
		}

		wastesDTO = append(wastesDTO, wasteDTO)
	}

	return wastesDTO, nil
}
