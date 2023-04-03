package algorithms

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm IHashAlgorithm
}

func NewPasswordProtector(user string, passwordName string, hashAlgorithm IHashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hashAlgorithm,
	}
}

func (p *PasswordProtector) SetPasswordName(passwordName string) {
	p.passwordName = passwordName
}

func (p *PasswordProtector) SetHashAlgorithm(hash IHashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}
