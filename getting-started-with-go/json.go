package main

import (
	"encoding/json"
	"fmt"
)

type Weather struct {
	Temp     float64 `json:"temp"`
	Clouds   int     `json:"clouds"`
	Forecast string  `json:"forecast"`
}

func main() {
	w1 := Weather{
		Temp:     2.4,
		Clouds:   100,
		Forecast: "Bad",
	}

	w1Json, err := json.Marshal(w1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(w1Json), "|", w1)

	var w2 Weather

	err = json.Unmarshal(w1Json, &w2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w2)
}
