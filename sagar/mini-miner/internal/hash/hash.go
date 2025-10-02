package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Test(data []byte) {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hashedData := hasher.Sum(nil)
	hexHash := hex.EncodeToString(hashedData)

	fmt.Println("SHA-256 Hash:", hexHash)

}

