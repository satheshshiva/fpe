package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/capitalone/fpe/ff1"
)

//export EncryptDecrypt
func EncryptDecrypt(ipText *C.char) *C.char {
	fmt.Println("ipText", C.GoString(ipText))
	key, err := hex.DecodeString("2B7E151628AED2A6ABF7158809CF4F3CEF4359D8D580AA4F7F036D6F04FC6A94")

	tweak, err := hex.DecodeString("")

	// 16 is an arbitrary number for maxTlen
	cipher, err := ff1.NewCipher(36, 16, key, tweak)
	if err != nil {
		panic(fmt.Errorf("Go: Unable to create cipher: %v", err))
	}

	plaintext := "xs8a0azh2avyalyzuwdxs8a0azh2avyalyzuwdxs8a0azh2avyalyzuwdxs8a0azh2avyalyzuwdxs8a0azh2avyalyzuwdxs8a0azh2avyalyzuwdxs8a0azh2avyal"

	fmt.Println("GO: INPUT TEXT: ", plaintext)
	ciphertext, err := cipher.Encrypt(plaintext)
	if err != nil {
		panic(fmt.Errorf("Unable to create cipher: %v", err))
	}

	fmt.Println("GO: CIPHER TEXT: ", ciphertext)
	decrypted, err := cipher.Decrypt(ciphertext)
	if err != nil {
		panic(fmt.Errorf("Go: Unable to create cipher: %v", err))
	}

	if plaintext != decrypted {
		panic(fmt.Errorf("Go: Long Decrypt Failed. \n Expected: %v \n Got: %v \n", plaintext, decrypted))

	}
	fmt.Println("GO: PLAIN = DECRYPTED: VALIDATED")

	returnVal := C.CString("AASD")
	//defer C.free(unsafe.Pointer(returnVal))
	return returnVal

}

//export Encrypt
func Encrypt(plaintext_ *C.char) *C.char {
	plaintext := C.GoString(plaintext_)
	key, err := hex.DecodeString("2B7E151628AED2A6ABF7158809CF4F3CEF4359D8D580AA4F7F036D6F04FC6A94")

	tweak, err := hex.DecodeString("")

	// 16 is an arbitrary number for maxTlen
	cipher, err := ff1.NewCipher(36, 16, key, tweak)
	if err != nil {
		panic(fmt.Errorf("Go: calling NewCipher fail: %v", err))
	}

	fmt.Println("GO: INPUT TEXT: ", plaintext)
	ciphertext, err := cipher.Encrypt(plaintext)
	if err != nil {
		panic(fmt.Errorf("Go: Calling Encrypt fail:  %v", err))
	}
	returnVal := C.CString(ciphertext)
	//defer C.free(unsafe.Pointer(returnVal))
	return returnVal
}

//export Decrypt
func Decrypt(ciphertext_ *C.char) *C.char {
	ciphertext := C.GoString(ciphertext_)

	key, err := hex.DecodeString("2B7E151628AED2A6ABF7158809CF4F3CEF4359D8D580AA4F7F036D6F04FC6A94")

	tweak, err := hex.DecodeString("")

	// 16 is an arbitrary number for maxTlen
	cipher, err := ff1.NewCipher(36, 16, key, tweak)
	if err != nil {
		panic(fmt.Errorf("Go: Unable to create cipher: %v", err))
	}

	decrypted, err := cipher.Decrypt(ciphertext)
	if err != nil {
		panic(fmt.Errorf("Go: Unable to create cipher: %v", err))
	}

	returnVal := C.CString(decrypted)
	//defer C.free(unsafe.Pointer(returnVal))
	return returnVal
}

func main() {

}
