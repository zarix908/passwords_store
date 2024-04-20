package key

import (
	"bytes"
	"fmt"

	"filippo.io/age"
)

type KeyGenerator interface {
	GenerateEncrypted(password string) ([]byte, error)
}

func NewKeyGenerator() KeyGenerator {
	return &keyGenerator{}
}

type keyGenerator struct{}

func (g *keyGenerator) GenerateEncrypted(pass string) ([]byte, error) {
	recipient, err := age.NewScryptRecipient(pass)
	if err != nil {
		return nil, fmt.Errorf("failed to get new age recipient: %w", err)
	}

	buf := bytes.NewBuffer(nil)

	encryptor, err := age.Encrypt(buf, recipient)
	if err != nil {
		return nil, fmt.Errorf("failed to create encryptor: %w", err)
	}

	key, err := age.GenerateX25519Identity()
	if err != nil {
		return nil, fmt.Errorf("failed to generate age key: %w", err)
	}

	fmt.Fprint(encryptor, key)

	return buf.Bytes(), nil
}
