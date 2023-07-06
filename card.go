package culqi

import (
	"net/url"
)

const (
	cardURL = baseURL + "/cards"
)

// Create método para crear una tarjeta
func CreateCard(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(cardURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una tarjeta por id
func GetByIDCard(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(cardURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de las tarjetas
func GetAllCard(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(cardURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una tarjeta
func UpdateCard(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(cardURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar una tarjeta por id
func DeleteCard(id string, body []byte) (int, string, error) {
	statusCode, res, err := Delete(cardURL, id, body)
	return statusCode, res, err
}
