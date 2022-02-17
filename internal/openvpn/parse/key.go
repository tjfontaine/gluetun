package parse

import (
	"fmt"
)

func ExtractPrivateKey(b []byte) (keyData string, err error) {
	keyData, err = extractPEM(b, "PRIVATE KEY")
	if err != nil {
		return "", fmt.Errorf("cannot extract PEM data: %w", err)
	}

	return keyData, nil
}

func ExtractEncryptedPrivateKey(b []byte) (keyData string, err error) {
	keyData, err = extractPEM(b, "ENCRYPTED PRIVATE KEY")
	if err != nil {
		return "", fmt.Errorf("cannot extract PEM data: %w", err)
	}

	return keyData, nil
}
