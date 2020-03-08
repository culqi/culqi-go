package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	subscriptionURL = baseURL + "/subscriptions"
)

// Subscription objeto request de Subscription
type Subscription struct {
	CardID   string            `json:"card_id"`
	PlanID   string            `json:"plan_id"`
	Metadata map[string]string `json:"metadata"`
}

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
	Card               Card              `json:"card"`
	Metadata           map[string]string `json:"metadata"`
}

// ResponseSubscriptionAll respuesta de subscripción para GetAll
type ResponseSubscriptionAll struct {
	Data []ResponseSubscription `json:"data"`
	WrapperResponse
}

// Create método para crear una Subscripción
func (sub *Subscription) Create() (*ResponseSubscription, error) {
	j, err := json.Marshal(sub)
	if err != nil {
		return nil, err
	}

	res, err := do("POST", subscriptionURL, nil, bytes.NewBuffer(j))
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
func (sub *Subscription) GetByID(id string) (*ResponseSubscription, error) {
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
func (sub *Subscription) GetAll(queryParams url.Values) (*ResponseSubscriptionAll, error) {
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
func (sub *Subscription) Update(id string, metadata map[string]string) (*ResponseSubscription, error) {
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
	res, err := do("PATCH", subscriptionURL+"/"+id, nil, bytes.NewBuffer(j))
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
func (sub *Subscription) Delete(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", subscriptionURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
