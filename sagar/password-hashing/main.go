package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// password: the password you'll operate on
// salt: the salt we'll use - also user as a secret where necessary; keep in mind it comes base64 encoded - decode for the raw bytes
// pbkdf2:
// hash: the digest to use
// rounds: the number of rounds to use
// scrypt:
// N: the N parameter for scrypt's KDF
// p: the parallelization parameter
// r: the blocksize parameter
// buflen: intended output length in octets
// _control: example scrypt calculated for password="rosebud", salt="pepper", N=128, p=8, n=4


type problemStatement struct {
	Password string `jsong:"password"`
	Salt string `json:"salt"`
	Pbkdf2 struct {
		Hash string `json:"hash"`
		Rounds string `json:"rounds"`
	}`json:"pbkdf2"`
	Scrypt struct {
		N int `json:"N"` 
		P int `json:"p"`
		R int `json:"r"`
		Buflen int `json:"buflen"`
	}`json:"scrypt"`

}

type solution struct {
	Sha256 string `json:"sha256"`
	Hmac string `json:"hmac"`
	Pbkdf2 string `json:"pbkdf2"`
	Scrypt string `json:"scrypt"`
}
// get: https://hackattic.com/challenges/password_hashing/problem?access_token=aaa699dde38ea86a
// send: https://hackattic.com/challenges/password_hashing/solve?access_token=aaa699dde38ea86a

func main() {

	resp, err := http.Get("https://hackattic.com/challenges/password_hashing/problem?access_token=aaa699dde38ea86a")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var problem problemStatement
	json.Unmarshal(body, &problem)
	// fmt.Print(problem)
	// converting salt string into byte 
	saltByte , err := base64.StdEncoding.DecodeString(problem.Salt)
	fmt.Println(saltByte)


	// sample usage of sha256 and similar cryptographic methods 
	// h := sha256.New()
	// h.Write([]byte("hello world\n"))
	// fmt.Printf("%x", h.Sum(nil))

	h := sha256.New()
	h.Write([]byte(problem.Password))
	fmt.Printf("%x", h.Sum(nil))




	// sending response 
	// Create JSON data
	data := solution{
		Sha256: hex.EncodeToString(h.Sum(nil)),
		Hmac: hex.EncodeToString(h.Sum(nil)),
		Pbkdf2: hex.EncodeToString(h.Sum(nil)),
		Scrypt: hex.EncodeToString(h.Sum(nil)),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
	panic(err)
	}

	// Make POST request with JSON data
	resp, err = http.Post("https://hackattic.com/challenges/password_hashing/solve?access_token=aaa699dde38ea86a", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
	panic(err)
	}
	defer resp.Body.Close()
	resBytes, _ := io.ReadAll(resp.Body)
	fmt.Println(">>>>>>>>>", "<<<<<<<<<<<<<")
	fmt.Println(string(resBytes))

}
