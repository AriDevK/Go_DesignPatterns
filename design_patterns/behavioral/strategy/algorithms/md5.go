package algorithms

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type MD5 struct {
}

func (MD5) Hash(p *PasswordProtector) {
	h := md5.New()
	h.Write([]byte(p.passwordName))
	md5Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using MD5 for %s: %s\n", p.user, md5Hash)
}
