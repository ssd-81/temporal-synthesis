package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
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
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Pbkdf2   struct {
		Hash   string `json:"hash"`
		Rounds int    `json:"rounds"`
	} `json:"pbkdf2"`
	Scrypt struct {
		N      int `json:"N"`
		P      int `json:"p"`
		R      int `json:"r"`
		Buflen int `json:"buflen"`
	} `json:"scrypt"`
}

type solution struct {
	Sha256 string `json:"sha256"`
	Hmac   string `json:"hmac"`
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
	fmt.Print(problem)
	// converting salt string into byte
	saltByte, err := base64.StdEncoding.DecodeString(problem.Salt)
	fmt.Println(saltByte)

	// sample usage of sha256 and similar cryptographic methods
	// h := sha256.New()
	// h.Write([]byte("hello world\n"))
	// fmt.Printf("%x", h.Sum(nil))

	h := sha256.New()
	h.Write([]byte(problem.Password))
	fmt.Printf("%x", h.Sum(nil))

	// mac := hmac.New(sha256.New, key)
	// mac.Write(message)
	// expectedMAC := mac.Sum(nil)
	// return hmac.Equal(messageMAC, expectedMAC)

	mac := hmac.New(sha256.New, saltByte)
	mac.Write([]byte(problem.Password))

	// pbkdf2
	// rounds , _ := strconv.Atoi(problem.Pbkdf2.Rounds)
	// changed Rounds type to integer
	rounds := problem.Pbkdf2.Rounds

	// dk:= pbkdf2.Key([]byte(problem.Password),saltByte,rounds,32, sha256.New)
	// dk:= pbkdf2.Key(mac.Sum(nil),saltByte,rounds,32, sha256.New)
	dk := pbkdf2.Key([]byte(problem.Password), saltByte, rounds, 32, sha256.New)

	if err != nil {
		fmt.Println("error encountered while deriving key using pbkdf2")
	}

	// scrypt
	sr, _ := scrypt.Key([]byte(problem.Password), saltByte, problem.Scrypt.N, problem.Scrypt.R, problem.Scrypt.P, problem.Scrypt.Buflen)

	// sending response
	// Create JSON data
	data := solution{
		Sha256: hex.EncodeToString(h.Sum(nil)),
		Hmac:   hex.EncodeToString(mac.Sum(nil)),
		Pbkdf2: hex.EncodeToString(dk),
		Scrypt: hex.EncodeToString(sr),
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
	fmt.Println()
	fmt.Println(">>>>>>>>>", "<<<<<<<<<<<<<")
	fmt.Println(string(resBytes))

}
