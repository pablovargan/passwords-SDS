//package sal_hash
package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

const (
	directory string = "directory"
)

type Pass struct {
	Sal         []byte `json:"sal"`
	PasswordSal string `json:"passwordSal"`
}

type User map[string]Pass

func convertPointer(p *Pass) (n int, pValue Pass) {
	pValue = Pass{p.Sal, p.PasswordSal}
	return len(pValue.PasswordSal), pValue
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func MakeSal(sal *[]byte) {
	*sal = make([]byte, 16)
	_, err := rand.Read(*sal)
	check(err)
}

func CreatePass(user string, password string) Pass {
	if user == "" || password == "" {
		log.Fatal("User/Password is null")
	}

	pUser := new(Pass)
	MakeSal(&pUser.Sal)
	pPass := []byte(password)

	tmp := make([]byte, len(pUser.Sal)+len(pPass))
	for i := 0; i < len(pPass); i++ {
		tmp = append(pUser.Sal, pPass[i])
	}

	hasher := sha512.New()
	hasher.Reset()
	_, err := hasher.Write(tmp)
	check(err)

	pUser.PasswordSal = base64.URLEncoding.EncodeToString(hasher.Sum(tmp))
	length, pToValue := convertPointer(pUser)
	if length != len(pUser.PasswordSal) {
		log.Fatal("Error converting pointer to value")
	}

	return pToValue
}

func StoreUser(email string, pass Pass) {
	//var users [...]User
	bytes, err := ioutil.ReadFile(directory)
	check(err)
	json.Unmarshal(bytes, &users)

	var u User
	u = make(map[string]Pass)
	u[email] = pass

	//bytes, err = json.Marshal(users)
	//u = pass
	//bytes, err = json.Marshal(u)
	//ioutil.WriteFile(user, bytes, 0666)
}

/*u := new(User)
	bytes, err := ioutil.ReadFile(user)
	if err != nil {
		u.Id = make(map[string]Pass)
	}
	/*
		 Unmarshal parses the JSON-encoded data and
		   stores the result in the value pointed to by 'bytes'.

		length, pToValue := convertPointer(pUser)
		if length != len(pUser.passwordSalt) {
			log.Fatal("Error converting pointer to value")
		}
		u.Id[user] = pToValue
		err = json.Unmarshal(bytes, &u)
		check(err)
		/*length, pToValue := convertPointer(pUser)
		if length != len(pUser.passwordSalt) {
			log.Fatal("Error converting pointer to value")
		}
		u.Id[user] = pToValue

		bytes, err = json.Marshal(u)
		ioutil.WriteFile(user, bytes, 0666)

		/*
			u := new(User)
			u.Id = make(map[string]Pass)

			length, pToValue := convertPointer(pUser)
			if length != len(pUser.passwordSalt) {
				log.Fatal("Error converting pointer to value")
			}
			u.Id[user] = pToValue
}
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	u, _ := reader.ReadString('\n')
	username := strings.TrimSpace(u)

	fmt.Print("Enter Password: ")
	p, _ := reader.ReadString('\n')
	password := strings.TrimSpace(p)

	fmt.Println(username)
	fmt.Println(password)
	StoreUser(username, CreatePass(username, password))
}
