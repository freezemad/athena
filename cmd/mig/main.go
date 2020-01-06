package main

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  UserName varchar(50) NOT NULL,
  Email varchar(100) DEFAULT NULL,
  HashPassword varchar(200) NOT NULL,
  Created datetime NOT NULL,
  Updated datetime NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4`

func main() {
  var pool *sqlx.DB
  dsn := "root:@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local"
  pool, err := sqlx.Open("mysql", dsn)
  if err != nil {
    panic(err)
  }
  pool.MustExec(schema)
  fmt.Println("migrated!")
}
