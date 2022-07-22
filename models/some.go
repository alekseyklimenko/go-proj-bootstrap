package models

import "gorm.io/gorm"

const (
	SomeStatusIdle       = 0
	SomeStatusProcessing = 1
	SomeStatusError      = 2
)

type Some struct {
	gorm.Model
	Name       string `json:"name"`
	Url        string `json:"url"`
	ClientId   int    `json:"client_id"`
	Status     int    `json:"status" gorm:"index:idx_status"`
	StatusText string `json:"status_text"`
}

func NewSome() *Some {
	return &Some{}
}

func (s *Some) TableName() string {
	return "some_table"
}

func (s *Some) ValidationErrorMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "correctUrl":
		return "URL are incorrect. Must be in format ..."
	}
	return ""
}
