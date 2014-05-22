package passwords

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Password struct {
	Login string `json:"login"`
	Pass  []byte `json:"password"`
	Notes []byte `json:"notes"`
}

func CreatePasswordEntry(login string, pass []byte, notes []byte) Password {
	var entry Password
	entry.Login = login
	entry.Pass = pass
	entry.Notes = notes

	return entry
}

func StorePassword(user string, site string, entry Password) {
	var warehouse map[string]Password

	bytes, err := ioutil.ReadFile(user)

	if err != nil {
		warehouse = make(map[string]Password)
	}

	json.Unmarshal(bytes, &warehouse)

	warehouse[site] = entry

	bytes, err = json.Marshal(warehouse)

	ioutil.WriteFile(user, bytes, 0666)
}

func GetPassword(user string, site string) Password {
	var warehouse map[string]Password

	warehouse = getWarehouse(user)
	return warehouse[site]
}

func ListSites(user string) {
	warehouse := getWarehouse(user)

	for site, _ := range warehouse {
		fmt.Println(site)
	}
}

func getWarehouse(user string) map[string]Password {
	var warehouse map[string]Password

	bytes, err := ioutil.ReadFile(user)
	if err != nil {
		log.Fatal("The file does not exist")
	}

	json.Unmarshal(bytes, &warehouse)

	return warehouse
}
