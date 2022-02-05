package main

import (
	"answers_app/dto"
	"answers_app/service"
	"math/rand"
	"testing"
	"time"
)

var answerService service.AnswerService

var key string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	rand.Seed(int64(time.Now().Minute()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestCreateAnswer(t *testing.T) {
	key := randSeq(10)
	answer := dto.AnswerRequest{Key: key, Value: "value_1"}
	status := answerService.CreateAnswer(answer)

	if status == false {
		t.Errorf("FAILED : Unable to create answer ! Either key alreay exists or something went wrong.")
	}
}

func TestUpdateAnswer(t *testing.T) {
	key := randSeq(10)
	answer := dto.AnswerRequest{Key: key, Value: "value_2"}
	status := answerService.UpdateAnswer(answer)

	if status == false {
		t.Errorf("FAILED : Unable to update answer! Either answer key doesn't exists or something went wrong.")
	}
}

func TestGetAnswer(t *testing.T) {
	key := randSeq(10)
	status := answerService.GetAnswer(key)

	if status.Key == "" {
		t.Errorf("FAILED : Unable to get answer! Either answer key doesn't exists or something went wrong.")
	}

}

func TestDeleteAnswer(t *testing.T) {
	key := randSeq(10)
	status := answerService.DeleteAnswer(key)
	if status == false {
		t.Errorf("FAILED : Unable to delete answer! Either answer key doesn't exists or something went wrong.")
	}
}

func TestAnswerHistory(t *testing.T) {
	key := randSeq(10)
	resp := answerService.AnswerHistory(key)

	if len(resp) == 0 {
		t.Errorf("FAILED : Unable to get answer history ! Either answer key doesn't exists or something went wrong.")
	}
}
