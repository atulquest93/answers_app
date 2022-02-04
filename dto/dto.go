package dto

type AnswerRequest struct {
	Key   string
	Value string
}

type Event struct {
	Event string
	Data  Answer
}

type Answer struct {
	Key   string
	Value string
}
