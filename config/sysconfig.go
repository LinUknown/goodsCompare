package config

import (
	"encoding/json"
	"io/ioutil"
)

type sysconfig struct {
	Port       string    `json:"Port"`
	DBUserName   string `json:"DBUserName"`
	DBPassword string    `json:"DBPassword"`
	DBIp   string    `json:"DBIp"`
	DBPort    string    `json:"DBPort"`
	DBName string    `json:"DBName"`

}
var Sysconfig = &sysconfig{}

func init(){
	b,err := ioutil.ReadFile("./config.json")
	if err != nil{
		panic("sys conf read err")
	}
	err = json.Unmarshal(b,Sysconfig)
	if err != nil{
		panic(err)
	}
}
