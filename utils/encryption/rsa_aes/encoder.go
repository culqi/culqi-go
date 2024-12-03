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
	"strings"
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

	bodyByte, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	// Crea un nuevo GCM
	aesGCM, err := cipher.NewGCM(blockAES)
	if err != nil {
		panic(err)
	}

	// Cifra los datos con GCM
	ciphertext := aesGCM.Seal(nil, nonce, bodyByte, nil)

	//encryptedData64 := base64.StdEncoding.EncodeToString(ciphertext)

	// Calcula la longitud de los datos cifrados
	cipherTextLen := len(ciphertext)

	// Saca los últimos 16 bytes
	ciphertextWithoutTag := ciphertext[:cipherTextLen-16]

	// Ahora puedes convertir ciphertextWithoutTag a base64
	encryptedData64 := base64.StdEncoding.EncodeToString(ciphertextWithoutTag)

	reRsa := regexp.MustCompile(`"rsa_id":\s*"([^"]+)"`)
	rePk := regexp.MustCompile(`"rsa_public_key":\s*"([^"]+)"`)

	matchesIdRsa := reRsa.FindSubmatch(encryptionData)
	matchesPk := rePk.FindSubmatch(encryptionData)

	rsaID := string(matchesIdRsa[1])
	publicKeyString := string(matchesPk[1])

	// Replace \n with newline character
	publicKeyString = strings.ReplaceAll(publicKeyString, "\\n", "\n")

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
	encryptedKey, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		key,
		nil,
	)

	encryptedNonce, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		nonce,
		nil,
	)

	// Encode ciphertext as base64
	keyBase64 := base64.StdEncoding.EncodeToString(encryptedKey)
	nonceBase64 := base64.StdEncoding.EncodeToString(encryptedNonce)

	var data = []byte(`{		
		"encrypted_data": "` + encryptedData64 + `",
		"encrypted_key": "` + keyBase64 + `",
		"encrypted_iv": "` + nonceBase64 + `"
	}`)

	return bytes.NewBuffer(data), rsaID
}
