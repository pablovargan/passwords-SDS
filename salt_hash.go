package salt_hash

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"log"
)

type Pass struct {
	salt         []byte
	passwordSalt string
}

type User struct {
	Email string
	// Map with string as key and Pass as value
	Id map[string]Pass
}

func convertPointer(p *Pass) (n int, pValue Pass) {
	pValue = Pass{p.salt, p.passwordSalt}
	return len(pValue.passwordSalt), pValue
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func MakeSalt(salt *[]byte) {
	*salt = make([]byte, 16)
	_, err := rand.Read(*salt)
	check(err)
}

func CreateUser(user string, password string) {
	if user == "" || password == "" {
		log.Fatal("User/Password does not null")
	}

	pUser := new(Pass)
	MakeSalt(&pUser.salt)
	pPass := []byte(password)

	tmp := make([]byte, len(pUser.salt)+len(pPass))
	for i := 0; i < len(pPass); i++ {
		tmp = append(pUser.salt, pPass[i])
	}

	hasher := sha512.New()
	hasher.Reset()
	_, err := hasher.Write(tmp)
	check(err)

	pUser.passwordSalt = base64.URLEncoding.EncodeToString(hasher.Sum(tmp))

	// STORE?

	u := new(User)
	u.Email = user
	u.Id = make(map[string]Pass)

	length, pToValue := convertPointer(pUser)
	if length != len(pUser.passwordSalt) {
		log.Fatal("Error converting pointer to value")
	}
	u.Id[u.Email] = pToValue
}