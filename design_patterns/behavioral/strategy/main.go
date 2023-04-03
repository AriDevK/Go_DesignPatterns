package main

import (
	"Go_DesignPatterns_Platzi/design_patterns/behavioral/strategy/algorithms"
)

func main() {
	user := "Ariadne"
	pwd := "my ultra hiper galaxy secret password 1234"

	pwdProtector := algorithms.NewPasswordProtector(user, pwd, &algorithms.Sha{})
	pwdProtector.Hash()

	pwdProtector.SetHashAlgorithm(&algorithms.SHA256{})
	pwdProtector.Hash()

	pwdProtector.SetHashAlgorithm(&algorithms.MD5{})
	pwdProtector.Hash()
}
