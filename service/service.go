package service

import (
	"answers_app/constant"
	"answers_app/dto"
	"answers_app/model"
	"answers_app/repository"
)

type AnswerService struct{}

var answerRepo repository.AnswerRepository

func (a AnswerService) CreateAnswer(answer dto.AnswerRequest) bool {
	isExists := answerRepo.IsAnswerExists(answer.Key)

	if isExists == 0 {
		event := model.EventModel{Name: constant.CREATE}
		dataObj := model.AnswerModel{AnswerKey: answer.Key, AnswerValue: answer.Value, Event: event, IsDeleted: false}
		answerRepo.SaveAnswerEvent(&dataObj)
		return true
	} else {
		return false
	}
}

func (a AnswerService) UpdateAnswer(answer dto.AnswerRequest) bool {
	isExists := answerRepo.IsAnswerExists(answer.Key)

	if isExists == 1 {
		event := model.EventModel{Name: constant.UPDATE}
		dataObj := model.AnswerModel{AnswerKey: answer.Key, AnswerValue: answer.Value, Event: event, IsDeleted: false}
		answerRepo.SaveAnswerEvent(&dataObj)
		return true
	} else {
		return false
	}
}

func (a AnswerService) GetAnswer(key string) dto.Answer {
	var checkAnswer model.AnswerModel
	isExists := answerRepo.IsAnswerExists(key)

	if isExists == 1 {
		result := answerRepo.GetAnswerByKey(key, checkAnswer)
		answer := dto.Answer{Key: result.AnswerKey, Value: result.AnswerValue}
		return answer
	} else {
		var answer dto.Answer
		return answer
	}

}

func (a AnswerService) DeleteAnswer(key string) bool {
	isExists := answerRepo.IsAnswerExists(key)

	if isExists == 0 {
		return false
	} else {
		event := model.EventModel{Name: constant.DELETE}
		dataObj := model.AnswerModel{AnswerKey: key, AnswerValue: constant.EMPTY, Event: event, IsDeleted: true}
		answerRepo.SaveAnswerEvent(&dataObj)
		answerRepo.DeleteAnswers(key)
		return true
	}
}

func (a AnswerService) AnswerHistory(key string) []dto.Event {
	var history []dto.Event
	isExists := answerRepo.IsHistoryExists(key)

	if isExists == 0 {
		return history
	} else {
		result := answerRepo.GetAnswersHistory(key)
		for _, ans := range result {
			answerObj := dto.Answer{Key: ans.AnswerKey, Value: ans.AnswerValue}
			eventObj := dto.Event{Event: ans.Event.Name, Data: answerObj}
			history = append(history, eventObj)
		}
		return history
	}
}
