package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type blockDetails struct {
	Nonce interface{}     `json:"nonce"`
	Data  [][]interface{} `json:"data"`
}

type problemStat struct {
	Difficulty uint8        `json:"difficulty"`
	Block      blockDetails `json:"block"`
}

// struct for sending json payload with POST
type Post struct {
	Nonce string `json:"nonce"`
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
	if err != nil {
		fmt.Println("error while decoding the response data: ", err)
	}
	// fmt.Print(problem)

	// sending test post request
	postData := Post{
		Nonce: "123232223",
	}
	postUrl := "https://hackattic.com/challenges/mini_miner/solve?access_token=aaa699dde38ea86a"
	jsonPayload, err := json.Marshal(postData)
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

}
