package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoover-reviews-api-go/models"
	"github.com/zoover-reviews-api-go/services"
	"net/http"
	"strconv"
)

type ReviewsController struct {
	Service *services.Reviews
}

func NewReviewsController(service *services.Reviews) *ReviewsController {
	return &ReviewsController{
		Service: service,
	}
}

func (rc *ReviewsController) GetReviews(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "entryDate")
	order := c.DefaultQuery("order", "desc")

	pageQuery := c.DefaultQuery("page", "1")
	limitQuery := c.DefaultQuery("limit", "20")

	page, _ := strconv.Atoi(pageQuery)
	limit, _ := strconv.Atoi(limitQuery)

	var reviews []models.Review

	if traveledWith := c.Query("traveledWith"); traveledWith != "" {
		reviews = rc.Service.FilterByTraveledWith(traveledWith)
	}

	reviews = rc.Service.GetSortedCollection(sortBy, order, reviews)

	meta, reviews := rc.Service.Paginate(page, limit, reviews)

	c.JSON(http.StatusOK, gin.H{
		"resources": reviews,
		"meta": meta,
	})
}

func (rc *ReviewsController) GetReviewsStats(c *gin.Context) {
	c.JSON(http.StatusOK, rc.Service.GetStats())
}