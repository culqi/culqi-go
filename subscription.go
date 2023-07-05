package culqi

import (
	"bytes"
	"net/url"
)

const (
	subscriptionURL = baseURL + "/subscriptions"
)

// Create método para crear una Subscripción
func CreateSubscription(body []byte, encryptionData ...byte) (string, error) {
	res, err := do("POST", subscriptionURL, nil, bytes.NewBuffer(body), encryptionData...)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetByID método para obtener una Subscripción por id
func GetByIDSubscription(id string, body []byte) (string, error) {
	if id == "" {
		return "", ErrParameter
	}

	res, err := do("GET", subscriptionURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// GetAll método para obtener la lista de las subscripciones
func GetAllSubscription(queryParams url.Values, body []byte) (string, error) {
	res, err := do("GET", subscriptionURL, queryParams, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una Subscripción
func UpdateSubscription(id string, body []byte, encryptionData ...byte) (string, error) {
	res, err := do("PATCH", subscriptionURL+"/"+id, nil, bytes.NewBuffer(body), encryptionData...)
	if err != nil {
		return "", err
	}
	response := string(res[:])

	return response, nil
}

// Delete método para eliminar una Subscripción por id
func DeleteSubscriptions(id string, body []byte) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", subscriptionURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return nil
}
