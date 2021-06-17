package models

import (
	_ "github.com/jinzhu/gorm"
)

type CertificateList struct {
	Id              int
	CertificateName string
	DomainName      string
	Status          int
	PersonInCharge  string
	Describe        string
	ExpirationTime  int
	AddTime         int
}

func (CertificateList) TableName() string {
	return "certificatelist"
}
