package migration

import (
	"github.com/TUM-Dev/Campus-Backend/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

//migrate20210709193000

func (m TumDBMigrator) migrate20220713000000() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220713000000",
		Migrate: func(tx *gorm.DB) error {

			if err := tx.AutoMigrate(
				&model.Cafeteria{},
				&model.CafeteriaRating{},
				&model.CafeteriaRatingAverage{},
				&model.CafeteriaRatingTag{},
				&model.CafeteriaRatingTagsAverage{},
				&model.CafeteriaRatingTagOption{},
				&model.Dish{},
				&model.DishNameTagOption{},
				&model.DishNameTagOptionIncluded{},
				&model.DishNameTagOptionExcluded{},
				&model.DishNameTag{},
				&model.DishNameTagAverage{},
				&model.DishRating{},
				&model.DishRatingAverage{},
				&model.DishRatingTag{},
				&model.DishRatingTagAverage{},
				&model.DishRatingTagOption{},
				&model.DishToDishNameTag{},
			); err != nil {
				return err
			}
			return nil
		},
	}
}
