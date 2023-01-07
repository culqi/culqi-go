package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	ordersURL = baseURL + "/orders"
)

// Order objeto request orden por tipo
type OrderTipo struct {
	OrderId    string    `json:"order_id"`
	OrderTypes [2]string `json:"order_types"`
}

// ResponseOrder objeto respuesta de orden
type ResponseOrder struct {
	Object         string            `json:"object"`
	ID             string            `json:"id"`
	Amount         int               `json:"amount"`
	PaymentCode    string            `json:"payment_code"`
	CurrencyCode   string            `json:"currency_code"`
	Description    string            `json:"description"`
	OrderNumber    string            `json:"order_number"`
	State          string            `json:"state"`
	TotalFee       int               `json:"total_fee"`
	NetAmount      int               `json:"net_amount"`
	FeeDetails     int               `json:"fee_details"`
	CreationDate   int               `json:"creation_date"`
	ExpirationDate int               `json:"expiration_date"`
	UpdateAt       int               `json:"update_at"`
	PaidAt         int               `json:"pait_at"`
	AvailableOn    int               `json:"available_on"`
	Metadata       map[string]string `json:"metadata"`
}

// ResponseOrderAll respuesta de orden para GetAll y Update
type ResponseOrderAll struct {
	Data []ResponseOrder `json:"data"`
	WrapperResponse
}

// Create método para crear una orden
func CreateOrder(tk []byte) (*ResponseOrder, error) {
	res, err := do("POST", ordersURL, nil, bytes.NewBuffer(tk))
	if err != nil {
		return nil, err
	}

	rc := &ResponseOrder{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// GetByID método para obtener una orden por id
func GetByIDOrder(id string) (*ResponseOrder, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", ordersURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rc := &ResponseOrder{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// GetAll método para obtener la lista de Ordenes
func GetAllOrder(queryParams url.Values) (*ResponseOrderAll, error) {
	res, err := do("GET", ordersURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rcs := &ResponseOrderAll{}
	if err = json.Unmarshal(res, rcs); err != nil {
		return nil, err
	}

	return rcs, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una orden
func UpdateOrder(id string, tk []byte) (*ResponseOrder, error) {

	res, err := do("PATCH", ordersURL+"/"+id, nil, bytes.NewBuffer(tk))
	if err != nil {
		return nil, err
	}

	rc := &ResponseOrder{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// Delete método para eliminar una orden
func DeleteOrder(id string) (*ResponseOrder, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("DELETE", ordersURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rc := &ResponseOrder{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// Confirm método para confirmar una orden
func ConfirmOrder(id string) (*ResponseOrder, error) {
	res, err := do("POST", ordersURL+"/"+id+"/confirm", nil, nil)
	if err != nil {
		return nil, err
	}

	rc := &ResponseOrder{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// Confirm método para confirmar una orden por tipo
func ConfirmTipoOrder(tk []byte) (*ResponseOrder, error) {
	res, err := do("POST", ordersURL+"/confirm", nil, bytes.NewBuffer(tk))
	if err != nil {
		return nil, err
	}

	rc := &ResponseOrder{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}
