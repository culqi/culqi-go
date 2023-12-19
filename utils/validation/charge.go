package culqi

import utils "github.com/culqi/culqi-go/utils/validation"

// CulqiValidation contains methods for validating card data
type ChargeValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewChargeValidation() *ChargeValidation {
	return &ChargeValidation{}
}

func (t *ChargeValidation) Create(data map[string]interface{}) error {
	// Validate email
	if email, ok := data["email"].(string); ok {
		if !IsValidEmail(email) {
			return utils.NewCustomError("invalid email")
		}
	} else {
		return utils.NewCustomError("email must be a string")
	}

	// Validate amount
	if amount, ok := data["amount"].(float64); !ok || int(amount) != int(data["amount"].(float64)) {
		return utils.NewCustomError("invalid amount")
	}

	// Validate currency code
	if currencyCode, ok := data["currency_code"].(string); ok {
		if err := ValidateCurrencyCode(currencyCode); err != nil {
			return err
		}
	} else {
		return utils.NewCustomError("currency code must be a string")
	}

	// Validate source ID
	if sourceID, ok := data["source_id"].(string); ok {
		if err := ValidateStringStart(sourceID, "tkn"); err != nil {
			return err
		}
	} else {
		return utils.NewCustomError("source ID must be a string")
	}

	return nil
}
