package models

import (
	_ "github.com/jinzhu/gorm"
)

type Host struct {
	Id          int
	Sn          string
	EntryName   string
	Name        string
	Type        int
	Environment int
	Status      int
	Position    string
	Intranet    string
	External    string
	AddTime     int
}

func (Host) TableName() string {
	return "host"
}
