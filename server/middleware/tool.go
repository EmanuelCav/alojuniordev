package middleware

import (
	"context"
	"time"

	"github.com/EmanuelCav/alojuniordev/helper"
	"github.com/EmanuelCav/alojuniordev/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ToolValid(tool models.ToolCreateModel) string {

	var toolValid models.ToolModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := helper.ConnectionTool().FindOne(ctx, bson.M{"tool": tool.Tool}).Decode(&toolValid)

	if err == nil {
		if toolValid.Tool == tool.Tool {
			return "La herramienta ya ha sido registrada"
		}
	}

	if len(tool.Desciprtion) < 50 || len(tool.ShortDesciprtion) < 20 {
		return "La descripción propuesta es muy corta. Por favor escribe más"
	}

	if !helper.ValidateSomeCharacters(tool.Tool) || !helper.ValidateSomeCharacters(tool.Desciprtion) || !helper.ValidateSomeCharacters(tool.ShortDesciprtion) {
		return "Caracters como <,> no están permitidos"
	}

	return ""

}

func CategoryValid(category models.CategoryCreateModel) string {

	var categoryValid models.CategoryModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := helper.ConnectionCategory().FindOne(ctx, bson.M{"category": category.Category}).Decode(&categoryValid)

	if err == nil {
		if categoryValid.Category == category.Category {
			return "Category has been already registered"
		}
	}

	if !helper.ValidateString(category.Category) {
		return "Category only accepts letters"
	}

	return ""

}
