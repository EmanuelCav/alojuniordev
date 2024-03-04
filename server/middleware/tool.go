package middleware

import (
	"context"
	"time"

	"github.com/EmanuelCav/alojuniordev/helper"
	"github.com/EmanuelCav/alojuniordev/models"
	"go.mongodb.org/mongo-driver/bson"
)

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
