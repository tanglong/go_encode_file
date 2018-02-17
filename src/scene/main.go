package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func decrypt(cipherstring string, keystring string) string {
	// Byte array of the string
	ciphertext := []byte(cipherstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func encrypt(plainstring, keystring string) string {
	// Byte array of the string
	plaintext := []byte(plainstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func main() {
	fileTest()
}

func stringTest() {
	key := "testtesttesttest"
	string1 := "isdfadfasdfasdfa"
	ciphertext := encrypt(string1, key)

	fmt.Printf("after encrypt:[%s]\n", ciphertext)

	string3 := decrypt(ciphertext, key)
	fmt.Printf("after decrypt:[%s]\n", string3)
}

func fileTest() {
	key := "testtesttesttest"

	dat, err := ioutil.ReadFile("11111")
	if err != nil {
		fmt.Printf("read file err!file name:%s,err:%v \n", "11111", err)
		return
	}

	string1 := string(dat)
	fmt.Print(string(dat))

	ciphertext := encrypt(string1, key)
	fmt.Printf("after encrypt:[%s]\n", ciphertext)

	string3 := decrypt(ciphertext, key)
	fmt.Printf("after decrypt:[%s]\n", string3)

}
