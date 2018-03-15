package main

import (
	"crypto/md5"
	"fmt"
	"crypto/rand"
	"crypto/rsa"
	"math/big"

)

func main() {

	hash :=md5.New()
	hash.Write([]byte("du bist der Dude"))

	fmt.Println(hash.Sum(nil))

	reader := rand.Reader
	bitsize :=512
	key,err := rsa.GenerateKey(reader, bitsize)

	fmt.Println("publickey: ", key.PublicKey, err)
	fmt.Println("publickey exponent: " , key.PublicKey.E)
	fmt.Println("publickey N: " , key.PublicKey.N)
	fmt.Println("privatekey: ", key.D)
	fmt.Println("key: ", key)

	fmt.Println("prime0: ", key.Primes[0])
	fmt.Println("prime1: ", key.Primes[1])
	//p0 := *key.Primes[0]
	//p1:= *key.Primes[1]
	 p2 := *big.NewInt(0)
	p3 :=  p2.Mul(key.Primes[0], key.Primes[1])

	fmt.Println("p1 * p2: ", p3)

	fmt.Println("exponent: " , key.D.String)
	fmt.Println("publickey: " , key.E)


	fmt.Println("privatekey: ", key.D)


}
