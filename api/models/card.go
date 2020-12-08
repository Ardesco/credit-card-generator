package models

type Card struct {
	Issuer     string     `json:"issuer"`
	Pan        PAN        `json:"pan"`
	ExpiryDate ExpiryDate `json:"expiryDate"`
	CVV        string     `json:"cvv"`
}

type ExpiryDate struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

type PAN struct {
	Raw       string `json:"raw"`
	Formatted string `json:"formatted"`
}
type CardProperties struct {
	LongName string
	Prefix   []string
	PanSize  int
	CvvSize  int
}

var AvailableCardTypes = map[string]CardProperties{
	"amex": {
		LongName: "American Express",
		Prefix:   []string{"37", "34"},
		PanSize:  15,
		CvvSize:  4,
	},
	"visa": {
		LongName: "Visa",
		Prefix:   []string{"4"},
		PanSize:  16,
		CvvSize:  3,
	},
	"mc": {
		LongName: "Mastercard",
		Prefix:   []string{"51", "52", "53", "54", "55"},
		PanSize:  16,
		CvvSize:  3,
	},
	"dci": {
		LongName: "Diners Club International",
		Prefix:   []string{"36", "38"},
		PanSize:  16,
		CvvSize:  3,
	},
	"jcb": {
		LongName: "Japan Credit Bureau",
		Prefix:   []string{"35"},
		PanSize:  16,
		CvvSize:  3,
	},
	"discover": {
		LongName: "Discover",
		Prefix:   []string{"6011", "65"},
		PanSize:  16,
		CvvSize:  3,
	},
}
