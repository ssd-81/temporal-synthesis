package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type problemSet struct {
	Data string `json:"bytes"`
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
	rawBytes, err := base64.StdEncoding.DecodeString(problem.Data)
	fmt.Println("raw bytes", rawBytes)

	// strByte := base64.StdEncoding.EncodeToString(problem.Data)
	// fmt.Println("base64: ", strByte)

	// deB64, _ := base64.StdEncoding.DecodeString(strByte)
	// fmt.Println("decoded base64: ", deB64)

	// hexStr := hex.EncodeToString(problem.Data)
	// fmt.Println("hex: ", hexStr)

	// intV := problem.Data[0:4]
	// fmt.Println("test 0x1", base64.StdEncoding.EncodeToString(intV))

	// posting the solution
	// POST /challenges/help_me_unpack/solve?access_token=...

}
