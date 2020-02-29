package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/zoover-reviews-api-go/models"
	"testing"
)

var reviewStats = &ReviewStats{
	Rating: struct {
		General float32
		Count   float32
	}{
		0,
		0,
	},
	Aspects: make(map[string]struct {
		Sum   float32
		Count float32
	}),
	TraveledWith: make(map[string]int),
}

var reviews = []models.Review{
	{
		TraveledWith: "FAMILY",
		EntryDate:    1252359116000,
		TravelDate:   1252359116000,
		Ratings: struct {
			General struct {
				General float32 `json:"general"`
			} `json:"general"`
			Aspects struct {
				Location        int `json:"location"`
				Service         int `json:"service"`
				PriceQuality    int `json:"priceQuality"`
				Food            int `json:"food"`
				Room            int `json:"room"`
				ChildFriendly   int `json:"childFriendly"`
				Interior        int `json:"interior"`
				Size            int `json:"size"`
				Activities      int `json:"activities"`
				Restaurants     int `json:"restaurants"`
				SanitaryState   int `json:"sanitaryState"`
				Accessibility   int `json:"accessibility"`
				Nightlife       int `json:"nightlife"`
				Culture         int `json:"culture"`
				Surrounding     int `json:"surrounding"`
				Atmosphere      int `json:"atmosphere"`
				NoviceSkiArea   int `json:"noviceSkiArea"`
				AdvancedSkiArea int `json:"advancedSkiArea"`
				ApresSki        int `json:"apresSki"`
				Beach           int `json:"beach"`
				Entertainment   int `json:"entertainment"`
				Environmental   int `json:"environmental"`
				Pool            int `json:"pool"`
				Terrace         int `json:"terrace"`
			} `json:"aspects"`
		}{
			General: struct {
				General float32 `json:"general"`
			}{
				General: 8,
			},
			Aspects: struct {
				Location        int `json:"location"`
				Service         int `json:"service"`
				PriceQuality    int `json:"priceQuality"`
				Food            int `json:"food"`
				Room            int `json:"room"`
				ChildFriendly   int `json:"childFriendly"`
				Interior        int `json:"interior"`
				Size            int `json:"size"`
				Activities      int `json:"activities"`
				Restaurants     int `json:"restaurants"`
				SanitaryState   int `json:"sanitaryState"`
				Accessibility   int `json:"accessibility"`
				Nightlife       int `json:"nightlife"`
				Culture         int `json:"culture"`
				Surrounding     int `json:"surrounding"`
				Atmosphere      int `json:"atmosphere"`
				NoviceSkiArea   int `json:"noviceSkiArea"`
				AdvancedSkiArea int `json:"advancedSkiArea"`
				ApresSki        int `json:"apresSki"`
				Beach           int `json:"beach"`
				Entertainment   int `json:"entertainment"`
				Environmental   int `json:"environmental"`
				Pool            int `json:"pool"`
				Terrace         int `json:"terrace"`
			}{
				Location:        9,
				Service:         0,
				PriceQuality:    9,
				Food:            0,
				Room:            0,
				ChildFriendly:   9,
				Interior:        0,
				Size:            0,
				Activities:      0,
				Restaurants:     0,
				SanitaryState:   0,
				Accessibility:   0,
				Nightlife:       0,
				Culture:         0,
				Surrounding:     0,
				Atmosphere:      0,
				NoviceSkiArea:   0,
				AdvancedSkiArea: 0,
				ApresSki:        0,
				Beach:           0,
				Entertainment:   0,
				Environmental:   0,
				Pool:            0,
				Terrace:         0,
			},
		},
	},
	{
		TraveledWith: "FAMILY",
		EntryDate:    1266252490713,
		TravelDate:   1262304000000,
		Ratings: struct {
			General struct {
				General float32 `json:"general"`
			} `json:"general"`
			Aspects struct {
				Location        int `json:"location"`
				Service         int `json:"service"`
				PriceQuality    int `json:"priceQuality"`
				Food            int `json:"food"`
				Room            int `json:"room"`
				ChildFriendly   int `json:"childFriendly"`
				Interior        int `json:"interior"`
				Size            int `json:"size"`
				Activities      int `json:"activities"`
				Restaurants     int `json:"restaurants"`
				SanitaryState   int `json:"sanitaryState"`
				Accessibility   int `json:"accessibility"`
				Nightlife       int `json:"nightlife"`
				Culture         int `json:"culture"`
				Surrounding     int `json:"surrounding"`
				Atmosphere      int `json:"atmosphere"`
				NoviceSkiArea   int `json:"noviceSkiArea"`
				AdvancedSkiArea int `json:"advancedSkiArea"`
				ApresSki        int `json:"apresSki"`
				Beach           int `json:"beach"`
				Entertainment   int `json:"entertainment"`
				Environmental   int `json:"environmental"`
				Pool            int `json:"pool"`
				Terrace         int `json:"terrace"`
			} `json:"aspects"`
		}{
			General: struct {
				General float32 `json:"general"`
			}{
				General: 7,
			},
			Aspects: struct {
				Location        int `json:"location"`
				Service         int `json:"service"`
				PriceQuality    int `json:"priceQuality"`
				Food            int `json:"food"`
				Room            int `json:"room"`
				ChildFriendly   int `json:"childFriendly"`
				Interior        int `json:"interior"`
				Size            int `json:"size"`
				Activities      int `json:"activities"`
				Restaurants     int `json:"restaurants"`
				SanitaryState   int `json:"sanitaryState"`
				Accessibility   int `json:"accessibility"`
				Nightlife       int `json:"nightlife"`
				Culture         int `json:"culture"`
				Surrounding     int `json:"surrounding"`
				Atmosphere      int `json:"atmosphere"`
				NoviceSkiArea   int `json:"noviceSkiArea"`
				AdvancedSkiArea int `json:"advancedSkiArea"`
				ApresSki        int `json:"apresSki"`
				Beach           int `json:"beach"`
				Entertainment   int `json:"entertainment"`
				Environmental   int `json:"environmental"`
				Pool            int `json:"pool"`
				Terrace         int `json:"terrace"`
			}{
				Location:        7,
				Service:         0,
				PriceQuality:    0,
				Food:            0,
				Room:            0,
				ChildFriendly:   0,
				Interior:        0,
				Size:            0,
				Activities:      0,
				Restaurants:     6,
				SanitaryState:   0,
				Accessibility:   0,
				Nightlife:       0,
				Culture:         0,
				Surrounding:     0,
				Atmosphere:      0,
				NoviceSkiArea:   0,
				AdvancedSkiArea: 0,
				ApresSki:        0,
				Beach:           0,
				Entertainment:   0,
				Environmental:   0,
				Pool:            0,
				Terrace:         0,
			},
		},
	},
	{
		TraveledWith: "OTHER",
		EntryDate:    1449970631852,
		TravelDate:   1246406400000,
		Ratings: struct {
			General struct {
				General float32 `json:"general"`
			} `json:"general"`
			Aspects struct {
				Location        int `json:"location"`
				Service         int `json:"service"`
				PriceQuality    int `json:"priceQuality"`
				Food            int `json:"food"`
				Room            int `json:"room"`
				ChildFriendly   int `json:"childFriendly"`
				Interior        int `json:"interior"`
				Size            int `json:"size"`
				Activities      int `json:"activities"`
				Restaurants     int `json:"restaurants"`
				SanitaryState   int `json:"sanitaryState"`
				Accessibility   int `json:"accessibility"`
				Nightlife       int `json:"nightlife"`
				Culture         int `json:"culture"`
				Surrounding     int `json:"surrounding"`
				Atmosphere      int `json:"atmosphere"`
				NoviceSkiArea   int `json:"noviceSkiArea"`
				AdvancedSkiArea int `json:"advancedSkiArea"`
				ApresSki        int `json:"apresSki"`
				Beach           int `json:"beach"`
				Entertainment   int `json:"entertainment"`
				Environmental   int `json:"environmental"`
				Pool            int `json:"pool"`
				Terrace         int `json:"terrace"`
			} `json:"aspects"`
		}{
			General: struct {
				General float32 `json:"general"`
			}{
				General: 9,
			},
			Aspects: struct {
				Location        int `json:"location"`
				Service         int `json:"service"`
				PriceQuality    int `json:"priceQuality"`
				Food            int `json:"food"`
				Room            int `json:"room"`
				ChildFriendly   int `json:"childFriendly"`
				Interior        int `json:"interior"`
				Size            int `json:"size"`
				Activities      int `json:"activities"`
				Restaurants     int `json:"restaurants"`
				SanitaryState   int `json:"sanitaryState"`
				Accessibility   int `json:"accessibility"`
				Nightlife       int `json:"nightlife"`
				Culture         int `json:"culture"`
				Surrounding     int `json:"surrounding"`
				Atmosphere      int `json:"atmosphere"`
				NoviceSkiArea   int `json:"noviceSkiArea"`
				AdvancedSkiArea int `json:"advancedSkiArea"`
				ApresSki        int `json:"apresSki"`
				Beach           int `json:"beach"`
				Entertainment   int `json:"entertainment"`
				Environmental   int `json:"environmental"`
				Pool            int `json:"pool"`
				Terrace         int `json:"terrace"`
			}{
				Location:        9,
				Service:         0,
				PriceQuality:    9,
				Food:            0,
				Room:            0,
				ChildFriendly:   10,
				Interior:        0,
				Size:            0,
				Activities:      0,
				Restaurants:     9,
				SanitaryState:   0,
				Accessibility:   0,
				Nightlife:       0,
				Culture:         0,
				Surrounding:     0,
				Atmosphere:      0,
				NoviceSkiArea:   0,
				AdvancedSkiArea: 0,
				ApresSki:        0,
				Beach:           0,
				Entertainment:   0,
				Environmental:   0,
				Pool:            0,
				Terrace:         0,
			},
		},
	},
}

