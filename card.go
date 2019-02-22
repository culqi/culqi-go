package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	cardURL = baseURL + "/cards"
)

// Card objeto request de tarjeta
type Card struct {
	CustomerID string            `json:"customer_id"`
	TokenID    string            `json:"token_id"`
	Validate   bool              `json:"validate"`
	Metadata   map[string]string `json:"metadata"`
}

// ResponseCard objeto respuesta de tarjeta
type ResponseCard struct {
	Object       string            `json:"object"`
	ID           string            `json:"id"`
	CreationDate int               `json:"creation_date"`
	CustomerID   string            `json:"customer_id"`
	Source       ResponseToken     `json:"source"`
	Metadata     map[string]string `json:"metadata"`
}

// ResponseCardAll respuesta de tarjeta para GetAll
type ResponseCardAll struct {
	Data []ResponseCard `json:"data"`
	WrapperResponse
}

// Create método para crear una tarjeta
func (crd *Card) Create() (*ResponseCard, error) {
	j, err := json.Marshal(crd)
	if err != nil {
		return nil, err
	}

	res, err := do("POST", cardURL, nil, bytes.NewBuffer(j))
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
func (crd *Card) GetByID(id string) (*ResponseCard, error) {
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
func (crd *Card) GetAll(queryParams url.Values) (*ResponseCardAll, error) {
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
func (crd *Card) Update(id string, metadata map[string]string) (*ResponseCard, error) {
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

	res, err := do("PATCH", cardURL+"/"+id, nil, bytes.NewBuffer(j))
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
func (crd *Card) Delete(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", cardURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
