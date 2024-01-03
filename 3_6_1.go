package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	Birthday   string `json:"Birthday"`
	Address    string `json:"Address"`
	Phone      string `json:"Phone"`
	Rating     []int  `json:"Rating"`
}

type Group struct {
	ID       int       `json:"ID"`
	Number   string    `json:"Number"`
	Year     int       `json:"Year"`
	Students []Student `json:"Students"`
}

func main_3_6_1() {
	var group Group

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading data from stdin:", err)
		return
	}

	err = json.Unmarshal(data, &group)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	averageRating := calculateAverageRating(group.Students)
	result := map[string]float64{"Average": averageRating}

	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(resultJSON))
}

func calculateAverageRating(students []Student) float64 {
	total := 0
	count := 0
	for _, student := range students {
		total += len(student.Rating)
		count++
	}
	return float64(total) / float64(count)
}
