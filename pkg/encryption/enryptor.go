package encryption

import (
	"errors"
	"fmt"
	"time"

	"github.com/getsops/sops/v3"
	"github.com/getsops/sops/v3/aes"
	"github.com/getsops/sops/v3/age"
	"github.com/getsops/sops/v3/keys"
	"github.com/getsops/sops/v3/keyservice"
	"github.com/getsops/sops/v3/stores/yaml"
)

const SopsMetadataKey = "sops"

type Encryptor interface {
	Encrypt(file []byte, commaSeparatedAgeRecipients string) (encrypted []byte, _ error)
}

func NewEncryptor(sopsVersion string) Encryptor {
	return &encryptor{
		sopsVersion: sopsVersion,
	}
}

type encryptor struct {
	sopsVersion string
}

func (e *encryptor) Encrypt(file []byte, commaSeparatedAgeRecipients string) ([]byte, error) {
	ageKeys, err := age.MasterKeysFromRecipients(commaSeparatedAgeRecipients)
	if err != nil {
		return nil, fmt.Errorf("failed to parse to parse age recipients: %v", err)
	}

	sopsKeys := make([]keys.MasterKey, 0)
	for _, key := range ageKeys {
		sopsKeys = append(sopsKeys, key)
	}

	branches, err := (&yaml.Store{}).LoadPlainFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}
	if len(branches) < 1 {
		return nil, errors.New("file cannot be completely empty, it must contain at least one document")
	}

	if hasSopsTopLevelKey(branches[0]) {
		return nil, errors.New("file already encrypted")
	}

	tree := sops.Tree{
		Branches: branches,
		Metadata: sops.Metadata{
			KeyGroups:         []sops.KeyGroup{sopsKeys},
			UnencryptedSuffix: "_unencrypted",
			Version:           e.sopsVersion,
		},
	}

	dataKey, errs := tree.GenerateDataKeyWithKeyServices(
		[]keyservice.KeyServiceClient{keyservice.NewLocalClient()},
	)
	if len(errs) > 0 {
		err = fmt.Errorf("failed to generate data key: %s", errs)
		return nil, err
	}

	cipher := aes.NewCipher()

	mac, err := tree.Encrypt(dataKey, cipher)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt tree: %v", err)
	}

	if err := addMetadata(&tree, mac, cipher, dataKey); err != nil {
		return nil, fmt.Errorf("failed to add metadata to tree: %w", err)
	}

	encryptedFile, err := (&yaml.Store{}).EmitEncryptedFile(tree)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tree: %w", err)
	}

	return encryptedFile, nil
}

func addMetadata(tree *sops.Tree, mac string, cipher sops.Cipher, dataKey []byte) error {
	modifyTime := time.Now().UTC()

	encryptedMac, err := cipher.Encrypt(mac, dataKey, modifyTime.Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("failed to encrypt MAC: %v", err)
	}

	tree.Metadata.LastModified = time.Now().UTC()
	tree.Metadata.MessageAuthenticationCode = encryptedMac

	return nil
}

func hasSopsTopLevelKey(branch sops.TreeBranch) bool {
	for _, b := range branch {
		if b.Key == SopsMetadataKey {
			return true
		}
	}
	return false
}
