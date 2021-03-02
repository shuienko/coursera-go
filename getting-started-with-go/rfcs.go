package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.purpleair.com/json?key=AGZYH0ZCY7TFBSP6&show=49489")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)
}
