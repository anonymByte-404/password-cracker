package utils

import (
	"crypto/md5"
	"fmt"
)

func HashString(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	hashedBytes := hash.Sum(nil)
	return fmt.Sprintf("%x", hashedBytes)
}
