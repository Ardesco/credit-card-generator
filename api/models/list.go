package models

type ListOfKnownTypes []CardType

type CardType struct {
	Key      string `json:"type"`
	LongName string `json:"name"`
}
