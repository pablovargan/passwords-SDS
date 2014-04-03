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

func main() {

	m := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	password, _ := reader.ReadString('\n')

	key := strings.TrimSpace(username)
	value := strings.TrimSpace(password)
	m[key] = value

	// Generates 90 bytes random numbers
	salt := make([]byte, 90)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	pass := []byte(value)
	// Append pass to salt
	for i := 0; i < len(pass); i++ {
		salt = append(salt, pass[i])
	}
	// Hash complete salt(salt+password)
	hasher := sha512.New()
	hasher.Write(salt)
	valueSha := base64.URLEncoding.EncodeToString(hasher.Sum(salt))
	// Key -> User, Value = (salt+password) to SHA512
	m[key] = valueSha
	// Enjoy!
	fmt.Println(m[key])
}
