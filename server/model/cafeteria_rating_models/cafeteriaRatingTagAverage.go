package cafeteria_rating_models

import (
	"database/sql"
	"github.com/guregu/null"
)

var (
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// CafeteriaRatingTagsAverage stores all precomputed values for the cafeteria ratings
type CafeteriaRatingTagsAverage struct {
	CafeteriaRatingTagsAverage int32   `gorm:"primary_key;AUTO_INCREMENT;column:cafeteriaRatingTagsAverage;type:int;" json:"cafeteriaRatingTagsAverage" `
	CafeteriaID                int32   `gorm:"column:cafeteriaID;foreignKey:cafeteria;type:int;" json:"cafeteriaID"`
	TagID                      int32   `gorm:"column:tagID;foreignKey:cafeteriaRatingTagOption;type:int;" json:"tagID"`
	Average                    float32 `gorm:"column:average;type:float;" json:"average" `
	Min                        int8    `gorm:"column:min;type:int;" json:"min"`
	Max                        int8    `gorm:"column:max;type:int;" json:"max"`
	Std                        float32 `gorm:"column:std;type:float;" json:"std"`
}

// TableName sets the insert table name for this struct type
func (n *CafeteriaRatingTagsAverage) TableName() string {
	return "cafeteria_rating_tag_average"
}
