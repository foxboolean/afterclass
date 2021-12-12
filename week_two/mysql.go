package main


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// 在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
func main() {
	startService()
}

func startService() {
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	s := &Service{db: db}
	http.ListenAndServe(":8080", s)
}