package controllers

import (
	"io/ioutil"
	"net/http"
	processor "sum/processor"

	"github.com/gin-gonic/gin"
)

func Sum(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	mapInterface, err := processor.UnmarshallToInterface(jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
		return
	}
	total := processor.AddPayloadNumbers(mapInterface)
	result := processor.IntToHexSHA(total)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
