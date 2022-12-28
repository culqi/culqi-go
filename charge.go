package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	chargesURL = baseURL + "/charges"
)

// Charge objeto request cargo
type Charge struct {
	Amount            int               `json:"amount"`
	Capture           bool              `json:"capture"`
	CurrencyCode      string            `json:"currency_code"`
	Description       string            `json:"description"`
	Email             string            `json:"email"`
	Installments      int               `json:"installments"`
	Metadata          map[string]string `json:"metadata"`
	SourceID          string            `json:"source_id"`
	Address           string            `json:"address"`
	AddressCity       string            `json:"address_city"`
	CountryCode       string            `json:"country_code"`
	FirstName         string            `json:"first_name"`
	LastName          string            `json:"last_name"`
	PhoneNumber       int               `json:"phone_number"`
	Authentication3DS map[string]string `json:"authentication_3DS"`
}

// ResponseCharge objeto respuesta de cargo
type ResponseCharge struct {
	Duplicated         bool        `json:"duplicated"`
	Object             string      `json:"object"`
	ID                 string      `json:"id"`
	CreationDate       int         `json:"creation_date"`
	Amount             int         `json:"amount"`
	AmountRefunded     int         `json:"amount_refunded"`
	CurrentAmount      int         `json:"current_amount"`
	Installments       int         `json:"installments"`
	InstallmentsAmount int         `json:"installments_amount"`
	CurrencyCode       string      `json:"currency_code"`
	Email              string      `json:"email"`
	Description        string      `json:"description"`
	Source             interface{} `json:"source"`
	Outcome            struct {
		Type            string `json:"type"`
		Code            string `json:"code"`
		DeclineCode     string `json:"decline_code"`
		MerchantMessage string `json:"merchant_message"`
		UserMessage     string `json:"user_message"`
	} `json:"outcome"`
	FraudScore        float64           `json:"fraud_score"`
	AntifraudDetails  antifraud         `json:"antifraud_details"`
	Dispute           bool              `json:"dispute"`
	Capture           bool              `json:"capture"`
	ReferenceCode     string            `json:"reference_code"`
	AuthorizationCode string            `json:"authorization_code"`
	Metadata          map[string]string `json:"metadata"`
	TotalFee          int               `json:"total_fee"`
	FeeDetails        struct {
		FixedFee struct {
			Amount                   int    `json:"amount"`
			CurrencyCode             string `json:"currency_code"`
			ExchangeRate             string `json:"exchange_rate"`
			ExchangeRateCurrencyCode string `json:"exchange_rate_currency_code"`
			Total                    int    `json:"total"`
		} `json:"fixed_fee"`
		VariableFee struct {
			CurrencyCode string  `json:"currency_code"`
			Commision    float64 `json:"commision"`
			Total        int     `json:"total"`
		} `json:"variable_fee"`
	} `json:"fee_details"`
	TotalFeeTaxes       int    `json:"total_fee_taxes"`
	TransferAmount      int    `json:"transfer_amount"`
	Paid                bool   `json:"paid"`
	StatementDescriptor string `json:"statement_descriptor"`
	TransferID          string `json:"transfer_id"`
	ReviewCode          string `json:"action_code"`
}

type antifraud struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	AddressCity string `json:"address_city"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
}

// ResponseChargeAll respuesta de cargo para GetAll y Update
type ResponseChargeAll struct {
	Data []ResponseCharge `json:"data"`
	WrapperResponse
}

// Create método para crear un cargo
func (ch *Charge) Create() (*ResponseCharge, error) {
	j, err := json.Marshal(ch)
	if err != nil {
		return nil, err
	}

	res, err := do("POST", chargesURL, nil, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	rc := &ResponseCharge{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// GetByID método para obtener un cargo por id
func (ch *Charge) GetByID(id string) (*ResponseCharge, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", chargesURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rc := &ResponseCharge{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}

// GetAll método para obtener la lista de Cargos
func (ch *Charge) GetAll(queryParams url.Values) (*ResponseChargeAll, error) {
	res, err := do("GET", chargesURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rcs := &ResponseChargeAll{}
	if err = json.Unmarshal(res, rcs); err != nil {
		return nil, err
	}

	return rcs, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un cargo
func (ch *Charge) Update(id string, metadata map[string]string) (*ResponseCharge, error) {
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

	res, err := do("PATCH", chargesURL+"/"+id, nil, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	rc := &ResponseCharge{}
	if err = json.Unmarshal(res, rc); err != nil {
		return nil, err
	}

	return rc, nil
}
