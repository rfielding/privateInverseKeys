package main

import (
	"crypto/rsa"
	"fmt"
	"io"
	"math/big"
	"math/rand"
)

type randBytes struct {
}

func (r *randBytes) Read(arr []byte) (bytesRead int, err error) {
	L := len(arr)
	err = nil
	for i := 0; i < L; i++ {
		arr[i] = byte(rand.Int())
	}
	return L, err
}

type pubKey struct {
	N big.Int
	E big.Int
}

type rawPKeyPair struct {
	pub pubKey
	D   big.Int
}

func bigExp(x, y, n big.Int) big.Int {
	R := big.NewInt(int64(0))
	return *(R).Exp(&x, &y, &n)
}

func bigMul(a, b big.Int) big.Int {
	R := big.NewInt(int64(0))
	return *(R).Mul(&a, &b)
}

/**
  - this is effectively a message signed by us and granted to the other
  (((m1 ^ d1) mod n1) ^ e2) mod n2
*/
func apply(m big.Int, my rawPKeyPair, his pubKey) big.Int {
	return bigExp(bigExp(m, my.D, my.pub.N), his.E, his.N)
}

/** Generate the keypairs in a form we can do raw math with */
func rsaGen(randSource io.Reader, bitSize int) rawPKeyPair {
	kp, _ := rsa.GenerateKey(randSource, bitSize)
	return rawPKeyPair{
		pub: pubKey{
			N: *kp.N,
			E: *big.NewInt(int64(kp.E)),
		},
		D: *kp.D,
	}
}

func main() {
	randSource := &randBytes{}
	bitSize := 2048
	aliceTmp := rsaGen(randSource, bitSize)
	bobTmp := rsaGen(randSource, bitSize)
	fmt.Printf("Alice and Bob generate temporary RSA Keys\n")
	fmt.Printf("a: %v\nb: %v\n", aliceTmp, bobTmp)
	fmt.Printf("They exchange the public keys\n")
	//Alice and bob exchange public keys,
	//and they can use this to compute a pair of mutually inverse
	//keys that are not known to each other.

	//32byte secret key
	randKey := make([]byte, 32)
	randSource.Read(randKey)
	m := *big.NewInt(int64(91234))
	m.SetBytes(randKey)

	fmt.Printf("Alice and Bob use these to negotiate inverses with no public\n")
	fmt.Print("Alice makes up a giant key and encrypts it to bob\n")
	fmt.Printf("m:%v\n", m)
	m = apply(m, aliceTmp, bobTmp.pub)
	fmt.Printf("m:%v\n", m)
	m = apply(m, bobTmp, aliceTmp.pub)
	fmt.Printf("m:%v\n", m)
}
