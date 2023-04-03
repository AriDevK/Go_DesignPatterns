package algorithms

type IHashAlgorithm interface {
	Hash(p *PasswordProtector)
}
