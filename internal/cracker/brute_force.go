package cracker

import (
	"password-cracker/internal/hash"
)

func BruteForce(hashFunc hash.HashFunction, targetHash string, charset string, maxLength int) string {
	for length := 1; length <= maxLength; length++ {
		if result := tryCombinations(hashFunc, targetHash, "", charset, length); result != "" {
			return result
		}
	}
	return ""
}

func tryCombinations(hashFunc hash.HashFunction, targetHash, prefix, charset string, length int) string {
	if length == 0 {
		if hashFunc(prefix) == targetHash {
			return prefix
		}
		return ""
	}

	for _, char := range charset {
		if result := tryCombinations(hashFunc, targetHash, prefix+string(char), charset, length-1); result != "" {
			return result
		}
	}
	return ""
}
