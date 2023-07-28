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

	iv := make([]byte, 12) // GCM mode requires a 96-bit (12 bytes) random initialization vector
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	blockAES, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Read the contents of the file into a byte slice
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	// Create the GCM mode block
	aesGCM, err := cipher.NewGCM(blockAES)
	if err != nil {
		panic(err)
	}

	// Seal the plaintext using GCM encryption
	ciphertext := aesGCM.Seal(nil, iv, bodyBytes, nil)

	// Base64 encode the ciphertext
	encryptedData64 := base64.StdEncoding.EncodeToString(ciphertext)

	reRsa := regexp.MustCompile(`"rsa_id":\s*"([^"]+)"`)
	rePk := regexp.MustCompile(`"rsa_public_key":\s*"([^"]+)"`)

	matchesIdRsa := reRsa.FindSubmatch(encryptionData)
	matchesPk := rePk.FindSubmatch(encryptionData)

	rsaID := string(matchesIdRsa[1])
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

	// Encrypt the message with PKCS1_OAEP padding and SHA-256 hash function
	ciphertextKey, _ := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		key,
		nil,
	)

	ciphertextIV, _ := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		iv,
		nil,
	)

	// Encode ciphertext as base64
	keyBase64 := base64.StdEncoding.EncodeToString(ciphertextKey)
	ivBase64 := base64.StdEncoding.EncodeToString(ciphertextIV)

	var data = []byte(`{		
		"encrypted_data": "` + encryptedData64 + `",
		"encrypted_key": "` + keyBase64 + `",
		"encrypted_iv": "` + ivBase64 + `"
	}`)

	return bytes.NewBuffer(data), rsaID
}

// Pads a message to a multiple of the block size using PKCS#7 padding
func pad(message []byte, blockSize int) []byte {
	padding := blockSize - len(message)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(message, padText...)
}
