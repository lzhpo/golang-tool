package struct_to_map

import (
	"testing"
)

// UserRequestBody
type UserRequestBody struct {
	Username  string      `json:"username"`
	Password  string      `json:"password"`
	Profile   Profile     `json:"habit"`
	Addresses []Addresses `json:"addresses"`
}

// Habit
type Profile struct {
	Age   int    `json:"age"`
	Sex   string `json:"sex"`
	Habit string `json:"habit"`
}

// Addresses
type Addresses struct {
	AddressTypes []string `json:"addressTypes"`
	Street       string   `json:"street"`
	HouseNumber  string   `json:"houseNumber"`
	City         string   `json:"city"`
	Country      string   `json:"country"`
	PostalCode   string   `json:"postalCode"`
}

var userRequestBody = UserRequestBody{
	Username: "lzhpo",
	Profile: Profile{
		Age:   21,
		Sex:   "男",
		Habit: "Play basketball",
	},
	Addresses: []Addresses{
		{
			AddressTypes: []string{
				"Type1", "Type2",
			},
			Street:      "Berlin 123 Street",
			HouseNumber: "123456",
			City:        "Berlin",
			Country:     "Germany",
			PostalCode:  "123456",
		},
		{
			AddressTypes: []string{
				"Type1", "Type2",
			},
			Street:      "Berlin 456 Street",
			HouseNumber: "456789",
			City:        "Berlin",
			Country:     "Germany",
			PostalCode:  "456789",
		},
	},
}

func TestUnmarshalToMap(t *testing.T) {
	UnmarshalToMap(userRequestBody)
}

func TestReflectToMap(t *testing.T) {
	ReflectToMap(userRequestBody, "json")
}
