package culqi

import (
	"encoding/json"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	cardURL = baseURL + "/cards"
)

// Create método para crear una tarjeta
func CreateCard(body []byte, encryptionData ...byte) (int, string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return 0, "", err
	}
	// Perform validation
	validator := utils.NewCardValidation()
	err = validator.Create(data)
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Create(cardURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una tarjeta por id
func GetByIDCard(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "crd")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := GetById(cardURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de las tarjetas
func GetAllCard(queryParams url.Values, body []byte) (int, string, error) {
	data := make(map[string]interface{})
	for key, values := range queryParams {
		if len(values) > 0 {
			data[key] = values[0] // Taking only the first value for each key.
		}
	}

	err := utils.CardListValidation(data)
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := GetAll(cardURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una tarjeta
func UpdateCard(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "crd")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Update(cardURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar una tarjeta por id
func DeleteCard(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "crd")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Delete(cardURL, id, body)
	return statusCode, res, err
}
