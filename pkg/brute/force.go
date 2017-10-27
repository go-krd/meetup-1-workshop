package brute

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Force bruteforces password from hash.
func Force(hash string) (password string, err error) {
	b, err := hex.DecodeString(hash)
	if err != nil {
		return "", fmt.Errorf("decode hex: %s", err)
	}

	var target [md5.Size]byte
	copy(target[:], b)

	source := []byte("00000")
	for md5.Sum(source) != target {
		NextWord(source)

		if bytes.Equal(source, []byte("00000")) {
			return "", fmt.Errorf("password not found")
		}
	}

	return string(source), nil
}
