package main

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
	Salt []byte `json:"salt"`

}
type pb struct {
	Hash string `json:"hash"`
	Rounds string `json:"rounds"`
}
type sc struct {
	N string `json:"N"`
	p string `json:"p"`
	r string `json:"r"`
	BufLen int `json:"buflen"` // intended output length in octets
	Control string `json:"_control"` // not sure what this is 
}