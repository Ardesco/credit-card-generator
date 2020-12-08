package controllers

import (
	"github.com/Ardesco/credit-card-generator/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTypes(c *gin.Context) {

	var available models.ListOfKnownTypes

	for key, value := range models.AvailableCardTypes {
		available = append(available, models.CardType{
			Key:      key,
			LongName: value.LongName,
		})
	}

	c.JSON(http.StatusOK, gin.H{"known_cards": available})
}
