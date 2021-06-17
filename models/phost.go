package models

import (
	_ "github.com/jinzhu/gorm"
)

type Phost struct {
	Id           int
	Sn           string
	EntryName    string
	Name         string
	Environment  int
	Status       int
	Model        string
	Position     string
	RackPosition string
	Intranet     string
	AddTime      int
}

func (Phost) TableName() string {
	return "phost"
}
