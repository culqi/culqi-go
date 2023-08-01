package culqi

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io"
	"regexp"
)

func Encrypt(body io.Reader, encryptionData []byte) (io.Reader, string) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	blockAES, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Read the contents of the file into a byte slice
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	// Create a new GCM mode
	aesGCM, err := cipher.NewGCM(blockAES)
	if err != nil {
		panic(err)
	}

	// Generate a random nonce (IV) for GCM
	nonce = make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	// Encrypt the data with GCM
	ciphertext := aesGCM.Seal(nil, nonce, bodyByte, nil)

	encryptedData64 := base64.StdEncoding.EncodeToString(ciphertext)

	reRsa := regexp.MustCompile(`"rsa_id":\s*"([^"]+)"`)
	rePk := regexp.MustCompile(`"rsa_public_key":\s*"([^"]+)"`)

	matchesIDRsa := reRsa.FindSubmatch(encryptionData)
	matchesPk := rePk.FindSubmatch(encryptionData)

	rsaID := string(matchesIDRsa[1])
	publicKeyString := string(matchesPk[1])

	// Decode the public key
	block, _ := pem.Decode([]byte(publicKeyString))
	if block == nil {
		panic("failed to parse PEM block containing public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse public key: " + err.Error())
	}

	// Type assert the public key to RSA public key
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		panic("failed to get RSA public key")
	}

	// Encrypt the key and nonce with RSA public key
	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublicKey, key, nil)
	if err != nil {
		panic("failed to encrypt key with RSA: " + err.Error())
	}

	encryptedNonce, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublicKey, nonce, nil)
	if err != nil {
		panic("failed to encrypt nonce with RSA: " + err.Error())
	}

	// Encode encrypted key and nonce as base64
	encryptedKeyBase64 := base64.StdEncoding.EncodeToString(encryptedKey)
	encryptedNonceBase64 := base64.StdEncoding.EncodeToString(encryptedNonce)

	var data = []byte(`{		
		"encrypted_data": "` + encryptedData64 + `",
		"encrypted_key": "` + encryptedKeyBase64 + `",
		"encrypted_nonce": "` + encryptedNonceBase64 + `"
	}`)

	return bytes.NewBuffer(data), rsaID
}
