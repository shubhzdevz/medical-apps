package abe

import (
	"fmt"

	"github.com/cloudflare/circl/group"
	"github.com/cloudflare/circl/sign/edwards25519"
	"github.com/cloudflare/circl/sign/hfs"
)

// SetupParams contains the setup parameters for the CP-ABE scheme
type SetupParams struct {
	Policy string
}

// SecretKey contains the secret key for a specific attribute set
type SecretKey struct {
	MasterKey  *hfs.PrivateKey
	Attributes []string
}

// Encrypt encrypts the given message with the given attributes and returns the ciphertext
func Encrypt(message string, attributes []string) (string, error) {
	params := &SetupParams{
		Policy: fmt.Sprintf("(%s)", joinAttributes(attributes)),
	}

	// Generate the CP-ABE keys for the policy
	masterKey, err := hfs.GenerateKey(edwards25519.NewAES128SHA256Ed25519(true), params.Policy)
	if err != nil {
		return "", err
	}
	publicKey := masterKey.PublicKey().(*hfs.PublicKey)

	// Encrypt the message using the CP-ABE scheme
	ciphertext, err := publicKey.Encrypt([]byte(message))
	if err != nil {
		return "", err
	}

	// Generate the attributes for the secret key
	secretKey := &SecretKey{
		MasterKey:  masterKey,
		Attributes: attributes,
	}

	// Serialize the secret key and return it along with the ciphertext
	serializedSecretKey, err := secretKey.Serialize()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", string(serializedSecretKey), string(ciphertext)), nil
}

// Decrypt decrypts the given ciphertext using the given secret key and returns the plaintext message
func Decrypt(ciphertext string, secretKey *SecretKey) (string, error) {
	// Deserialize the secret key
	deserializedSecretKey, err := DeserializeSecretKey([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// Decrypt the ciphertext using the CP-ABE scheme
	plaintext, err := deserializedSecretKey.MasterKey.Decrypt([]byte(ciphertext))
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// joinAttributes joins the given attributes using the AND operator
func joinAttributes(attributes []string) string {
	if len(attributes) == 0 {
		return ""
	}
	if len(attributes) == 1 {
		return attributes[0]
	}
	result := "("
	for i, attr := range attributes {
		result += attr
		if i != len(attributes)-1 {
			result += " AND "
		}
	}
	result += ")"
	return result
}
