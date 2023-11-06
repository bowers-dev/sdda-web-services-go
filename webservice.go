package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dog struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner name"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
}

var dags = []dog{
	{ID: "1", Name: "Watson", Owner: "Jennie", Birthday: "9/23/19", Email: "email@email.com"},
	{ID: "2", Name: "Eli", Owner: "Charles", Birthday: "unknown", Email: "email@email.com"},
	{ID: "3", Name: "Julius", Owner: "Nick", Birthday: "unknown", Email: "email@email.com"},
}

func getDags(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dags)
}

func main() {
	router := gin.Default()
	router.GET("/dogs", getDags)

	router.Run("localhost:42069")
}
