package pbkdf2

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	mathrand "math/rand"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltMinLen = 8
	saltMaxLen = 32
	iter       = 1000
	keyLen     = 64
)

func EncryptPwd(rawPassword string) (string, error) {
	salt := make([]byte, mathrand.Intn(saltMaxLen-saltMinLen)+saltMinLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	encPassword := encryptPwdWithSalt([]byte(rawPassword), salt)
	encPassword = append(encPassword, salt...)
	return base64.StdEncoding.EncodeToString(encPassword), nil
}

func encryptPwdWithSalt(pwd, salt []byte) []byte {
	pwd = append(pwd, salt...)
	return pbkdf2.Key(pwd, salt, iter, keyLen, sha256.New)
}

func CheckPasswordMatch(rawPassword, encPassword string) bool {
	if len(encPassword) == 0 || len(rawPassword) == 0 {
		return false
	}
	encPasswordDecode, err := base64.StdEncoding.DecodeString(encPassword)
	if err != nil {
		return false
	}
	salt := encPasswordDecode[keyLen:]
	encPasswordBase64 := base64.StdEncoding.EncodeToString(encPasswordDecode[0:keyLen])
	encPasswordNowBase64 := base64.StdEncoding.EncodeToString(encryptPwdWithSalt([]byte(rawPassword), salt))
	return encPasswordBase64 == encPasswordNowBase64
}
