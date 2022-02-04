package controllers

import (
	"answers_app/constant"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"answers_app/dto"

	"answers_app/service"
)

type AnswerController struct{}

var answerService service.AnswerService

func (a AnswerController) CreateAnswer(c *gin.Context) {
	var requestBody dto.AnswerRequest

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Errorf("Error in request body", err)
	}

	status := answerService.CreateAnswer(requestBody)

	if status == true {
		c.JSON(http.StatusCreated, gin.H{
			constant.RESP_MSG:    constant.ANSWER_CREATED,
			constant.RESP_STATUS: http.StatusCreated,
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			constant.RESP_MSG:    constant.ANSWER_EXISTS,
			constant.RESP_STATUS: http.StatusConflict,
		})
	}
}

func (a AnswerController) UpdateAnswer(c *gin.Context) {
	var requestBody dto.AnswerRequest

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Errorf("Error in request body", err)
	}

	status := answerService.UpdateAnswer(requestBody)

	if status == false {
		c.JSON(http.StatusNotFound, gin.H{
			constant.RESP_MSG:    constant.ANSWER_NOTEXISTS,
			constant.RESP_STATUS: http.StatusNotFound,
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			constant.RESP_MSG:    constant.ANSWER_UPDATED,
			constant.RESP_STATUS: http.StatusAccepted,
		})
	}
}

func (a AnswerController) GetAnswer(c *gin.Context) {

	key := c.Param("key")
	status := answerService.GetAnswer(key)

	if status.Key == "" {
		c.JSON(http.StatusNoContent, gin.H{
			constant.RESP_MSG:    constant.ANSWER_NOTEXISTS,
			constant.RESP_STATUS: http.StatusNoContent,
		})
	} else {
		c.JSON(http.StatusOK, status)
	}
}

func (a AnswerController) DeleteAnswer(c *gin.Context) {
	key := c.Param("key")
	status := answerService.DeleteAnswer(key)

	if status == false {
		c.JSON(http.StatusNoContent, gin.H{
			constant.RESP_MSG:    constant.ANSWER_NOTEXISTS,
			constant.RESP_STATUS: http.StatusNoContent,
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			constant.RESP_MSG:    constant.ANSWER_DELETED,
			constant.RESP_STATUS: http.StatusAccepted,
		})
	}
}

func (a AnswerController) AnswerHistory(c *gin.Context) {
	key := c.Param("key")
	resp := answerService.AnswerHistory(key)

	if len(resp) == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			constant.RESP_MSG:    constant.ANSWER_NOTEXISTS,
			constant.RESP_STATUS: http.StatusNoContent,
		})
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
