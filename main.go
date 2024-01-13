package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// func init() {
// 	fmt.Println("Learn JSON Encoding & Decoding!")
// }

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

// !MARSHALLING || SERIALIZATION
// ?Encoding Go objects to JSON format is known as marshaling
// ?func Marshal(v interface{}) ([]byte, error)
// ?It accepts an empty interface. In other words, you can provide any Go data type to the function

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

// !UNMARSHALLING || DE-SERIALIZATION
// ?converting JSON to Go objects is called Un-Marshalling
// ?func Unmarshal(data []byte, v interface{}) error
// ?pass a reference to store the decoded content

func Decode(jsonData string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonData), obj)

	if err != nil {
		return err
	}
	return nil
}

func DecodingExamples() {
	// ?Decoding or De-Serializing JSON into Object

	jsonBook := `{
		"title": "My Book",
		"author": "Hossien",
		"year": 2023
	}`
	var book Book
	err := Decode(jsonBook, &book)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", book)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")

	// ?Decoding JSON to Map D.S
	jsonInput := `{
        "apples": 10,
        "mangos": 20,
        "grapes": 20
    }`
	fruits := map[string]int{}
	err = Decode(jsonInput, &fruits)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fruits)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")

	// ?Decoding complex JSON
	complexJsonInput := `[
		{
			"productId":50,
			"name":"Writing Book",
			"seller":{
				"sellerId":1,
				"name":"ABC Company",
				"countryCode":"US"
			},
			"price":100
		},
		{
			"productId":51,
			"name":"Kettle",
			"seller":{
				"sellerId":20,
				"name":"John Store",
				"countryCode":"DE"
			},
			"price":500
		}]
		`
	products := []Product{}
	err = Decode(complexJsonInput, &products)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", products)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")
}

type Window struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

type Config struct {
	Timeout     float32 `json:"timeout"`
	PluginsPath string  `json:"pluginsPath"`
	Window      Window  `json:"window"`
}

func ReadingJsonFiles(filepath string, obj interface{}) {
	byteData, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}
	err = Decode(string(byteData), obj)
	if err != nil {
		log.Fatal(err)
	}
}

func WritingJsonFiles(filepath string, obj interface{}) error {
	strData, err := Encode(obj)

	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, []byte(strData), 0666)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// EncodingExamples()
	// DecodingExamples()
	// Reading Json Files in a struct variable
	var c Config
	ReadingJsonFiles("config.json", &c)
	fmt.Printf("%#v\n", c)
	// Writing to Json Files
	c.PluginsPath = "usr/bin/plugins/"
	WritingJsonFiles("config.json", c)
}
