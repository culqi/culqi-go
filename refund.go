package culqi

import (
	"net/url"
)

const (
	refundURL = baseURL + "/refunds"
)

// Create método para crear una devolucion
func CreateRefund(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(refundURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una devolucion por id
func GetByIDRefund(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(refundURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de devoluciones
func GetAllRefund(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(refundURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una devolucion
func UpdateRefund(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(refundURL, id, body, encryptionData...)
	return statusCode, res, err
}
