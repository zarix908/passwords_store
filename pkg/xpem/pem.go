package xpem

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"os"
)

const PassHeimdallPEMType = "PASSHEIMDAL KEY"

func CreatePEMFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0400)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if err := pem.Encode(file, &pem.Block{
		Type:  PassHeimdallPEMType,
		Bytes: data,
	}); err != nil {
		return fmt.Errorf("failed to encode to pem format: %w", err)
	}

	return nil
}

func ReadPEMFile(path string) ([]byte, error) {
	pemKey, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pemKeyBlock, _ := pem.Decode(pemKey)
	if pemKeyBlock.Type != PassHeimdallPEMType {
		return nil, fmt.Errorf(
			"failed to read PEM: type %s is incompatible, type %s is required", 
			pemKeyBlock.Type, PassHeimdallPEMType,
		)
	}

	pemKeyBuf := bytes.NewBuffer(pemKeyBlock.Bytes)

	return pemKeyBuf.Bytes(), nil
}
