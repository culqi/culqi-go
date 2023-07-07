package culqi

import (
	"net/url"
)

const (
	planURL = baseURL + "/plans"
)

// Create método para crear un plan
func CreatePlan(body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Create(planURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener un plan por id
func GetByIDPlan(id string, body []byte) (int, string, error) {
	statusCode, res, err := GetById(planURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de los planes
func GetAllPlan(queryParams url.Values, body []byte) (int, string, error) {
	statusCode, res, err := GetAll(planURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un plan
func UpdatePlan(id string, body []byte, encryptionData ...byte) (int, string, error) {
	statusCode, res, err := Update(planURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar un plan por id
func DeletePlan(id string, body []byte) (int, string, error) {
	statusCode, res, err := Delete(planURL, id, body)
	return statusCode, res, err
}
