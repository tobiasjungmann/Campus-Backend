package cron

import (
	"github.com/TUM-Dev/Campus-Backend/model/cafeteria_rating_models"
	"log"
)

type averageRatingForCafeteria struct {
	Cafeteria string  `json:"cafeteria"`
	Average   float32 `json:"average"`
	Min       int     `json:"min"`
	Max       int     `json:"max"`
}

type averageRatingForMealInCafeteria struct {
	Cafeteria string  `json:"cafeteria"`
	Meal      string  `json:"meal"`
	Average   float32 `json:"average"`
	Min       int     `json:"min"`
	Max       int     `json:"max"`
}

//regularly computes the average rating for every cafeteria
func (c *CronService) averageRatingComputation() error {

	computeAverageForCafeteria(c)
	computeAverageForMealsInCafeterias(c)
	computeAverageCafeteriaTags(c)
	return nil
}

func computeAverageCafeteriaTags(c *CronService) {

	/*Todo
				alle ratings einer cafeteria sammeln -> nur die IDs merken, diese dann auf das tagrating tabelle anwenden - join
				-> alle tagratings zu einer cafeteria in einer Gruppe
				-> pro gruppe nach den tags gruppieren und den Durchschnitt bereschnen und in einemr result tabelle Speichern


			für alle drei tagarten berechnen
		nameratingtag ist das komplizierte, die beiden anderen können auf der tagrating tabelle bestimmt werden
	-> zusammenführen durch das parent rating um zu erfahren, zu welcher mensa die gerichte gehören
	*/

	/*
		err := s.db.Raw("SELECT r.*, a.campus, a.name "+
				"FROM roomfinder_rooms r "+
				"LEFT JOIN roomfinder_building2area a ON a.building_nr = r.building_nr "+
				"WHERE MATCH(room_code, info, address) AGAINST(?)", req.Query).Scan(&res).Error
	*/

	//nach der tagID gruppieren

}

func computeAverageForMealsInCafeterias(c *CronService) {
	var results []averageRatingForMealInCafeteria
	res := c.db.Model(cafeteria_rating_models.MealRating{}).
		Select("cafeteria, meal, AVG(rating) as average, MAX(rating) as max, MIN(rating) as min").
		Group("cafeteria,meal").Find(&results)

	if res.Error != nil {
		log.Println("Error in query")
		log.Println(res.Error)
	} else {
		for _, v := range results {
			cafeteria := cafeteria_rating_models.MealRatingsAverage{
				Cafeteria: v.Cafeteria,
				Average:   float32(v.Average),
				Meal:      v.Meal,
				Min:       v.Min,
				Max:       v.Max,
			} //todo add standard deviation

			var existing *cafeteria_rating_models.MealRatingsAverage
			testDish := c.db.Model(cafeteria_rating_models.MealRatingsAverage{}).
				Where("cafeteria = ?", cafeteria.Cafeteria).
				Where("meal = ?", cafeteria.Meal).
				First(&existing)

			if testDish.RowsAffected == 1 {
				errUpdate := c.db.Model(&cafeteria_rating_models.MealRatingsAverage{}).
					Where("cafeteria = ?", cafeteria.Cafeteria).
					Where("meal = ?", cafeteria.Meal).
					Updates(cafeteria)

				if errUpdate.Error != nil {
					log.Println(errUpdate.Error)
				}
			} else {
				log.Println("New average rating will be created for cafeteria: ", v.Cafeteria)
				errCreate := c.db.Create(&cafeteria)
				if errCreate.Error != nil {
					log.Println(errCreate.Error)
				}
			}
		}
	}
}

func computeAverageForCafeteria(c *CronService) {
	var results []averageRatingForCafeteria
	res := c.db.Model(cafeteria_rating_models.CafeteriaRating{}).
		Select("cafeteria, AVG(rating) as average, MAX(rating) as max, MIN(rating) as min").
		Group("cafeteria").Find(&results)

	if res.Error != nil {
		log.Println("Error in query")
		log.Println(res.Error)
	} else {
		for _, v := range results {
			cafeteria := cafeteria_rating_models.CafeteriaRatingResult{
				Cafeteria: v.Cafeteria,
				Average:   v.Average,
				Min:       v.Min,
				Max:       v.Max,
			} //todo add standard deviation

			var existing *cafeteria_rating_models.CafeteriaRatingResult
			testDish := c.db.Model(cafeteria_rating_models.CafeteriaRatingResult{}).Where("cafeteria = ?", cafeteria.Cafeteria).First(&existing)

			if testDish.RowsAffected == 1 {
				errUpdate := c.db.Model(&cafeteria_rating_models.CafeteriaRatingResult{}).
					Where("cafeteria = ?", cafeteria.Cafeteria).
					Updates(cafeteria)

				if errUpdate.Error != nil {
					log.Println(errUpdate.Error)
				}
			} else {
				log.Println("New rating will be created for cafeteria: ", v.Cafeteria)
				errCreate := c.db.Create(&cafeteria)
				if errCreate.Error != nil {
					log.Println(errCreate.Error)
				}
			}
		}
	}
}
