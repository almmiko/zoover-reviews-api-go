package services

import (
	"encoding/json"
	"github.com/zoover-reviews-api-go/models"
	"github.com/zoover-reviews-api-go/utils"
	"log"
	"os"
	"sort"
	"strings"
)

type Reviews struct {
	Data []models.Review
	Stats *utils.ReviewStats
}

type Meta struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	TotalItems int `json:"totalItems"`
	HasNext bool `json:"hasNext"`
}

func NewReviewsService() *Reviews {
	return &Reviews{}
}

func (reviews *Reviews) Load() {
	reviews.Stats = &utils.ReviewStats{
		Aspects:      make(map[string] struct{ Sum float32; Count float32}),
		TraveledWith: make(map[string] int),
	}

	reviewsJSON, err := os.Open("data/reviews.json")
	if err != nil {
		log.Fatal("error reading from reviews.json")
	}

	reviewsParser := json.NewDecoder(reviewsJSON)

	if err := reviewsParser.Decode(&reviews.Data); err != nil {
		log.Fatal("error parsing reviews")
	}

	defer reviewsJSON.Close()

	for _, review := range reviews.Data {
		reviews.Stats.Calculate(&review)
	}
}

func (reviews *Reviews) GetStats() *utils.StatsResponse {
	return reviews.Stats.GetStats()
}

func (reviews *Reviews) FilterByTraveledWith(value string) []models.Review {
	filter := strings.ToUpper(value)
	filters := []string{"FAMILY", "FRIENDS", "OTHER", "COUPLE", "SINGLE"}

	if !utils.Contains(filters, filter) {
		return []models.Review{}
	}

	b := reviews.getCopy()[:0]
	for _, x := range reviews.getCopy() {
		if x.TraveledWith == filter {
			b = append(b, x)
		}
	}

	return b
}

func (reviews *Reviews) GetSortedCollection(predicate string, order string, collection []models.Review) []models.Review {
	if collection == nil {
		collection = reviews.getCopy()
	}

	if predicate == "entryDate" {
		return reviews.sortByEntryDate(order, collection)
	}

	if predicate == "travelDate" {
		return reviews.sortByTravelDate(order, collection)
	}

	return reviews.sortByEntryDate(order, collection)
}

func (reviews *Reviews) Paginate(page int, limit int, collection []models.Review) (*Meta, []models.Review) {
	offset := (page - 1) * limit

	if offset > len(collection) { offset = 0 }

	sliced := collection[offset:]

	if limit > len(sliced) { limit = len(sliced)}

	paginatedCollection := sliced[0: limit]

	if len(paginatedCollection) != 0 {
		return &Meta{
			Page:       page,
			Limit:      limit,
			TotalItems: len(collection),
			HasNext:    true,
		}, paginatedCollection
	} else {
		return &Meta{
			HasNext:    false,
		}, nil
	}
}

func (reviews *Reviews) GetReviews() []models.Review {
	return reviews.Data
}

func (reviews *Reviews) sortByEntryDate(order string, collection []models.Review) []models.Review {
	if order == "asc" {
		sort.SliceStable(collection, func(i, j int) bool {
			return collection[i].EntryDate < collection[j].EntryDate
		})
	} else {
		sort.SliceStable(collection, func(i, j int) bool {
			return collection[i].EntryDate > collection[j].EntryDate
		})
	}

	return collection
}

func (reviews *Reviews) sortByTravelDate(order string, collection []models.Review) []models.Review {
	if order == "asc" {
		sort.Slice(collection, func(i, j int) bool {
			return collection[i].TravelDate < collection[j].TravelDate
		})
	} else {
		sort.Slice(collection, func(i, j int) bool {
			return collection[i].TravelDate > collection[j].TravelDate
		})
	}

	return collection
}

func (reviews *Reviews) getCopy() []models.Review {
	reviewsCopy := make([]models.Review, len(reviews.Data))
	copy(reviewsCopy, reviews.Data)

	return reviewsCopy
}