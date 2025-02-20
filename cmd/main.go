package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type User struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Balance float64 `json:"balance"`
  isBlocked bool `json:"isBlocked"`
}

var users []User

func main() {
  r := gin.Default()

  r.GET("/users", func(c *gin.Context){
    c.JSON(http.StatusOK, users)
  })

  r.Run(":8080")
}