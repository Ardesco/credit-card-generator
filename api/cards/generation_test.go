package cards

import (
	"github.com/Ardesco/credit-card-generator/api/models"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestThatExpiryDateIsGenerated(t *testing.T) {
	result := GenerateExpiryDate
	assert.GreaterOrEqual(t, result().Year, time.Now().Year())
	assert.LessOrEqual(t, result().Year, time.Now().Year()+5)
	assert.GreaterOrEqual(t, result().Month, 1)
	assert.LessOrEqual(t, result().Month, 12)
}

func TestThatExpiryCVVIsGeneratedCorrectlyForVisa(t *testing.T) {
	result := GenerateCVV(models.AvailableCardTypes["visa"].CvvSize)
	value, err := strconv.Atoi(result)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(result), 3)
	assert.GreaterOrEqual(t, value, 100)
	assert.LessOrEqual(t, value, 999)
}

func TestThatExpiryCVVIsGeneratedCorrectlyForAmex(t *testing.T) {
	result := GenerateCVV(models.AvailableCardTypes["amex"].CvvSize)
	value, err := strconv.Atoi(result)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(result), 4)
	assert.GreaterOrEqual(t, value, 1000)
	assert.LessOrEqual(t, value, 9999)
}

func TestGenerateAVisaPan(t *testing.T) {
	result := GeneratePAN(models.AvailableCardTypes["visa"])
	assert.Equal(t, len(result.Raw), 16)
	assert.Equal(t, result.Raw[0], "4"[0])
}

func TestGenerateAnAmexPan(t *testing.T) {
	result := GeneratePAN(models.AvailableCardTypes["amex"])
	assert.Equal(t, len(result.Raw), 15)
	assert.Equal(t, result.Raw[0], "3"[0])
}

func TestGenerateAMastercardPan(t *testing.T) {
	result := GeneratePAN(models.AvailableCardTypes["mastercard"])
	assert.Equal(t, len(result.Raw), 16)
	assert.Equal(t, result.Raw[0], "5"[0])
}
