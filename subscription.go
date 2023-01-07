package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	subscriptionURL = baseURL + "/subscriptions"
)

// ResponseSubscription objeto respuesta de Subscription
type ResponseSubscription struct {
	Object             string            `json:"object"`
	ID                 string            `json:"id"`
	CreationDate       int               `json:"creation_date"`
	Status             string            `json:"status"`
	CurrentPeriod      int               `json:"current_period"`
	TotalPeriods       int               `json:"total_periods"`
	CurrentPeriodStart int               `json:"current_period_start"`
	CurrentPeriodEnd   int               `json:"current_period_end"`
	CancelAtPeriodEnd  bool              `json:"cancel_at_period_end"`
	CanceledAt         int               `json:"canceled_at"`
	EndedAt            int               `json:"ended_at"`
	NextBillingDate    int               `json:"next_billing_date"`
	TrialStart         int               `json:"trial_start"`
	TrialEnd           int               `json:"trial_end"`
	Charges            []ResponseCharge  `json:"charges"`
	Plan               Plan              `json:"plan"`
	Metadata           map[string]string `json:"metadata"`
}

// ResponseSubscriptionAll respuesta de subscripción para GetAll
type ResponseSubscriptionAll struct {
	Data []ResponseSubscription `json:"data"`
	WrapperResponse
}

// Create método para crear una Subscripción
func CreateSubscription(tk []byte) (*ResponseSubscription, error) {
	res, err := do("POST", subscriptionURL, nil, bytes.NewBuffer(tk))
	if err != nil {
		return nil, err
	}

	rsub := &ResponseSubscription{}
	if err = json.Unmarshal(res, rsub); err != nil {
		return nil, err
	}

	return rsub, nil
}

// GetByID método para obtener una Subscripción por id
func GetByIDSubscription(id string) (*ResponseSubscription, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", subscriptionURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rsub := &ResponseSubscription{}
	if err = json.Unmarshal(res, rsub); err != nil {
		return nil, err
	}

	return rsub, nil
}

// GetAll método para obtener la lista de las subscripciones
func GetAllSubscription(queryParams url.Values) (*ResponseSubscriptionAll, error) {
	res, err := do("GET", subscriptionURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rsub := &ResponseSubscriptionAll{}
	if err = json.Unmarshal(res, rsub); err != nil {
		return nil, err
	}

	return rsub, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de una Subscripción
func UpdateSubscription(id string, tk []byte) (*ResponseSubscription, error) {
	res, err := do("PATCH", subscriptionURL+"/"+id, nil, bytes.NewBuffer(tk))
	if err != nil {
		return nil, err
	}

	rsub := &ResponseSubscription{}
	if err = json.Unmarshal(res, rsub); err != nil {
		return nil, err
	}

	return rsub, nil
}

// Delete método para eliminar una Subscripción por id
func DeleteSubscriptions(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", subscriptionURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
