package culqi

import (
	"encoding/json"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	customerURL = baseURL + "/customers"
)

// Create método para crear un cliente
func CreateCustomer(body []byte, encryptionData ...byte) (int, string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return 0, "", err
	}
	// Perform validation
	validator := utils.NewCustomerValidation()
	err = validator.Create(data)
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Create(customerURL, body, encryptionData...)
	return statusCode, res, err
}

// GetByID método para obtener un cliente por id
func GetByIDCustomer(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "cus")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := GetById(customerURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de clientes
func GetAllCustomer(queryParams url.Values, body []byte) (int, string, error) {
	data := make(map[string]interface{})
	for key, values := range queryParams {
		if len(values) > 0 {
			data[key] = values[0] // Taking only the first value for each key.
		}
	}

	err := utils.CustomerListValidation(data)
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := GetAll(customerURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un cliente
func UpdateCustomer(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "cus")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Update(customerURL, id, body, encryptionData...)
	return statusCode, res, err

}

// Delete método para eliminar un cliente por id
func DeleteCustomer(id string, body []byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "cus")
	if err != nil {
		return 0, "", err
	}
	statusCode, res, err := Delete(customerURL, id, body)
	return statusCode, res, err
}
