package culqi

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	tokensURL = baseURL + "/tokens"
)

// Token objeto request
type Token struct {
	CardNumber      string            `json:"card_number"`
	Cvv             string            `json:"cvv"`
	ExpirationMonth string            `json:"expiration_month"`
	ExpirationYear  string            `json:"expiration_year"`
	Email           string            `json:"email"`
	Metadata        map[string]string `json:"metadata"`
}

// ResponseToken objeto respuesta de token
type ResponseToken struct {
	Object       string `json:"object"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	CreationDate int    `json:"creation_date"`
	Email        string `json:"email"`
	CardNumber   string `json:"card_number"`
	LastFour     string `json:"last_four"`
	Active       bool   `json:"active"`
	Iin          struct {
		Object              string `json:"object"`
		Bin                 string `json:"bin"`
		CardBranch          string `json:"card_branch"`
		CardType            string `json:"card_type"`
		CardCategory        string `json:"card_category"`
		InstallmentsAllowed []int  `json:"installments_allowed"`
		Issuer              struct {
			Name        string `json:"name"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			Website     string `json:"website"`
			PhoneNumber string `json:"phone_number"`
		} `json:"issuer"`
	} `json:"iin"`
	Client struct {
		IP                string `json:"ip"`
		IPCountry         string `json:"ip_country"`
		IPCountryCode     string `json:"ip_country_code"`
		Browser           string `json:"browser"`
		DeviceFingerprint string `json:"device_fingerprint"`
		DeviceType        string `json:"device_type"`
	} `json:"client"`
	Metadata map[string]string `json:"metadata"`
}

// ResponseTokenAll respuesta de token para GetAll y Update
type ResponseTokenAll struct {
	Data []ResponseToken `json:"data"`
	WrapperResponse
}

// Create método para crear un token
func (tk *Token) Create() (*ResponseToken, error) {
	j, err := json.Marshal(tk)
	if err != nil {
		return nil, err
	}

	res, err := do("POST", tokensURL, nil, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	rt := &ResponseToken{}
	if err = json.Unmarshal(res, rt); err != nil {
		return nil, err
	}

	return rt, nil
}

// GetByID método para obtener un token por id
func (tk *Token) GetByID(id string) (*ResponseToken, error) {
	if id == "" {
		return nil, ErrParameter
	}

	res, err := do("GET", tokensURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	rt := &ResponseToken{}
	if err = json.Unmarshal(res, rt); err != nil {
		return nil, err
	}

	return rt, nil
}

// GetAll método para obtener la lista de tokens
func (tk *Token) GetAll(queryParams url.Values) (*ResponseTokenAll, error) {
	res, err := do("GET", tokensURL, queryParams, nil)
	if err != nil {
		return nil, err
	}

	rts := &ResponseTokenAll{}
	if err = json.Unmarshal(res, rts); err != nil {
		return nil, err
	}

	return rts, nil
}

// Update método para agregar o remplazar información a los valores de la metadata de un token
func (tk *Token) Update(id string, metadata map[string]string) (*ResponseToken, error) {
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

	res, err := do("PATCH", tokensURL+"/"+id, nil, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	rt := &ResponseToken{}
	if err = json.Unmarshal(res, rt); err != nil {
		return nil, err
	}

	return rt, nil
}
