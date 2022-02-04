package routes

import (
	"github.com/gin-gonic/gin"

	"answers_app/controllers"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	answer := new(controllers.AnswerController)
	router.PUT("/answer", answer.CreateAnswer)
	router.PATCH("/answer", answer.UpdateAnswer)
	router.GET("/answer/:key", answer.GetAnswer)
	router.DELETE("/answer/:key", answer.DeleteAnswer)
	router.GET("/answer/:key/history", answer.AnswerHistory)

	return router

}
