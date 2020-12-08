package cards

import (
	"github.com/Ardesco/credit-card-generator/api/models"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func calculateLuhn(digits string) string {
	number, _ := strconv.Atoi(digits)
	checkNumber := checksum(number)
	if checkNumber == 0 {
		return "0"
	}

	return strconv.Itoa(10 - checkNumber)
}

func checksum(number int) int {
	var sum int
	for i := 0; number > 0; i++ {
		digit := number % 10
		if i%2 == 0 { // even
			digit = digit * 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}
		sum += digit
		number = number / 10
	}

	return sum % 10
}

func GeneratePAN(properties models.CardProperties) models.PAN {
	var prefix = properties.Prefix[rand.Intn(len(properties.Prefix))]
	for len(prefix) < properties.PanSize-1 {
		prefix = prefix + generateRandomNumberOfSize(1)
	}

	number := prefix + calculateLuhn(prefix)
	return models.PAN{
		Raw:       number,
		Formatted: FormatPan(properties.LongName, number),
	}
}

func FormatPan(cardType string, pan string) string {
	var formattedPan string
	for i := 0; i < len(pan); i++ {
		formattedPan = formattedPan + string(pan[i])
		switch cardType {
		case "American Express":
			if i == 3 || i == 9 {
				formattedPan = formattedPan + " "
			}
		default:
			if (i+1)%4 == 0 {
				formattedPan = formattedPan + " "
			}
		}
	}

	return strings.TrimSpace(formattedPan)
}

func generateRandomNumberOfSize(size int) string {
	var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var randomNumber []string
	for i := 0; i < size; i++ {
		randomNumber = append(randomNumber, numbers[rand.Intn(len(numbers))])
	}
	return strings.Join(randomNumber, "")
}

func GenerateCVV(size int) string {
	return generateRandomNumberOfSize(size)
}

func GenerateExpiryDate() models.ExpiryDate {
	var year = time.Now().Year() + rand.Intn(6)
	var month = rand.Intn(12) + 1

	return models.ExpiryDate{
		Month: month,
		Year:  year,
	}
}
