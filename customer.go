package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	customerURL = baseURL + "/customers"
)

// Customer objeto request de cliente
type Customer struct {
	Address     string            `json:"address"`
	AddressCity string            `json:"address_city"`
	CountryCode string            `json:"country_code"`
	Email       string            `json:"email"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	PhoneNumber string            `json:"phone_number"`
	Metadata    map[string]string `json:"metadata"`
}

// ResponseCustomer objeto respuesta de cliente
type ResponseCustomer struct {
	Object           string            `json:"object"`
	ID               string            `json:"id"`
	CreationDate     int               `json:"creation_date"`
	Email            string            `json:"email"`
	AntifraudDetails antifraud         `json:"antifraud_details"`
	Cards            []Card            `json:"cards"`
	Metadata         map[string]string `json:"metadata"`
}

// ResponseCustomerAll respuesta de cliente para GetAll
type ResponseCustomerAll struct {
	Data []ResponseCustomer `json:"data"`
	WrapperResponse
}

// Create método para crear un cliente
func (cus *Customer) Create() (*ResponseCustomer, error) {
	j, err := json.Marshal(cus)
	if err != nil {
		return nil, err
	}

	res, err := do("POST", customerURL, nil, bytes.NewBuffer(j))
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
func (cus *Customer) GetByID(id string) (*ResponseCustomer, error) {
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
func (cus *Customer) GetAll(queryParams url.Values) (*ResponseCustomerAll, error) {
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
func (cus *Customer) Update(id string, metadata map[string]string) (*ResponseCustomer, error) {
	if id == "" || len(metadata) == 0 {
		return nil, ErrParameter
	}

	j, err := json.Marshal(
		struct {
			Metadata map[string]string `json:"metadata"`
		}{
			metadata,
		},
	)
	if err != nil {
		return nil, err
	}
	res, err := do("PATCH", customerURL+"/"+id, nil, bytes.NewBuffer(j))
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
func (cus *Customer) Delete(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", customerURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
