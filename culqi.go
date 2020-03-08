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
	PublicKey string
	SecretKey string
}

// Key crea una única instancia de key
func Key(pKey, sKey string) {
	keyInstance.PublicKey = pKey
	keyInstance.SecretKey = sKey
}

// GetKey retorna la instancia de key
func GetKey() *key {
	return keyInstance
}
