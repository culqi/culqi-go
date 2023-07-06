package culqi

import (
	"net/url"
)

const (
	subscriptionURL = baseURL + "/subscriptions"
)

// Create método para crear una Subscripción
func CreateSubscription(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(subscriptionURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una Subscripción por id
func GetByIDSubscription(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(subscriptionURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de las subscripciones
func GetAllSubscription(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(subscriptionURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una Subscripción
func UpdateSubscription(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(subscriptionURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar una Subscripción por id
func DeleteSubscriptions(id string, body []byte) (int, string, error) {
	statusCode, res, err := Delete(subscriptionURL, id, body)
	return statusCode, res, err
}
