package dao

import (
	"database/sql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type dbConfig struct {
	host     string `yaml:"host"`
	dbname   string `yaml:"dbname"`
	username string `yaml:"username"`
	password string `yaml:"password"`
}

func NewDB() *sql.DB {
	var dbc dbConfig
	dbc.loadYaml()
	dsn := dbc.username + ":" +
		dbc.password + "@tcp(" + dbc.host + ":3306)/" +
		dbc.dbname + ")?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func (dbc *dbConfig) loadYaml() *dbConfig {
	yamlFile, err := ioutil.ReadFile("../../config/conf.yml")
	err = yaml.Unmarshal(yamlFile, dbc)
	if err != nil {
		panic(err)
	}
	return dbc
}
