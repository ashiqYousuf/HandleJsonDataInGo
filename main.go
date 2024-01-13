package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// !MARSHALLING || SERIALIZATION
// ?Encoding Go objects to JSON format is known as marshaling
// ?func Marshal(v interface{}) ([]byte, error)
// ?It accepts an empty interface. In other words, you can provide any Go data type to the function

type Book struct {
	Title  string `json:"title,"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type Seller struct {
	Id          int    `json:"sellerId"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}
type Product struct {
	Id     int    `json:"productId"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Seller Seller `json:"seller,omitempty"`
}

func Encode(obj interface{}) (string, error) {
	// byteData, err := json.Marshal(obj)
	byteData, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return "", err
	}
	return string(byteData), nil
}

func EncodingExamples() {
	// *Encoding map D.S. to JSON

	fileCounter := map[string]int{
		"cpp":        10,
		"go":         3,
		"python":     11,
		"javascript": 5,
	}
	jsonData, err := Encode(fileCounter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", jsonData)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")

	// *Encoding struct to JSON

	book := Book{"Data Structure", "Ashiq Hussain", 2022}
	bookJson, err := Encode(book)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", bookJson)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")

	// *Encoding Complex Objects to JSON

	products := []Product{
		{
			Id:     50,
			Name:   "Writing Book",
			Seller: Seller{1, "ABC Company", "US"},
			Price:  100,
		},
		{
			Id:     51,
			Name:   "Kettle",
			Seller: Seller{20, "John Store", "DE"},
			Price:  500,
		},
		{
			Id:    52,
			Name:  "Laptop",
			Price: 6790,
		},
	}

	complexJson, err := Encode(products)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", complexJson)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")
}

func main() {
	// EncodingExamples()
}
