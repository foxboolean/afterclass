package dao

import (
	"database/sql"
)

type UserDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO{
	return &UserDAO{
		db: db,
	}
}

func (u *UserDAO) GetUser(id string) (*User, error) {
	var name string
	sqlStr := "select name from people where id=?"
	err := u.db.QueryRow(sqlStr,id).Scan(&name)
	return &User{
		Name: name,
	}, err
}

type User struct {
	Name string
}
