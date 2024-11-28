package culqi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	subscriptionURL = baseURL + "/recurrent/subscriptions"
)

// Create método para crear una Subscripción
func CreateSubscription(body []byte, encryptionData ...byte) (int, string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return 0, "", err
	}
	// Perform validation
	validator := utils.NewSubscriptionValidation()
	err = validator.Create(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := Create(subscriptionURL+"/create", body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una Subscripción por id
func GetByIDSubscription(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "sxn")
	if err != nil {
		return ErrorBadRequest, "", err
	}
	validator := utils.NewSubscriptionValidation()
	if err := validator.ValidateId(id); err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := GetById(subscriptionURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de las subscripciones
func GetAllSubscription(queryParams url.Values, body []byte) (int, string, error) {
	data := make(map[string]interface{})
	for key, values := range queryParams {
		if len(values) > 0 {
			data[key] = values[0] // Taking only the first value for each key.
		}
	}

	err := utils.SubscriptionListValidation(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := GetAll(subscriptionURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una Subscripción
func UpdateSubscription(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "sxn")
	if err != nil {
		return ErrorBadRequest, "", err
	}

	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		return 0, "", fmt.Errorf("error al decodificar JSON: %w", err)
	}

	if data == nil {
		return 0, "", errors.New("el cuerpo JSON no se deserializó correctamente")
	}

	validator := utils.NewSubscriptionValidation()
	if err := validator.Update(data, id); err != nil {
		return ErrorBadRequest, "", err
	}

	statusCode, res, err := Update(subscriptionURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar una Subscripción por id
func DeleteSubscriptions(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "sxn")
	if err != nil {
		return ErrorBadRequest, "", err
	}
	validator := utils.NewSubscriptionValidation()
	if err := validator.ValidateId(id); err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := Delete(subscriptionURL, id, body)
	return statusCode, res, err
}
