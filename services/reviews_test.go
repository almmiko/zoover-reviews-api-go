package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/zoover-reviews-api-go/models"
	"github.com/zoover-reviews-api-go/utils"
	"testing"
)

var reviews = []models.Review{
	{
		ID: "1",
		TraveledWith: "FAMILY",
		EntryDate:    1252359116000, //Monday, 7 September 2009 21:31:56
		TravelDate:   1252359116000,
	},
	{
		ID: "2",
		TraveledWith: "FAMILY",
		EntryDate:    1266252490713, //Monday, 15 February 2010 16:48:10.713
		TravelDate:   1262304000000, //Friday, 1 January 2010 00:00:00
	},
	{
		ID: "3",
		TraveledWith: "OTHER",
		EntryDate:    1449970631852, //Sunday, 13 December 2015 01:37:11.852
		TravelDate:   1246406400000, // Wednesday, 1 July 2009 00:00:00
	},
}

var filterTestCases = []struct {
	input    string
	expected int
}{
	{
		"OTHER",
		1,
	},
	{
		"FAMILY",
		2,
	},
	{
		"unknown",
		0,
	},
}

var sortingTestCases = []struct {
	predicate string
	order string
	expected []string
	collection []models.Review
}{
	{
		"entryDate",
		"asc",
		[]string{"1", "2", "3"},
		reviews,
	},
	{
		"entryDate",
		"desc",
		[]string{"3", "2", "1"},
		reviews,
	},
	{
		"travelDate",
		"asc",
		[]string{"3", "1", "2"},
		reviews,
	},
	{
		"travelDate",
		"desc",
		[]string{"2", "1", "3"},
		reviews,
	},
	{
		"travelDate",
		"desc",
		[]string{"2", "1", "3"},
		nil,
	},
	{
		"",
		"",
		[]string{"3", "2", "1"},
		nil,
	},
}

var paginationTestCases = []struct{
	page int
	limit int
	expected struct{
		meta *Meta
		collectionLen int
	}
	collection []models.Review
} {
	{
		1,
		1,
		struct {
			meta       *Meta
			collectionLen int
		}{meta: &Meta{
			Page:       1,
			Limit:      1,
			TotalItems: 3,
			HasNext:    true,
		}, collectionLen: 1 },
		reviews,
	},
	{
		4,
		1,
		struct {
			meta       *Meta
			collectionLen int
		}{meta: &Meta{
			HasNext:    false,
		}, collectionLen: 0 },
		reviews,
	},
}

func TestReviews_FilterByTraveledWith(t *testing.T) {
	reviews := &Reviews{
		Data: reviews,
	}

	for _, tt := range filterTestCases {
		t.Run("should filter using traveledWith value", func(t *testing.T) {
			got := reviews.FilterByTraveledWith(tt.input)
			assert.Equal(t, tt.expected,  len(got))
		})
	}
}

func TestReviews_GetSortedCollection(t *testing.T) {
	reviews := &Reviews{
		Data: reviews,
	}

	for _, tt := range sortingTestCases {
		t.Run("should return collection sorted by date", func(t *testing.T) {
			got := reviews.GetSortedCollection(tt.predicate, tt.order, tt.collection)

			var reviewsIds []string

			for _, review := range got {
				reviewsIds = append(reviewsIds, review.ID)
			}

			assert.Equal(t, tt.expected,  reviewsIds)
		})
	}
}

func TestReviews_Paginate(t *testing.T) {
	reviews := &Reviews{
		Data: reviews,
	}

	for _, tt := range paginationTestCases {
		t.Run("should return pagination", func(t *testing.T) {
			meta, collection := reviews.Paginate(tt.page, tt.limit, tt.collection)

			assert.Equal(t, tt.expected.meta,  meta)
			assert.Equal(t, tt.expected.collectionLen,  len(collection))
		})
	}

}

func TestNewReviewsService(t *testing.T) {
	expected := &Reviews{}
	actual := NewReviewsService()
	assert.Equal(t, expected, actual)
}

func TestReviews_GetStats(t *testing.T) {
	reviews := &Reviews{
		Data: reviews,
	}

	reviews.Stats = &utils.ReviewStats{
		Aspects:      make(map[string] struct{ Sum float32; Count float32}),
		TraveledWith: make(map[string] int),
	}

	expected := &utils.StatsResponse{}
	actual := reviews.GetStats()

	assert.IsType(t, expected, actual)
}

func TestReviews_GetReviews(t *testing.T) {
	reviews := &Reviews{
		Data: reviews,
	}

	expected := reviews.Data
	actual := reviews.GetReviews()
	assert.Equal(t, expected, actual)
	assert.IsType(t, expected, actual)
}