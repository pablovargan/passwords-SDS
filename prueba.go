package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func escribirFichero(json []byte) {
	f, err := os.Create("/tmp/contrasenyas")
	check(err)

	f.Write(json)
}

func codificarJSON(mapa map[string]string) []byte {
	b, err := json.Marshal(mapa)
	check(err)

	return b
}

func decodificarJSON(coded []byte) map[string]string {
	mapa := make(map[string]string)

	json.Unmarshal(coded, &mapa)
	return mapa
}

func leerFichero() []byte {
	leido, err := ioutil.ReadFile("/tmp/contrasenyas")
	check(err)

	return leido
}

func main() {
	m := map[string]string{"petordos.com": "4567", "flaisbook.com": "1234"}

	//os.Create("/tmp/contrasenyas")

	escribirFichero(codificarJSON(m))

	mapa := decodificarJSON(leerFichero())

	fmt.Println("Claves almacenadas")
	for k, _ := range mapa {
		fmt.Println(k)
	}
}
