package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	cardURL = baseURL + "/cards"
)

// ResponseCard objeto respuesta de tarjeta
type ResponseCard struct {
	Object       string            `json:"object"`
	ID           string            `json:"id"`
	CreationDate int               `json:"creation_date"`
	CustomerID   string            `json:"customer_id"`
	Source       ResponseToken     `json:"source"`
	Metadata     map[string]string `json:"metadata"`
	ReviewCode   string            `json:"action_code"`
	UserMessage  string            `json:"user_message"`
}

// ResponseCardAll respuesta de tarjeta para GetAll
type ResponseCardAll struct {
	Data []ResponseCard `json:"data"`
	WrapperResponse
}

// Create método para crear una tarjeta
func CreateCard(tk []byte) (*ResponseCard, error) {

	res, err := do("POST", cardURL, nil, bytes.NewBuffer(tk))
	if err != nil {
		return nil, err
	}

	rcrd := &ResponseCard{}
	if err = json.Unmarshal(res, rcrd); err != nil {
		return nil, err
	}

	return rcrd, nil
}

// GetByID método para obtener una tarjeta por id
func GetByIDCard(id string) (*ResponseCard, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", cardURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rcrd := &ResponseCard{}
	if err = json.Unmarshal(res, rcrd); err != nil {
		return nil, err
	}

	return rcrd, nil
}

// GetAll método para obtener la lista de las tarjetas
func GetAllCard(queryParams url.Values) (*ResponseCardAll, error) {
	res, err := do("GET", cardURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rcrd := &ResponseCardAll{}
	if err = json.Unmarshal(res, rcrd); err != nil {
		return nil, err
	}

	return rcrd, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una tarjeta
func UpdateCard(id string, tk []byte) (*ResponseCard, error) {

	res, err := do("PATCH", cardURL+"/"+id, nil, bytes.NewBuffer(tk))
	if err != nil {
		return nil, err
	}

	rcrd := &ResponseCard{}
	if err = json.Unmarshal(res, rcrd); err != nil {
		return nil, err
	}

	return rcrd, nil
}

// Delete método para eliminar una tarjeta por id
func DeleteCard(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", cardURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
