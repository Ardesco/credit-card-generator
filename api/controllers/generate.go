package controllers

import (
	"github.com/Ardesco/credit-card-generator/api/cards"
	"github.com/Ardesco/credit-card-generator/api/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const MaxCount = 500

func GenerateCards(c *gin.Context) {
	countString := c.DefaultQuery("count", "1")
	cardType := c.DefaultQuery("type", "visa")

	count, countError := strconv.Atoi(countString)
	cardProperties, exists := models.AvailableCardTypes[cardType]

	var errors models.Errors

	if countError != nil {
		errors = append(errors, models.Error{
			Parameter: "count",
			Issue:     "invalid value supplied",
		})
	}

	if !exists {
		errors = append(errors, models.Error{
			Parameter: "type",
			Issue:     "invalid card type supplied",
		})
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	if count > MaxCount {
		count = MaxCount
	}

	var cardList []models.Card
	for i := 0; i < count; i++ {

		rand.Seed(time.Now().UnixNano())

		var card = models.Card{
			Issuer:     cardProperties.LongName,
			Pan:        cards.GeneratePAN(cardProperties),
			ExpiryDate: cards.GenerateExpiryDate(),
			CVV:        cards.GenerateCVV(cardProperties.CvvSize),
		}

		cardList = append(cardList, card)
	}

	c.JSON(http.StatusOK, cardList)
}

func GenerateCard(c *gin.Context) {
	cardType := c.DefaultQuery("type", "visa")
	cardProperties, exists := models.AvailableCardTypes[cardType]

	var errors models.Errors

	if !exists {
		errors = append(errors, models.Error{
			Parameter: "type",
			Issue:     "invalid card type supplied",
		})
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	rand.Seed(time.Now().UnixNano())

	var card = models.Card{
		Issuer:     cardProperties.LongName,
		Pan:        cards.GeneratePAN(cardProperties),
		ExpiryDate: cards.GenerateExpiryDate(),
		CVV:        cards.GenerateCVV(cardProperties.CvvSize),
	}

	c.JSON(http.StatusOK, card)
}
