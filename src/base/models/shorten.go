package models

import (
//"fmt"
)

type MetaParams struct {
	Source    string
	Version   string
	SecretKey string
	RequestID string
}

type DataParams struct {
	LongUrl string
	Ip      string
}

type Shorten struct {
	Meta MetaParams
	Data DataParams
}

func GetParams(shorten Shorten) Shorten {
	return shorten
}
