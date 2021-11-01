package main

import "fmt"

type PasswordProctect struct {
	user          string
	PasswordName  string
	hashAlgorithm HashAlgorithm
}
type HashAlgorithm interface {
	Hash(p *PasswordProctect)
}

func newPasswordProctect(user string, passwordName string, hashAlgorithm HashAlgorithm) *PasswordProctect {
	return &PasswordProctect{user, passwordName, hashAlgorithm}
}
func (p *PasswordProctect) setHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}
func (p *PasswordProctect) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct {
}

func (SHA) Hash(p *PasswordProctect) {
	fmt.Println("Hasing using SHA for ",p.PasswordName)
}
type MD5 struct {
}

func (MD5) Hash(p *PasswordProctect) {
	fmt.Println("Hasing using MD5 for ",p.PasswordName)
}

func main() {
	sha := SHA{}
	md5 := MD5{}	
	passwordProctect := newPasswordProctect("user", "password", sha)
	passwordProctect.Hash()
	passwordProctect.setHashAlgorithm(md5)
	passwordProctect.Hash()
}