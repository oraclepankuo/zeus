package models

import (
	_ "github.com/jinzhu/gorm"
)

type NetworkDevice struct {
	Id            int
	Sn            string
	InstanceName  string
	Model         string
	Position      string
	RackPosition  string
	Intranet      string
	Type          int
	MachineStatus int
	BusinessType  string
	AddTime       int
}

func (NetworkDevice) TableName() string {
	return "networkdevice"
}
