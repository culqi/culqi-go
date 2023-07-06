package culqi

import (
	"net/url"
)

const (
	chargesURL = baseURL + "/charges"
)

// Create método para crear un cargo
func CreateCharge(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(chargesURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener un cargo por id
func GetByIdCharge(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(chargesURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de Cargos
func GetAllCharge(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(chargesURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un cargo
func UpdateCharge(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(chargesURL, id, body, encryptionData...)
	return statusCode, res, err
}
