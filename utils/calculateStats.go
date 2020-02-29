package utils

import (
	"github.com/zoover-reviews-api-go/models"
	"math"
	"reflect"
	"strings"
)

type ReviewStats struct {
	Rating struct {
		General float32
		Count float32
	}
	Aspects map[string] struct{
		Sum float32
		Count float32
	}
	TraveledWith map[string] int
}

type StatsResponse struct {
	AverageRating float64 `json:"averageRating"`
	TraveledWith map[string] int `json:"traveledWith"`
	AverageRatingAspects map[string]float64 `json:"averageRatingAspects"`
}

func (stats *ReviewStats) calculateGeneralRatingSum(review *models.Review) {
	ratings := review.Ratings
	reviewWeight := review.ReviewWeight()

	stats.Rating.General += reviewWeight * ratings.General.General
	stats.Rating.Count += 1
}

func (stats *ReviewStats) calculateRatingAspectsSum(review *models.Review) {
	v := reflect.ValueOf(review.Ratings.Aspects)
	aspectsType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(aspectsType.Field(i).Name)
		value := v.Field(i).Interface().(int)

		_, ok := stats.Aspects[key]
		if !ok {
			stats.Aspects[key] = struct {
				Sum   float32
				Count float32
			}{Sum: 0, Count: 0}
		}

		sum := stats.Aspects[key].Sum
		count := stats.Aspects[key].Count

		stats.Aspects[key] = struct {
			Sum   float32
			Count float32
		}{Sum:
			sum + (review.ReviewWeight() * float32(value)),
			Count: count + 1,
		}

	}

}

func (stats *ReviewStats) calculateTraveledWith(review *models.Review) {
	_, ok := stats.TraveledWith[review.TraveledWith]
	if !ok {
		stats.TraveledWith[review.TraveledWith] = 0
	}

	stats.TraveledWith[review.TraveledWith] += 1
}

func (stats *ReviewStats) getAverageGeneralRating() float64 {
	return math.Floor(float64(stats.Rating.General / stats.Rating.Count) * 100) / 100
}

func (stats *ReviewStats) getTraveledWith() map[string]int {
	return stats.TraveledWith
}

func (stats * ReviewStats) getAverageAspectsRating() map[string]float64 {
	average := make(map[string]float64)

	for key, value := range stats.Aspects {
		average[key] = math.Round(float64((value.Sum/value.Count)*100)) / 100
	}

	return average
}

func (stats *ReviewStats) Calculate(review *models.Review) {
	stats.calculateGeneralRatingSum(review)
	stats.calculateTraveledWith(review)
	stats.calculateRatingAspectsSum(review)
}

func (stats *ReviewStats) GetStats() *StatsResponse {
	return &StatsResponse{
		AverageRating: stats.getAverageGeneralRating(),
		AverageRatingAspects: stats.getAverageAspectsRating(),
		TraveledWith:stats.getTraveledWith(),
	}
}
