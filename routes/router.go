package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zoover-reviews-api-go/controllers"
	"github.com/zoover-reviews-api-go/services"
)

type Srv []struct{
	Name string
	Service interface{}
}

type Controllers []struct{
	Name string
	Controller interface{}
}

type Router struct {
	App *gin.Engine

	services Srv
	controllers Controllers
}

func NewRouter(app *gin.Engine) *Router {
	return &Router{App: app}
}

func (r *Router) RegisterAll() *Router {
	reviewsService := services.NewReviewsService()
	reviewsService.Load()
	reviewsController := controllers.NewReviewsController(reviewsService)

	r.App.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	r.App.GET("/v1/reviews", reviewsController.GetReviews)
	r.App.GET("/v1/reviews/stats", reviewsController.GetReviewsStats)

	return r
}