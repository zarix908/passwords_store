package key

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"filippo.io/age"
)

func DecryptKey(pass string, data []byte) (*age.X25519Identity, error) {
	identity, err := age.NewScryptIdentity(pass)
	if err != nil {
		return nil, fmt.Errorf("failed to get age identity: %w", err)
	}

	buf := bytes.NewBuffer(data)
	decrypter, err := age.Decrypt(buf, identity)
	if err != nil {
		return nil, fmt.Errorf("failed to create age decrypter: %w", err)
	}

	rawKey, err := io.ReadAll(decrypter)
	if err != nil {
		return nil, errors.New("failed to decrypt key")
	}

	key, err := age.ParseX25519Identity(string(rawKey))
	if err != nil {
		return nil, errors.New("failed to parse key")
	}

	return key, nil
}