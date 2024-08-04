package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonDataResponse struct { //sending to service1
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	SchoolName string `json:"schoolName"`
}

type PersonDataReceive struct { //received from service1
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetData(c *gin.Context) {
	var dataReceive PersonDataReceive
	err := c.BindJSON(&dataReceive)
	if err != nil {
		fmt.Println("Error ocuured in BindJson", err)
	}
	output := dataManipulation(dataReceive)
	fmt.Println("jsonByte::", output)
	c.JSON(http.StatusOK, output)
}

func dataManipulation(dataReceive PersonDataReceive) PersonDataResponse {
	fmt.Println("data:", dataReceive)
	schoolName := "gayatri Convent"
	var dataResponse PersonDataResponse
	dataResponse.FirstName = dataReceive.FirstName
	dataResponse.LastName = dataReceive.LastName
	dataResponse.SchoolName = schoolName

	fmt.Println("dataResponse::", dataResponse)
	return dataResponse
}
