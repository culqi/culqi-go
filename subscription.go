package culqi

import (
	"encoding/json"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	subscriptionURL = baseURL + "/subscriptions"
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
		return 0, "", err
	}
	statusCode, res, err := Create(subscriptionURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una Subscripción por id
func GetByIDSubscription(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "sxn")
	if err != nil {
		return 0, "", err
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
		return 0, "", err
	}
	statusCode, res, err := GetAll(subscriptionURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una Subscripción
func UpdateSubscription(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "sxn")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Update(subscriptionURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar una Subscripción por id
func DeleteSubscriptions(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "sxn")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Delete(subscriptionURL, id, body)
	return statusCode, res, err
}
