package auth

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "123456"
	hash, _ := Encrypt(password)

	fmt.Println("password:", password)
	fmt.Println("hash:", hash)
}

func TestCompare(t *testing.T) {
	password := "123456"
	hash, _ := Encrypt(password)

	match := Compare(password, hash)
	fmt.Println("match:", match)

}