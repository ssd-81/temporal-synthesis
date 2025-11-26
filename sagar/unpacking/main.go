package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type problemSet struct {
	Data string `json:"bytes"`
}

type solutionPaylod struct {
	// int: the signed integer value
	// uint: the unsigned integer value
	// short: the decoded short value
	// float: surprisingly, the float value
	// double: the double value - shockingly
	// big_endian_double: you get the idea by now!
	IntVal          int32   `json:"int"`
	UintVal         uint32  `json:"uint"`
	ShortVal        int16   `json:"short"`
	FloatVal        float32 `json:"float"`
	DoubleVal       float64 `json:"double"`
	BigEndianDouble float64 `json:"big_endian_double"`
}

func main() {

	// get problem set from: https://hackattic.com/challenges/help_me_unpack/problem?access_token=aaa699dde38ea86a
	getUrl := "https://hackattic.com/challenges/help_me_unpack/problem?access_token=aaa699dde38ea86a"
	// problem?access_token=aaa699dde38ea86a
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

	// a regular int (signed), to start off
	// an unsigned int
	// a short (signed) to make things interesting
	// a float because floating point is important
	// a double as well
	// another double but this time in big endian (network byte order)

	var regInt int32
	var unsignedInt uint32
	var shortSig int16
	var floatVal float32
	var doubleVal float64
	var doubleValBigEnd float64

	rawBytes, err := base64.StdEncoding.DecodeString(problem.Data)
	if err != nil {
		fmt.Println("error encounterd while decoding provided input")
		return
	}

	// for debugging
	fmt.Printf("Total bytes: %d\n", len(rawBytes))
	fmt.Printf("bytes: % x\n", rawBytes)

	buf := bytes.NewReader(rawBytes)
	err = binary.Read(buf, binary.LittleEndian, &regInt)
	if err != nil {
		fmt.Println("error encounterd while parsing binary")
		return
	}
	fmt.Printf("After regInt, position: %d\n", len(rawBytes)-buf.Len())

	fmt.Println(regInt)

	binary.Read(buf, binary.LittleEndian, &unsignedInt)
	fmt.Printf("After unsignedInt, position: %d\n", len(rawBytes)-buf.Len())

	binary.Read(buf, binary.LittleEndian, &shortSig)
	fmt.Printf("After shortSig, position: %d\n", len(rawBytes)-buf.Len())

	// min memory block = 32 bits = 4 bytes
	// shortSig => 2 bytes
	// 2 bytes is padded
	var padding uint16
	binary.Read(buf, binary.LittleEndian, &padding)

	binary.Read(buf, binary.LittleEndian, &floatVal)
	fmt.Printf("After float val , position: %d\n", len(rawBytes)-buf.Len())

	binary.Read(buf, binary.LittleEndian, &doubleVal)
	fmt.Printf("After doubleVal, position: %d\n", len(rawBytes)-buf.Len())

	binary.Read(buf, binary.BigEndian, &doubleValBigEnd)
	fmt.Printf("After doubleValBigEnd, position: %d\n", len(rawBytes)-buf.Len())

	sol := solutionPaylod{
		IntVal:          regInt,
		UintVal:         unsignedInt,
		ShortVal:        shortSig,
		FloatVal:        floatVal,
		DoubleVal:       doubleVal,
		BigEndianDouble: doubleValBigEnd,
	}
	jsonData, err := json.Marshal(sol)
	if err != nil {
		fmt.Println(err)
		return
	}
	solUrl := "https://hackattic.com/challenges/help_me_unpack/solve?access_token=aaa699dde38ea86a"
	req, err := http.NewRequest("POST", solUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("request object could not be created")
		return
	}
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("response could not be sent")
		return
	}
	defer resp.Body.Close()
	resBytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(resBytes))

}
