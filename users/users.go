package users

import (
	//"bufio"
	"crypto/rand"
	"crypto/sha256"
	//"encoding/base64"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	//"os"
	//"strings"
)

const (
	directory string = "directory"
)

type Pass struct {
	Sal         []byte `json:"sal"`
	PasswordSal []byte `json:"passwordSal"`
}

/*func convertPointer(p *Pass) (n int, pValue Pass) {
	pValue = Pass{p.Sal, p.PasswordSal}
	return len(pValue.PasswordSal), pValue
}*/

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func createHash(sal []byte, pass []byte) []byte {

	tmp := make([]byte, len(sal)+len(pass))

	copy(tmp[:16], sal)
	copy(tmp[16:], pass)

	hasher := sha256.New()
	hasher.Reset()
	_, err := hasher.Write(tmp)
	check(err)

	resume := hasher.Sum(nil)

	return resume
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

	//pUser := new(Pass)
	var pUser Pass
	MakeSal(&pUser.Sal)

	pUser.PasswordSal = createHash(pUser.Sal, []byte(password))

	//Codigo simplicado.

	//length, pToValue := convertPointer(pUser)
	/*if length != len(pUser.PasswordSal) {
		log.Fatal("Error converting pointer to value")
	}*/

	//return pToValue
	return pUser
}

func StoreUser(email string, pass Pass) {
	var warehouse map[string]Pass

	bytes, err := ioutil.ReadFile(directory)
	if err != nil {
		warehouse = make(map[string]Pass)
	}
	json.Unmarshal(bytes, &warehouse)

	warehouse[email] = pass
	bytes, err = json.Marshal(warehouse)
	ioutil.WriteFile(directory, bytes, 0666)
}

func GetUser(email string, password string) bool {
	var warehouse map[string]Pass

	bytes, err := ioutil.ReadFile(directory)
	if err != nil {
		log.Fatal("The file does not exist")
	}
	json.Unmarshal(bytes, &warehouse)

	// Check if email exists
	data := warehouse[email]
	if data.Sal == nil {
		return false
	}
	// Get password+sal generated
	var passSaltGen = createHash(data.Sal, []byte(password))
	if string(passSaltGen) == string(data.PasswordSal) {
		return true
	} else {
		return false
	}
}

//Return password of specified user.
func GetPassword(email string) []byte {
	var warehouse map[string]Pass

	bytes, err := ioutil.ReadFile(directory)
	if err != nil {
		log.Fatal("The file does not exist")
	}
	json.Unmarshal(bytes, &warehouse)

	return warehouse[email].PasswordSal
}
