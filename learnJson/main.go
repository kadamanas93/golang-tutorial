package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"Title"`
	Author Author `json:"Author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}
	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading SensorReading
	json.Unmarshal([]byte(jsonString), &reading)
	fmt.Printf("%+v\n", reading)

	str := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var sreading map[string]interface{}
	json.Unmarshal([]byte(str), &sreading)
	fmt.Printf("%+v\n", sreading)
}
