package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func GenerateRSA() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", err
	}
	publicKey := privateKey.PublicKey
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return "", "", err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}
	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDer,
	}
	publicKeyPem := string(pem.EncodeToMemory(publicBlock))
	privateKeyPem := string(pem.EncodeToMemory(privateBlock))
	return privateKeyPem, publicKeyPem, nil
}

func Encrypt(publicKeyPem string, plainText string) (string, error) {
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyPem))
	if publicKeyBlock == nil {
		return "", fmt.Errorf("public key invalid")
	}
	publicKeyAny, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return "", err
	}
	publicKey, ok := publicKeyAny.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("public key invalid")
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(privateKeyPem string, plainText string) (string, error) {
	plainTextBytes, err := base64.StdEncoding.DecodeString(plainText)
	if err != nil {
		return "", err
	}
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyPem))
	if privateKeyBlock == nil {
		return "", fmt.Errorf("private key invalid")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, plainTextBytes)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
