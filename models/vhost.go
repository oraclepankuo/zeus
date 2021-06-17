package models

import (
	_ "github.com/jinzhu/gorm"
)

type Vhost struct {
	Id                   int
	Sn                   string
	InstanceName         string
	Environment          int
	MachineStatus        int
	ResourcePool         int
	VirtualMachineType   int
	VirtualMachineStatus int
	Position             string
	Intranet             string
	IdleRate             int
	AddTime              int
}

func (Vhost) TableName() string {
	return "vhost"
}
