package hash

import (
	"crypto/sha256"
)

func CheckSolution(data []byte, diff uint8) bool {
	hashedData := sha256.Sum256(data) // this returns a 32 bit array
	fullZeroBytes := diff / 8
	remainingBits := diff % 8

	// checking for full zero bytes
	for i := uint8(0); i < fullZeroBytes; i++ {
		if hashedData[i] != 0x00 {
			return false
		}
	}

	if remainingBits > 0 {
		// mask is used to compare the stuck bits
		mask := uint8(0xFF << (8 - remainingBits))

		// this extracts only the relevent bits that are part of the difficulty; it will ignore all other
		// bits that are irrelevent
		if mask&hashedData[fullZeroBytes] != 0 { // might need changes
			return false
		}

	}
	return true
}
