package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type blockDetails struct {
	Data [][]string `json:"data"`
}

type problemStat struct {
	Difficulty uint8        `json:"difficulty"`
	Block      blockDetails `json:"block"`
}

func main() {

	// getting the problem statement
	// from : https://hackattic.com/challenges/mini_miner/problem?access_token=aaa699dde38ea86a

	url := "https://hackattic.com/challenges/mini_miner/problem?access_token=aaa699dde38ea86a"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error while getting response from the url: ", url)
	}
	defer resp.Body.Close()
	// fmt.Print(resp)

	// initializing the target struct
	var problem problemStat
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&problem)
	fmt.Print(problem)
}
