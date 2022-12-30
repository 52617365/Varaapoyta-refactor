package endpoints

import (
	"net/http"
	"strings"
	"varaapoyta-backend-refactor/requests"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func InitApi() {
	router := gin.Default()
	InitEndpoints(router)
	router.Run()
}

func InitEndpoints(router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	setCorsRules(router)
	router.GET("/tables/:city", func(c *gin.Context) {
		city := c.Param("city")
		if !cityIsValid(city) {
			c.JSON(http.StatusBadRequest, "Provided city does not exist on the Raflaamo page.")
			return
		}
		restaurantsWithTimeSlots, err := requests.GetRestaurantsWithTimeSlots(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, restaurantsWithTimeSlots)
	})
}

func setCorsRules(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"https://raflaamo.rasmusmaki.com"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))
}

func cityIsValid(city string) bool {
	return cityIsOnRaflaamoList(city)
}

func cityIsOnRaflaamoList(city string) bool {
	return slices.Contains(allPossibleCities, strings.ToLower(city))
}
