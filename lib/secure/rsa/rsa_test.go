package rsa

import (
	"fmt"
	"testing"
)

func TestGenerateAndEncAndDec(T *testing.T) {
	message := "test"
	privateKey, publicKey, err := GenerateRSA()
	if err != nil {
		T.Fatal(err)
	}
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	encrypted, err := Encrypt(publicKey, message)
	if err != nil {
		T.Fatal(err)
	}
	fmt.Println(encrypted)
	decrypted, err := Decrypt(privateKey, encrypted)
	if err != nil {
		T.Fatal(err)
	}
	fmt.Println(decrypted)
}

func TestEncrypt(T *testing.T) {
	rawPassword := `admin`
	publicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDNSa+zzgPycrpzV/v0sgVASqjM
ZdVeMr1nVrpzo5WIIZ1wtDg0jtQi21XXKWAn7BNgoq1Tf8W17FriPv5XkAxt8gQg
yTh8oMYmyDxQVbIv/6GLOHO1W3Sui+96SKSY4H4LIHu4dGdOYdGBNaRpEkL13YkL
3ZULo5UOWJLo/yT06wIDAQAB
-----END PUBLIC KEY-----
`
	encPassword, err := Encrypt(publicKey, rawPassword)
	if err != nil {
		T.Error(err)
		return
	}
	fmt.Println(encPassword)
}
