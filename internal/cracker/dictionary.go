package cracker

import (
	"bufio"
	"os"
	"password-cracker/internal/hash"
)

func DictionaryAttack(hashFunc hash.HashFunction, targetHash string, dictionaryPath string) string {
	file, err := os.Open(dictionaryPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := scanner.Text()
		if hashFunc(password) == targetHash {
			return password
		}
	}
	return ""
}
