package culqi

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	//
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"regexp"
)

const (
	apiVersion   = "v2.0"
	baseURL      = "https://api.culqi.com/v2"
	baseURLToken = "https://secure.culqi.com/v2"
)

// Errors API
var (
	ErrInvalidRequest = errors.New("La petición tiene una sintaxis inválida")
	ErrAuthentication = errors.New("La petición no pudo ser procesada debido a problemas con las llaves")
	ErrParameter      = errors.New("Algún parámetro de la petición es inválido")
	ErrCard           = errors.New("No se pudo realizar el cargo a una tarjeta")
	ErrLimitAPI       = errors.New("Estás haciendo muchas peticiones rápidamente al API o superaste tu límite designado")
	ErrResource       = errors.New("El recurso no puede ser encontrado, es inválido o tiene un estado diferente al permitido")
	ErrAPI            = errors.New("Error interno del servidor de Culqi")
	ErrUnexpected     = errors.New("Error inesperado, el código de respuesta no se encuentra controlado")
)

// WrapperResponse respuesta generica para respuestas GetAll
type WrapperResponse struct {
	Paging struct {
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Cursors  struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

func do(method, endpoint string, params url.Values, body io.Reader, encryptionData ...byte) ([]byte, error) {
	idRsaHeader := ""
	if encryptionData != nil {
		body, idRsaHeader = encrypt(body, encryptionData)
	}
	if len(params) != 0 {
		endpoint += "?" + params.Encode()
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(buf.Bytes()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+keyInstance.Key)
	if idRsaHeader != "" {
		req.Header.Set("x-culqi-rsa-id", idRsaHeader)
	}

	c := &http.Client{}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	obj, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(obj)

	switch res.StatusCode {
	case 400:
		err = ErrInvalidRequest
	case 401:
		err = ErrAuthentication
	case 422:
		err = ErrParameter
	case 402:
		err = ErrCard
	case 429:
		err = ErrLimitAPI
	case 404:
		err = ErrResource
	case 500, 503:
		err = ErrAPI
	}

	if err != nil {
		err = fmt.Errorf("%v: %s", err, string(obj))
		return nil, err
	}

	if res.StatusCode >= 200 && res.StatusCode <= 206 {
		return obj, nil
	}

	return nil, ErrUnexpected
}

func encrypt(body io.Reader, encryptionData []byte) (io.Reader, string) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}

	iv := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	block_aes, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Read the contents of the file into a byte slice
	body_byte, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	paddedMessage := pad([]byte(body_byte), block_aes.BlockSize())

	mode := cipher.NewCBCEncrypter(block_aes, iv)

	ciphertext := make([]byte, len(paddedMessage))
	mode.CryptBlocks(ciphertext, paddedMessage)
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
	ciphertext_key, _ := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		key,
		nil,
	)

	ciphertext_iv, _ := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		iv,
		nil,
	)

	// Encode ciphertext as base64
	keyBase64 := base64.StdEncoding.EncodeToString(ciphertext_key)
	ivBase64 := base64.StdEncoding.EncodeToString(ciphertext_iv)

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
