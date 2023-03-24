package repository

import (
	"io"
	"net/http"
	"strings"

	"github.com/zthiagovalle/waste-api-go/internal/entity"
)

type WasteGoogleSheetsRepository struct {
	Url string
}

func NewWasteGoogleSheetsRepository(url string) *WasteGoogleSheetsRepository {
	return &WasteGoogleSheetsRepository{Url: url}
}

func (w WasteGoogleSheetsRepository) GetWastes() ([]entity.Waste, error) {
	request, err := http.NewRequest("GET", w.Url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	wastes := convertCSVInStruct(string(respBody))

	return wastes, nil
}

func convertCSVInStruct(csv string) []entity.Waste {
	rows := strings.Split(csv, "\n")
	wastes := []entity.Waste{}
	for _, row := range rows {
		blocks := strings.Split(row, ",")

		wasteNameBlock := strings.TrimSpace(blocks[0])
		wasteImageUrlBlock := strings.TrimSpace(blocks[3])

		addressBlock := strings.Split(blocks[1], "|")
		coordinatesBlock := strings.Split(blocks[2], "|")

		addresses := []entity.Address{}
		for i := 0; i < len(addressBlock); i++ {
			addresses = append(addresses,
				entity.NewAddress(
					strings.TrimSpace(addressBlock[i]),
					strings.TrimSpace(strings.Split(coordinatesBlock[i], " ")[0]),
					strings.TrimSpace(strings.Split(coordinatesBlock[i], " ")[1])))
		}

		waste := entity.NewWaste(wasteNameBlock, addresses, wasteImageUrlBlock)
		wastes = append(wastes, waste)
	}

	return wastes
}
