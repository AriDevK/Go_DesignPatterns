package algorithms

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Sha struct {
}

func (Sha) Hash(p *PasswordProtector) {
	h := sha1.New()
	h.Write([]byte(p.passwordName))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using SHA for %s: %s\n", p.user, sha1Hash)
}
