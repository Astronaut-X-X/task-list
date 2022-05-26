package util

import (
	"crypto/md5"
	"fmt"
)

// MD5 encryption for the code.
func MD5Encipher(code string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(code)))
}
