package main

import (
	"bufio"
	"fmt"
	"os"
	"passwords-SDS/passcipher"
	"passwords-SDS/passwords"
	"passwords-SDS/users"
)

func main() {
	opt := "0"
	logged, user := setCredentials()

	for opt != "3" {
		if logged == true {
			opt = menu()
			switch opt {
			case "1":
				addPass(user)

			case "2":
				getPass(user)

			case "3":
				fmt.Println("\nSaliendo...")

			default:
				fmt.Println("Opcion no valida")
			}
		} else {
			fmt.Println("\nCredenciales no validas. O error al registrarse.\n")
			logged, user = setCredentials()
		}
	}
}

func menu() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nElige la opcion que desees")
	fmt.Println("1. Añadir password")
	fmt.Println("2. Consultar password")
	fmt.Println("3. Salir")
	fmt.Print("> ")
	scanner.Scan()
	opt := scanner.Text()
	return opt
}

func setCredentials() (bool, string) {
	opt := "0"
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bienvenido al gestor de contraseñas")
	fmt.Println("1. Iniciar sesion")
	fmt.Println("2. Crear usuario")

	for opt != "1" && opt != "2" {
		fmt.Print("> ")
		scanner.Scan()
		opt = scanner.Text()
		switch opt {
		case "1":
			ok, user := doLogin()
			return ok, user
		case "2":
			user := createUser()
			return true, user

		default:
			fmt.Println("Opcion no valida")
		}
	}

	return false, ""
}

func doLogin() (bool, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Introduce tu usuario: ")
	scanner.Scan()
	user := scanner.Text()

	fmt.Print("Introduce tu contraseña: ")
	scanner.Scan()
	pass := scanner.Text()

	ok := users.GetUser(user, pass)

	if ok {
		return true, user
	} else {
		return false, ""
	}
}

func createUser() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Introduce tu usuario: ")
	scanner.Scan()
	user := scanner.Text()

	fmt.Print("Introduce tu contraseña: ")
	scanner.Scan()
	pass := scanner.Text()

	passSt := users.CreatePass(user, pass)
	users.StoreUser(user, passSt)
	return user
}

func addPass(user string) {
	scanner := bufio.NewScanner(os.Stdin)
	cipher_pass := []byte(users.GetPassword(user))

	fmt.Print("Nombre del sitio web: ")
	scanner.Scan()
	site := scanner.Text()

	fmt.Print("Nombre de usuario del sitio web: ")
	scanner.Scan()
	login := scanner.Text()

	fmt.Print("Contraseña del sitio web: ")
	scanner.Scan()
	pass := passcipher.Cipher(scanner.Bytes(), cipher_pass)

	fmt.Print("Alguna nota o pista: ")
	scanner.Scan()
	notes := passcipher.Cipher(scanner.Bytes(), cipher_pass)

	entry := passwords.CreatePasswordEntry(login, pass, notes)
	passwords.StorePassword(user, site, entry)

}

func getPass(user string) {
	scanner := bufio.NewScanner(os.Stdin)
	cipher_pass := []byte(users.GetPassword(user))

	fmt.Println("\nSitios almacenados:")
	passwords.ListSites(user)

	fmt.Print("\nIntroduce el site que quieres consultar:")
	scanner.Scan()
	site := scanner.Text()

	entry := passwords.GetPassword(user, site)

	pass := passcipher.Decipher(entry.Pass, cipher_pass)
	notes := passcipher.Decipher(entry.Notes, cipher_pass)

	fmt.Printf("Sitio %s\n", site)
	fmt.Printf("Login: %s\n", entry.Login)
	fmt.Printf("Password: %s\n", string(pass))
	fmt.Printf("Notas: %s\n", string(notes))
}
