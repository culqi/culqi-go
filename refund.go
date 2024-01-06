package culqi

import (
	"encoding/json"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	refundURL = baseURL + "/refunds"
)

// Create método para crear una devolucion
func CreateRefund(body []byte, encryptionData ...byte) (int, string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return 0, "", err
	}
	// Perform validation
	validator := utils.NewRefundValidation()
	err = validator.Create(data)
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Create(refundURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener una devolucion por id
func GetByIDRefund(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "ref")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := GetById(refundURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de devoluciones
func GetAllRefund(queryParams url.Values, body []byte) (int, string, error) {
	data := make(map[string]interface{})
	for key, values := range queryParams {
		if len(values) > 0 {
			data[key] = values[0] // Taking only the first value for each key.
		}
	}

	err := utils.RefundListValidation(data)
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := GetAll(refundURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de una devolucion
func UpdateRefund(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "ref")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Update(refundURL, id, body, encryptionData...)
	return statusCode, res, err
}
