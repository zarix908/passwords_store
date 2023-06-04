package manager

import (
	"bytes"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"

	"filippo.io/age"
)

const keyFile = "key"

func GenerateMasterKey() error {
	_, err := os.Stat(keyFile)
	switch {
	case err == nil:
		fmt.Println("Already initialized. Master key file exists.")
		return nil
	case errors.Is(err, os.ErrNotExist):
	default:
		return fmt.Errorf("failed to check existence of master key file: %w", err)
	}

	pass, err := readPassWithConfirm()
	if err != nil {
		return fmt.Errorf("failed to read password with confirmation: %w", err)
	}

	recipient, err := age.NewScryptRecipient(pass)
	if err != nil {
		return fmt.Errorf("failed to get new age recipient: ")
	}
	fmt.Println("Password accepted.")
	fmt.Println()

	fmt.Println("Generating master key...")

	buf := bytes.NewBuffer(nil)

	encryptor, err := age.Encrypt(buf, recipient)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	key, err := age.GenerateX25519Identity()
	if err != nil {
		return fmt.Errorf("failed to generate age key: %w", err)
	}

	fmt.Fprint(encryptor, key)

	file, err := os.OpenFile(keyFile, os.O_CREATE|os.O_WRONLY, 0400)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if err := pem.Encode(file, &pem.Block{
		Type:  "STORE KEY",
		Bytes: buf.Bytes(),
	}); err != nil {
		return fmt.Errorf("failed to encode to pem format: %w", err)
	}

	return nil
}

func restoreKey() (*age.X25519Identity, error) {
	pemKey, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	pemKeyBlock, _ := pem.Decode(pemKey)
	pemKeyBuf := bytes.NewBuffer(pemKeyBlock.Bytes)

	pass, err := readPassWithConfirm()
	switch {
	case errors.Is(err, os.ErrNotExist):
		fmt.Println("Master key file doesn't exist. Maybe you need to run initialization command first.")
	case err != nil:
	default:
		return nil, fmt.Errorf("failed to read password with confirmation: %w", err)
	}

	identity, err := age.NewScryptIdentity(pass)
	if err != nil {
		return nil, fmt.Errorf("failed to get age identity: %w", err)
	}

	decrypter, err := age.Decrypt(pemKeyBuf, identity)
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
