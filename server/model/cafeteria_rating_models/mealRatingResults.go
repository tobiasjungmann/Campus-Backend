package cafeteria_rating_models

import (
	"database/sql"
	"github.com/guregu/null"
)

var (
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// News struct is a row record of the mensa table in the tca database
type MealRatingResult struct {
	Id        int32   `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id" `
	Cafeteria string  `gorm:"column:cafeteria;type:mediumtext;" json:"cafeteria"`
	Meal      string  `gorm:"column:meal;type:mediumtext;" json:"meal"`
	Average   float32 `gorm:"column:latitude;type:float;" json:"latitude" `
	Min       int     `gorm:"column:min;type:int;" json:"min"`
	Max       int     `gorm:"column:max;type:int;" json:"max"`
}

// TableName sets the insert table name for this struct type
func (n *MealRatingResult) TableName() string {
	return "meal_rating_results"
}
