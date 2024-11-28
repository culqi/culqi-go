package culqi

import (
	"encoding/json"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	chargesURL = baseURL + "/charges"
)

// Create método para crear un cargo
func CreateCharge(body []byte, encryptionData ...byte) (int, string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return 0, "", err
	}
	// Perform validation
	validator := utils.NewChargeValidation()
	err = validator.Create(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := Create(chargesURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener un cargo por id
func GetByIdCharge(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "chr")
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := GetById(chargesURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de Cargos
func GetAllCharge(queryParams url.Values, body []byte) (int, string, error) {
	data := make(map[string]interface{})
	for key, values := range queryParams {
		if len(values) > 0 {
			data[key] = values[0] // Taking only the first value for each key.
		}
	}

	err := utils.ChargeListValidation(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := GetAll(chargesURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un cargo
func UpdateCharge(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "chr")
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := Update(chargesURL, id, body, encryptionData...)
	return statusCode, res, err
}
