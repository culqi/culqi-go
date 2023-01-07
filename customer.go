package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	customerURL = baseURL + "/customers"
)

// ResponseCustomer objeto respuesta de cliente
type ResponseCustomer struct {
	Object           string            `json:"object"`
	ID               string            `json:"id"`
	CreationDate     int               `json:"creation_date"`
	Email            string            `json:"email"`
	AntifraudDetails antifraud         `json:"antifraud_details"`
	Metadata         map[string]string `json:"metadata"`
}

// ResponseCustomerAll respuesta de cliente para GetAll
type ResponseCustomerAll struct {
	Data []ResponseCustomer `json:"data"`
	WrapperResponse
}

// Create método para crear un cliente
func CreateCustomer(body []byte) (*ResponseCustomer, error) {
	res, err := do("POST", customerURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	rcus := &ResponseCustomer{}
	if err = json.Unmarshal(res, rcus); err != nil {
		return nil, err
	}

	return rcus, nil
}

// GetByID método para obtener un cliente por id
func GetByIDCustomer(id string) (*ResponseCustomer, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", customerURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rcus := &ResponseCustomer{}
	if err = json.Unmarshal(res, rcus); err != nil {
		return nil, err
	}

	return rcus, nil
}

// GetAll método para obtener la lista de clientes
func GetAllCustomer(queryParams url.Values) (*ResponseCustomerAll, error) {
	res, err := do("GET", customerURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rcus := &ResponseCustomerAll{}
	if err = json.Unmarshal(res, rcus); err != nil {
		return nil, err
	}

	return rcus, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un cliente
func UpdateCustomer(id string, body []byte) (*ResponseCustomer, error) {
	res, err := do("PATCH", customerURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	rcus := &ResponseCustomer{}
	if err = json.Unmarshal(res, rcus); err != nil {
		return nil, err
	}

	return rcus, nil
}

// Delete método para eliminar un cliente por id
func DeleteCustomer(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", customerURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
