package passStore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//Almacena el byte de inicio de un pass y su tamaño.
type Direccion struct {
	Init   int64 `json:"init"`
	Length int   `json:"length"`
}

//Comprueba si se ha producido error, añadiendo una entrada al log en caso afirmativo.
func Check_error(desc string, err error) {
	if err != nil {
		log.Fatal(desc, err)
	}
}

//Almacena un nuevo password actualizando el directorio.
func StorePass(user string, site string, pass string) {
	dir := GetDirFile(user)
	var mf *os.File
	fileName := user + "_main"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		mf, _ = os.Create(fileName)
		defer mf.Close()
	}

	mf, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	Check_error("opening main file to write", err)
	defer mf.Close()

	var d Direccion

	stats, err := mf.Stat()
	Check_error("retriving main file stats", err)

	d.Init = stats.Size()
	mf.Seek(d.Init, 0)

	writed, err := mf.Write([]byte(pass))
	Check_error("writing main:", err)

	d.Length = writed
	dir[site] = d

	WriteDirFile(user, dir)
}

//Devuelve un pass concreto.
func GetPass(user string, site string) string {
	directory := GetDirFile(user)
	dir := directory[site]

	mf, err := os.Open(user + "_main")
	Check_error("opening for reading main file", err)

	mf.Seek(dir.Init, 0)
	b := make([]byte, dir.Length)
	mf.Read(b)

	return string(b)
}

//Devuelve el contenido del fichero de directorio del usuario.
func GetDirFile(user string) map[string]Direccion {
	dirByte, err := ioutil.ReadFile(user + "_directory")

	//Si el fichero no existe aun, devuelvo un directorio vacio.
	if err != nil {
		return make(map[string]Direccion)
	}
	var dir map[string]Direccion
	json.Unmarshal(dirByte, &dir)

	return dir
}

//Escribe el directorio actualizado en su fichero correspondiente.
func WriteDirFile(user string, dir map[string]Direccion) {
	f, err := os.Create(user + "_directory")
	Check_error("opening directory file", err)
	defer f.Close()

	b, err := json.Marshal(dir)
	Check_error("encoding dir", err)

	f.Write(b)
}
