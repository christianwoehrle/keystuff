package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"

	"crypto/aes"
	"os"
)

func main() {

	hash := md5.New()
	hash.Write([]byte("Text to be hashed"))
	//fmt.Println(hash.Sum(nil))

	aesKey := []byte("this is my key !")
	aesCipher, err := aes.NewCipher(aesKey)
	handleResult("Error creating aesCipher: ", err)
	fmt.Printf("Cipher: %v\n\n", aesCipher)

	reader := rand.Reader
	bitsize := 512
	key, err := rsa.GenerateKey(reader, bitsize)

	fmt.Println("prime0: ", key.Primes[0])
	fmt.Println("prime1: ", key.Primes[1])
	//p0 := *key.Primes[0]
	//p1:= *key.Primes[1]
	p2 := *big.NewInt(0)
	p3 := p2.Mul(key.Primes[0], key.Primes[1])

	fmt.Println("Expected Modulus: ", p3)

	fmt.Println("Modulus         : ", key.N)
	fmt.Println("publickey E:  ", key.E)

	fmt.Println("privatekey D: ", key.D)

}

func handleResult(text string, err error) {
	if err != nil {
		fmt.Printf("Error occurrend: %s, --> %v", text, err)
		os.Exit(1)
	}
}
