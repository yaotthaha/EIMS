package pbkdf2

import (
	"fmt"
	"testing"
	"time"
)

func TestPBKDF2(T *testing.T) {
	rawPassword := `admin`
	t1 := time.Now()
	encPassword, err := EncryptPwd(rawPassword)
	t2 := time.Now()
	if err != nil {
		T.Error(err)
		return
	}
	fmt.Println(encPassword)
	fmt.Println(len(encPassword))
	fmt.Println(t2.Sub(t1))
	fmt.Println(CheckPasswordMatch(rawPassword, encPassword))
}
