package main

import (
	"encoding/json"
	"io/ioutil"
)


type ConfAwsDB struct {
	User		string `json:"User"`
	Password	string `json:"Password"`
	Endpoint	string `json:"Endpoint"`
	DBName		string `json:"DBName"`
	Region		string `json:"Region"`
}

type ConfRSA struct {
	PathPrivate		string `json:"PathPrivate"`
	PathPublic		string `json:"PathPublic"`
}

type Conf struct {
	AwsDB	ConfAwsDB	`json:"AwsDB"`
	RSA		ConfRSA		`json:"RSA"`
}

func loadConf() (Conf) {
	var conf Conf

	data, err := ioutil.ReadFile("conf/conf.json")
	fatal(err)

	err = json.Unmarshal(data, &conf)
	fatal(err)
	return conf
}