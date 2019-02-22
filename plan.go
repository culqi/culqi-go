package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	planURL = baseURL + "/plans"
)

// Plan objeto request de plan
type Plan struct {
	Name          string            `json:"name"`
	Amount        int               `json:"amount"`
	CurrencyCode  string            `json:"currency_code"`
	Interval      string            `json:"interval"`
	IntervalCount int               `json:"interval_count"`
	TrialDays     int               `json:"trial_days"`
	Limit         int               `json:"limit"`
	Metadata      map[string]string `json:"metadata"`
}

// ResponsePlan objeto respuesta de plan
type ResponsePlan struct {
	Plan
	Object             string `json:"object"`
	ID                 string `json:"id"`
	CreationDate       int    `json:"creation_date"`
	TotalSubscriptions int    `json:"total_subscriptions"`
}

// ResponsePlanAll respuesta de tarjeta para GetAll
type ResponsePlanAll struct {
	Data []ResponsePlan `json:"data"`
	WrapperResponse
}

// Create método para crear un plan
func (pln *Plan) Create() (*ResponsePlan, error) {
	j, err := json.Marshal(pln)
	if err != nil {
		return nil, err
	}

	res, err := do("POST", planURL, nil, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	rpln := &ResponsePlan{}
	if err = json.Unmarshal(res, rpln); err != nil {
		return nil, err
	}

	return rpln, nil
}

// GetByID método para obtener un plan por id
func (pln *Plan) GetByID(id string) (*ResponsePlan, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", planURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rpln := &ResponsePlan{}
	if err = json.Unmarshal(res, rpln); err != nil {
		return nil, err
	}

	return rpln, nil
}

// GetAll método para obtener la lista de los planes
func (pln *Plan) GetAll(queryParams url.Values) (*ResponsePlanAll, error) {
	res, err := do("GET", planURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rpln := &ResponsePlanAll{}
	if err = json.Unmarshal(res, rpln); err != nil {
		return nil, err
	}

	return rpln, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un plan
func (pln *Plan) Update(id string, metadata map[string]string) (*ResponsePlan, error) {
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

	res, err := do("PATCH", planURL+"/"+id, nil, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	rpln := &ResponsePlan{}
	if err = json.Unmarshal(res, rpln); err != nil {
		return nil, err
	}

	return rpln, nil
}

// Delete método para eliminar un plan por id
func (pln *Plan) Delete(id string) error {
	if id == "" {
		return ErrParameter
	}

	_, err := do("DELETE", planURL+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
