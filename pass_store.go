package passStore

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Almacena el byte de inicio de un pass y su tamaño.
type Direccion struct {
	init   int64
	length int
}

//Comprueba si se ha producido error, añadiendo una entrada al log en caso afirmativo.
func check_error(desc string, err error) {
	if err != nil {
		log.Fatal(desc, err)
	}
}

//Almacena un nuevo password actualizando el directorio.
func storePass(user string, site string, pass string) {
	dir := getDirFile(user)

	mf, err := os.Create(user + "_main")
	defer mf.Close()
	check_error("opening main file:", err)

	var d Direccion

	stats, err := mf.Stat()
	check_error("retriving main file stats", err)

	d.init = stats.Size()

	writed, err := mf.Write([]byte(pass))
	check_error("writing main:", err)

	d.length = writed
	dir[site] = d
	writeDirFile(user, dir)
}

//Devuelve un pass concreto.
func getPass(user string, site string) string {
	directory := getDirFile(user)
	dir := directory[site]

	mf, err := os.Open(user + "_main")
	check_error("opening for reading main file", err)

	mf.Seek(dir.init, 0)
	b := make([]byte, dir.length)
	mf.Read(b)

	return string(b)
}

//Devuelve el contenido del fichero de directorio del usuario.
func getDirFile(user string) map[string]Direccion {
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
func writeDirFile(user string, dir map[string]Direccion) {
	f, err := os.Create(user + "_directory")
	check_error("opening directory file", err)
	defer f.Close()

	b, err := json.Marshal(dir)
	check_error("encoding dir", err)

	f.Write(b)
}

/*func main() {
	storePass("panfri", "petardas.com", "1234")
	fmt.Println(getPass("panfri", "petardas.com"))
}*/
