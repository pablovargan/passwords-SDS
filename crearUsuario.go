package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

type Pass struct {
	salt         []byte
	passwordSalt string
}

type User struct {
	Email string
	// Hashmap with string as key and Pass as value
	Id map[string]Pass
}

func convert(p *Pass) (n int, pValue Pass) {
	pValue = Pass{p.salt, p.passwordSalt}
	return len(pValue.passwordSalt), pValue
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create user
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	u, _ := reader.ReadString('\n')
	username := strings.TrimSpace(u)

	fmt.Print("Enter Password: ")
	p, _ := reader.ReadString('\n')
	password := strings.TrimSpace(p)

	// Make a salt
	pStruct := new(Pass)
	pStruct.salt = make([]byte, 16)
	_, err := rand.Read(pStruct.salt)
	check(err)
	// Convert password to bytes
	pS := []byte(password)
	// Append pS with salt
	tmp := make([]byte, 16+len(pS))
	for i := 0; i < len(pS); i++ {
		tmp = append(pStruct.salt, pS[i])
	}
	// 'Hashing' time :)
	hasher := sha512.New()
	hasher.Reset()
	_, err = hasher.Write(tmp)
	check(err)
	// Set to passwordSalt the hashed tmp
	pStruct.passwordSalt = base64.URLEncoding.EncodeToString(hasher.Sum(tmp))

	// New instance of user
	user := new(User)
	user.Email = username
	user.Id = make(map[string]Pass)

	// Convert the pStruct pointer to value ;)
	length, pToValue := convert(pStruct)
	// Check if the conversion was correct
	if length != len(pStruct.passwordSalt) {
		panic(length)
	}
	user.Id[user.Email] = pToValue
	fmt.Println(user)
}
