package models

import (
	_ "github.com/jinzhu/gorm"
)

type DomainNameList struct {
	Id                int
	DomainName        string
	Intranet          string
	ResolveRecordType int
	CameValue         string
	EntryName         string
	WafLabel          int
	LbLabel           int
	AddTime           int
}

func (DomainNameList) TableName() string {
	return "domainnamelist"
}
