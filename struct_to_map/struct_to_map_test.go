package struct_to_map

import (
	"fmt"
	"testing"
)

// UserRequestBody
type UserRequestBody struct {
	ID        int         `json:"id"`
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

// 测试userRequestBody
var userRequestBody = UserRequestBody{
	ID:       123456,
	Username: "lzhpo",
	Password: "123456",
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
	toMapValue, unmarshalToMapError := UnmarshalToMap(userRequestBody)
	if unmarshalToMapError != nil {
		fmt.Println("UnmarshalToMap:", unmarshalToMapError)
	}
	for k, v := range toMapValue {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
}

func TestReflectToMap(t *testing.T) {
	reflectToMap, reflectToMapError := ReflectToMap(userRequestBody, "json")
	if reflectToMapError != nil {
		fmt.Println("ReflectToMap Error:", reflectToMapError)
	}
	for k, v := range reflectToMap {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
}
