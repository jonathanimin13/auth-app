package bcrypt

import "golang.org/x/crypto/bcrypt"

type Bcrypt interface {
	GenerateFromPassword(password []byte) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type bcryptImpl struct{}

func NewBcrypt() Bcrypt {
	return bcryptImpl{}
}

func (b bcryptImpl) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (b bcryptImpl) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}