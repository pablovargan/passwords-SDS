package passcipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

/*
Cifra una cadena. Devuelve la cadena cifrada y el vector de inicializacion.
*/
func Cipher(plain_text []byte, cipher_pass []byte) []byte {
	block, err := aes.NewCipher(cipher_pass)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plain_text))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plain_text)

	return ciphertext
}

/*
Descifra una cadena. Devuelve la cadena descifrada.
*/
func Decipher(ciphertext []byte, cipher_pass []byte) []byte {

	//Falta validar que el password de cifrado sea correcto.

	block, err := aes.NewCipher(cipher_pass)
	if err != nil {
		panic(err)
	}

	plain_text := make([]byte, len(ciphertext[aes.BlockSize:]))
	stream := cipher.NewCTR(block, ciphertext[:aes.BlockSize])
	stream.XORKeyStream(plain_text, ciphertext[aes.BlockSize:])

	return plain_text
}
