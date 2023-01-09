package culqi

import (
	"bytes"
	"net/url"
)

const (
	subscriptionURL = baseURL + "/subscriptions"
)

// Create método para crear una Subscripción
func CreateSubscription(body []byte) ([]byte, error) {
	res, err := do("POST", subscriptionURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID método para obtener una Subscripción por id
func GetByIDSubscription(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", subscriptionURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAll método para obtener la lista de las subscripciones
func GetAllSubscription(queryParams url.Values) ([]byte, error) {
	res, err := do("GET", subscriptionURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una Subscripción
func UpdateSubscription(id string, body []byte) ([]byte, error) {
	res, err := do("PATCH", subscriptionURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete método para eliminar una Subscripción por id
func DeleteSubscriptions(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", subscriptionURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
