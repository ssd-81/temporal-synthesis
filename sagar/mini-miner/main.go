package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/ssd-81/temporal-synthesis/sagar/mini-miner/internal/hash"
)

type blockDetails struct {
	Nonce any     `json:"nonce"`
	Data  [][]any `json:"data"`
}

type problemStat struct {
	Difficulty uint8        `json:"difficulty"`
	Block      blockDetails `json:"block"`
}

// struct for sending json payload with POST
type Post struct {
	Nonce uint `json:"nonce"`
}

type hashBlock struct {
	Data [][]any `json:"data"`
	Nonce uint `json:"nonce"`
}

func main() {

	// getting the problem statement
	// from : https://hackattic.com/challenges/mini_miner/problem?access_token=aaa699dde38ea86a

	url := "https://hackattic.com/challenges/mini_miner/problem?access_token=aaa699dde38ea86a"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error while getting response from the url: %v", url)
	}
	defer resp.Body.Close()
	// fmt.Print(resp)

	// initializing the target struct
	var problem problemStat
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&problem)
	if err != nil {
		fmt.Println("error while decoding the response data: ", err)
	}


	// finding the solution nonce 
	tempBlock := hashBlock{
		Data: problem.Block.Data,
		Nonce: 0,
	}
	jsonSerializedData , err := json.Marshal(tempBlock)
	if err != nil {
		fmt.Println("error while converting go struct to JSON: ", err)
		return 
	}
	diff := ""
	for i := 0 ; i < int(problem.Difficulty); i++ {
		diff += "0"
	}
	for {
		if x := hash.CheckSolution(jsonSerializedData, diff); x {
			fmt.Println("found it")
			break
		}
		fmt.Println("no match")
	}

	hash.Test(jsonSerializedData)


	// sending the solution to required endpoint 
	solutionNonce := Post{
		Nonce: 0,
	}
	postUrl := "https://hackattic.com/challenges/mini_miner/solve?access_token=aaa699dde38ea86a"
	jsonPayload, err := json.Marshal(solutionNonce)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	reqBody := bytes.NewBuffer(jsonPayload)
	req, err := http.NewRequest("POST", postUrl, reqBody)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}
	// setting the Header for letting the server know the format of data
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error while making POST request: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	fmt.Println("successly sent post request")

	// creating the json serialized data for hashing

	

}
