package main

import (
	"bufio"
	//"crypto/sha512"
	//"encoding/base64"
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

	//bv := []byte(key)
	//hasher := sha512.New()
	//hasher.Sum(bv)

	//fmt.Print(hasher)
	//toString := base64.URLEncoding.EncodeToString(hasher.Sum(bv))
	//fmt.Print(toString)
}
