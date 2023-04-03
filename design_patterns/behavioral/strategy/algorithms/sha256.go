package algorithms

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type SHA256 struct{}

func (SHA256) Hash(p *PasswordProtector) {
	h := sha256.New()
	h.Write([]byte(p.passwordName))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using SHA256 for %s: %s\n", p.user, sha256Hash)
}
