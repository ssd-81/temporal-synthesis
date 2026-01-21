package main

import (
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
	Password string `json:"password"`
	Salt     string `json:"salt"` // base64 encoded
	Pb pb `json:"pbkdf2"`
	Sc sc `json:"scrypt"`

}
type pb struct {
	Hash   string `json:"hash"`
	Rounds int32 `json:"rounds"`
	
}
type sc struct {
	N       int32 `json:"N"`
	P       int32 `json:"p"`
	R       int32 `json:"r"`
	BufLen  int32    `json:"buflen"`   // intended output length in octets
	Control int32 `json:"_control"` // not sure what this is
}

// get: https://hackattic.com/challenges/password_hashing/problem?access_token=aaa699dde38ea86a

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
}