var testTable = []struct {
	name string
	expected interface{}
} {
	{
		"should return correct average general rating",
		float64(4),
	},
	{
		"should return correct average aspect rating",
		map[string]float64{"accessibility": 0, "activities": 0, "advancedskiarea": 0, "apresski": 0, "atmosphere": 0, "beach": 0, "childfriendly": 3.17, "culture": 0, "entertainment": 0, "environmental": 0, "food": 0, "interior": 0, "location": 4.17, "nightlife": 0, "noviceskiarea": 0, "pool": 0, "pricequality": 3, "restaurants": 2.5, "room": 0, "sanitarystate": 0, "service": 0, "size": 0, "surrounding": 0, "terrace": 0},
	},
	{
		"should return traveledWith with count",
		map[string]int{"FAMILY": 2, "OTHER": 1},
	},
	{
		"should return stats response",
		&StatsResponse{AverageRating: 4, TraveledWith:map[string]int{"FAMILY": 2, "OTHER":1}, AverageRatingAspects:map[string]float64{"accessibility": 0, "activities":0, "advancedskiarea":0, "apresski":0, "atmosphere":0, "beach":0, "childfriendly":3.17, "culture":0, "entertainment":0, "environmental":0, "food":0, "interior":0, "location":4.17, "nightlife":0, "noviceskiarea":0, "pool":0, "pricequality":3, "restaurants":2.5, "room":0, "sanitarystate":0, "service":0, "size":0, "surrounding":0, "terrace":0}},
	},
}

func init() {
	for _, r := range reviews {
		reviewStats.Calculate(&r)
	}
}

func TestReviewStats_Calculate(t *testing.T) {
	assert.Equal(t, testTable[0].expected, reviewStats.getAverageGeneralRating(), testTable[0].name)
	assert.Equal(t, testTable[1].expected, reviewStats.getAverageAspectsRating(), testTable[1].name)
	assert.Equal(t, testTable[2].expected, reviewStats.getTraveledWith(), testTable[2].name)
}

func TestReviewStats_GetStats(t *testing.T) {
	assert.Equal(t, testTable[3].expected, reviewStats.GetStats(), testTable[3].name)
}