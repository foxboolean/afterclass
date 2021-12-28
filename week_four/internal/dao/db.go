package dao

import (
	"database/sql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type dbConfig struct {
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewDB() *sql.DB {
	var dbc dbConfig
	dbc.loadYaml()
	dsn := dbc.Username + ":" +
		dbc.Password + "@tcp(" + dbc.Host + ")/" +
		dbc.Dbname + "?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func (dbc *dbConfig) loadYaml() *dbConfig {
	WorkPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(WorkPath, "/week_four/config/config.yml")
	yamlFile, err := ioutil.ReadFile(appConfigPath)
	err = yaml.Unmarshal(yamlFile, dbc)
	if err != nil {
		panic(err)
	}
	return dbc
}
