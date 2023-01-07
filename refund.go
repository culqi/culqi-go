package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	refundURL = baseURL + "/refunds"
)

// Refund litar objeto request devolucion
type RefundList struct {
	CreationDate         int    `json:"creation_date"`
	CreationDateFrom     string `json:"creation_date_from"`
	CreationDateTo       string `json:"creation_date_to"`
	Reason               string `json:"reason"`
	Limit                string `json:"limit"`
	Before               string `json:"before"`
	After                string `json:"after"`
	ModificationDateFrom string `json:"modification_date_from"`
	ModificationDateTo   string `json:"modification_date_to"`
	Status               string `json:"status"`
}

// ResponseRefunf objeto respuesta de devoluciones
type ResponseRefund struct {
	Object       string            `json:"object"`
	ID           string            `json:"id"`
	ChargeID     string            `json:"charge_id"`
	CreationDate int               `json:"creation_date"`
	Amount       int               `json:"amount"`
	Reason       string            `json:"reason"`
	Metadata     map[string]string `json:"metadata"`
	Status       int               `json:"status"`
	LastModified int               `json:"last_modified"`
}

// ResponseRefundAll respuesta de devolucion para GetAll y Update
type ResponseRefundAll struct {
	Data []ResponseRefund `json:"data"`
	WrapperResponse
}

// Create método para crear una devolucion
func CreateRefund(body []byte) (*ResponseRefund, error) {
	res, err := do("POST", chargesURL, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	rc := &ResponseRefund{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// GetByID método para obtener una devolucion por id
func GetByIDRefund(id string) (*ResponseRefund, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", refundURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rc := &ResponseRefund{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// GetAll método para obtener la lista de devoluciones
func GetAllRefund(queryParams url.Values) (*ResponseRefundAll, error) {
	res, err := do("GET", refundURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rcs := &ResponseRefundAll{}
	if err = json.Unmarshal(res, rcs); err != nil {
		return nil, err
	}

	return rcs, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una devolucion
func UpdateRefund(id string, body []byte) (*ResponseRefund, error) {
	res, err := do("PATCH", chargesURL+"/"+id, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	rc := &ResponseRefund{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}
