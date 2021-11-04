package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Db *gorm.DB

type Config struct {
	Mysql struct {
		Dsn      string `yaml:"dsn"`
		MaxOpen  int    `yaml:"max_open"`
		MaxIdle  int    `yaml:"max_idle"`
		LifeTime string `yaml:"life_time"`
		Log      bool   `yaml:"log"`
	}
}

func init() {
	dsn := "root:wang0805@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db = db
}

func Mysql() {
	dbConfig := Config{}

	data, err := ioutil.ReadFile("src/config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &dbConfig)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", dbConfig.Mysql.Dsn)
	if err != nil {
		panic(err)
	}
	Db = db
}