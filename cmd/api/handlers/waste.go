package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/zthiagovalle/waste-api-go/internal/infra/repository"
	"github.com/zthiagovalle/waste-api-go/internal/usecase"
)

func ListWastes(c *fiber.Ctx) error {
	repository := repository.NewWasteGoogleSheetsRepository(os.Getenv("GOOGLE_SHEET_URL"))
	usecase := usecase.GetWaste{WasteRepository: repository}
	wastes, err := usecase.Execute()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(200).JSON(wastes)
}
