package opsec

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

// Encrypts data with SecretBox Encryption. Good for small messages.
func SBoxEnc(Key []byte, content []byte) ([]byte, error) {
	if Key == nil {
		Key = *SboxKey
	}

	var secretKey [32]byte
	copy(secretKey[:], Key)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}

	// This encrypts content and appends to nonce
	encrypted := secretbox.Seal(nonce[:], content, &nonce, &secretKey)

	return encrypted, nil
}

// Decrypts data with SecretBox Encryption. Good for small messages.
func SBoxDec(Key []byte, content []byte) ([]byte, error) {

	if Key == nil {
		Key = *SboxKey
	}

	var secretKey [32]byte
	copy(secretKey[:], Key)

	var decryptNonce [24]byte
	copy(decryptNonce[:], content[:24])
	decrypted, ok := secretbox.Open(nil, content[24:], &decryptNonce, &secretKey)
	if !ok {
		return nil, fmt.Errorf("SBOX: Error decrypting.")
	}

	return decrypted, nil

}
