package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
  "models"
)

func main() {
  var pool *sqlx.DB
  dsn := "root:@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local"
  pool, err := sqlx.Open("mysql", dsn)
  if err != nil {
    panic(err)
  }

  r := gin.Default()
  r.GET("/", func(c *gin.Context){
    c.JSON(200, gin.H{"message": "pong"})
  })
  r.GET("/users", func(c *gin.Context){
    people := []models.User{}
    pool.Select(&models.User, "SELECT * FROM users ORDER BY id ASC")

    c.JSON(200, people)
  })
  r.Run()
}
