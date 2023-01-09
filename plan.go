package culqi

import (
	"bytes"
	"net/url"
)

const (
	planURL = baseURL + "/plans"
)

// Create método para crear un plan
func CreatePlan(body []byte) ([]byte, error) {
	res, err := do("POST", planURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID método para obtener un plan por id
func GetByIDPlan(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", planURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAll método para obtener la lista de los planes
func GetAllPlan(queryParams url.Values) ([]byte, error) {
	res, err := do("GET", planURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un plan
func UpdatePlan(id string, body []byte) ([]byte, error) {
	res, err := do("PATCH", planURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete método para eliminar un plan por id
func DeletePlan(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", planURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
