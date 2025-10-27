package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type problemSet struct {
	Data []byte `json:"bytes"`
}

func main() {

	// get problem set from: https://hackattic.com/challenges/help_me_unpack/problem?access_token=aaa699dde38ea86a
	getUrl := "https://hackattic.com/challenges/help_me_unpack/problem?access_token=aaa699dde38ea86a"
	resp, err := http.Get(getUrl)

	if err != nil {
		fmt.Printf("error while getting the base64 encoded data")
	}
	defer resp.Body.Close()

	var problem problemSet
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&problem)
	if err != nil {
		fmt.Println("error while decoding the response data")
	}
	// strContent, err := decoder.String(problem.Data)
	fmt.Println(problem.Data)

	// understanding encoding basics
	var a byte = 'A'
	fmt.Println(a)
	var b string = "A"
	fmt.Println("string: ", b)

	understanding := base64.StdEncoding.EncodeToString(problem.Data)
	fmt.Println("base64: ", understanding)

	hexStr := hex.EncodeToString(problem.Data)
	fmt.Println("hex: ", hexStr)

	
}
