package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"

	"bytes"
	"crypto/aes"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"os"
	"time"
)

func main() {

	hash := md5.New()
	hash.Write([]byte("Text to be hashed"))
	//fmt.Println(hash.Sum(nil))

	aesKey := []byte("this is my key !")
	aesCipher, err := aes.NewCipher(aesKey)
	handleResult("Error creating aesCipher: ", err)
	fmt.Printf("Cipher: %v\n\n", aesCipher)

	rreader := rand.Reader
	bitsize := 512
	key, err := rsa.GenerateKey(rreader, bitsize)

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

	now := time.Now()
	until := now.Add(365 * 24 * time.Hour)

	template := x509.Certificate{
		SerialNumber: big.NewInt(1504),
		Subject: pkix.Name{
			CommonName:   "chrisi.de",
			Country:      []string{"de"},
			Organization: []string{"dude.co"}},
		NotBefore:             now,
		NotAfter:              until,
		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDataEncipherment | x509.KeyUsageKeyEncipherment,
		BasicConstraintsValid: true,
		IsCA:     true,
		DNSNames: []string{"chrissi.de", "localhost"},
	}

	fmt.Printf("%T|\n", key)
	derBytes, err := x509.CreateCertificate(rreader, &template, &template, &key.PublicKey, key)
	handleResult("Couldn´t create certificate", err)

	fmt.Println(derBytes)
	var certbytes []byte
	//certbytes := make([]byte, 0, 5000)
	certbuffer := bytes.NewBuffer(certbytes)
	err = pem.Encode(certbuffer, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	handleResult("Certificat not Encoded", err)
	fmt.Println("Cert\n ", certbuffer, "\n\n")
	var keybytes []byte
	keybuffer := bytes.NewBuffer(keybytes)
	err = pem.Encode(keybuffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	handleResult("Key not Encoded", err)
	fmt.Println("Key\n", keybuffer, "\n\n")

}

func handleResult(text string, err error) {
	if err != nil {
		fmt.Printf("Error occurrend: %s, --> %v", text, err)
		os.Exit(1)
	}
}
