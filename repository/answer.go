package repository

import (
	"answers_app/database"

	"answers_app/model"

	"gorm.io/gorm"
)

var connection *gorm.DB

type AnswerRepository struct{}

func (a AnswerRepository) IsAnswerExists(answer_key string) int64 {

	if connection == nil {
		connection = database.Open()
	}

	var checkExistance model.AnswerModel
	isExists := connection.First(&checkExistance, "answer_key = ? AND is_deleted = false", answer_key)
	return isExists.RowsAffected

}

func (a AnswerRepository) IsHistoryExists(answer_key string) int64 {

	if connection == nil {
		connection = database.Open()
	}

	var checkExistance model.AnswerModel
	isExists := connection.First(&checkExistance, "answer_key = ?", answer_key)
	return isExists.RowsAffected

}

func (a AnswerRepository) SaveAnswerEvent(answer *model.AnswerModel) int64 {

	if connection == nil {
		connection = database.Open()
	}
	result := connection.Create(&answer)
	return result.RowsAffected

}

func (a AnswerRepository) GetAnswerByKey(answerkey string, answer model.AnswerModel) model.AnswerModel {

	if connection == nil {
		connection = database.Open()
	}
	connection.Last(&answer, "answer_key = ? AND is_deleted = false", answerkey)
	return answer
}

func (a AnswerRepository) DeleteAnswers(key string) int64 {
	if connection == nil {
		connection = database.Open()
	}
	var deletetItem model.AnswerModel
	result := connection.Model(&deletetItem).Where("answer_key = ?", key).Update("is_deleted", true)
	return result.RowsAffected
}

func (a AnswerRepository) GetAnswersHistory(key string) []model.AnswerModel {
	if connection == nil {
		connection = database.Open()
	}

	var answer []model.AnswerModel
	connection.Preload("Event").Find(&answer, "answer_key = ?", key)

	return answer
}
