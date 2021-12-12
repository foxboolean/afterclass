package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

type Service struct {
	db *sql.DB
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/name":
		s.QueryNameService(w)
		return
	}
}

func (s *Service) QueryNameService(w http.ResponseWriter) {
	name, err := s.QueryNameDao()
	if err != nil {
		fmt.Printf("%v", err)
		if IsNotRows(err) {
			http.Error(w, "no data", 500)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, name)
	return
}

func (s *Service) QueryNameDao() (string, error) {
	db := s.db
	var name string
	sqlStr := "select name from people where id=?"
	err := db.QueryRow(sqlStr,4).Scan(&name)
	if IsNotRows(err) {
		err = fmt.Errorf("database has no data \n%w", err)
	}
	return name, err
}

func IsNotRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
