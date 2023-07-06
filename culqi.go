package culqi

import "sync"

var (
	once        sync.Once
	keyInstance *key
)

func init() {
	once.Do(func() {
		keyInstance = &key{}
	})
}

type key struct {
	publicKey string
	secretKey string
}

// Key crea una única instancia de key
func Key(publicKey string, secretKey string) {
	keyInstance.publicKey = publicKey
	keyInstance.secretKey = secretKey
}

/*
func Key(key string) {
	keyInstance.Key = key
}
*/
// GetKey retorna la instancia de key
func GetKey() *key {
	return keyInstance
}
