package main

import (
	"fmt"
	"password-cracker/internal/cracker"
	"password-cracker/internal/hash"
	"password-cracker/internal/utils"
	"time"
)

func main() {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$&'()*+,-./:;<=>?@[]^_`{|}~"

	var attackType string
	var input string

	fmt.Println("Type the password: ")
	fmt.Scanln(&input)
	targetHash := utils.HashString(input)

	fmt.Println("Choose an attack type (brute-force/dictionary): ")
	fmt.Scanln(&attackType)

	switch attackType {
	case "brute-force":
		fmt.Println("Starting brute-force attack...")
		start := time.Now()
		password := cracker.BruteForce(hash.MD5, targetHash, charset, 8)
		elapse := time.Since(start)
		if password != "" {
			fmt.Printf("Password found: %s\n", password)
		} else {
			fmt.Println("Password not found.")
		}
		fmt.Printf("Brute-force attack took: %s\n", elapse)

	case "dictionary":
		fmt.Println("Starting dictionary attack...")
		start := time.Now()
		password := cracker.DictionaryAttack(hash.MD5, targetHash, "dictionary.txt")
		elapse := time.Since(start)
		if password != "" {
			fmt.Printf("Password found: %s\n", password)
		} else {
			fmt.Println("Password not found.")
		}
		fmt.Printf("Dictionary attack took: %s\n", elapse)

	default:
		fmt.Println("Invalid attack type. Please choose 'brute-force' or 'dictionary'.")
	}
}
