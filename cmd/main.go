package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	IsBlocked bool    `json:"isBlocked"`
}

var users []User

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to convert id(string) into integer")
		}
		for _, user := range users {
			if user.ID == i {
				c.JSON(http.StatusOK, user)
				return
			}
		}
	})

	r.POST("/users", func(c *gin.Context) {

	})

	r.Run(":8080")
}
