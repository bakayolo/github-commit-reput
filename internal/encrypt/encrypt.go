package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"github.com/rs/zerolog/log"
)

var publicKey rsa.PublicKey

const (
	keyLength = 2048
)

func GenerateKey() {
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		log.Panic().Err(err).Msgf("Error generating private key")
	}
	publicKey = privateKey.PublicKey
}

func Encrypt(message string) (string, error) {
	// Check rsa.EncryptOAEP() definition
	// The message must be no longer than the length of the public modulus minus
	// twice the hash length, minus a further 2.
	hash := sha256.New()
	pubSize := publicKey.Size()
	messageResized := make([]byte, pubSize-2*hash.Size()-2) // exact check made in rsa.EncryptOAEP function
	copy(messageResized[:], message)

	cipherText, err := rsa.EncryptOAEP(hash, rand.Reader, &publicKey, messageResized, []byte(""))
	if err != nil {
		log.Error().Err(err).Msgf("Error encrypting the message %v", message)
		return "", err
	}
	return string(cipherText), nil
}
