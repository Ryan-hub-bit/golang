package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 19, []string{"xingye", "zhangbozhi"}}
	jsonStr, error := json.Marshal(movie)
	if error != nil {
		fmt.Println("json marsal error", error)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	myMovie := Movie{}
	error = json.Unmarshal(jsonStr, &myMovie)
	if error != nil {
		fmt.Println("json unmarshal error", error)
		return
	}

	fmt.Printf("%v\n", myMovie)
}
