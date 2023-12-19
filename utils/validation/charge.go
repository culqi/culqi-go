package culqi

import (
	"strconv"
	"strings"

	utils "github.com/culqi/culqi-go/utils"
)

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
			return NewCustomError("invalid email")
		}
	} else {
		return NewCustomError("email must be a string")
	}

	// Validate amount
	if amount, ok := data["amount"].(float64); !ok || int(amount) != int(data["amount"].(float64)) {
		return NewCustomError("invalid amount")
	}

	// Validate currency code
	if currencyCode, ok := data["currency_code"].(string); ok {
		if err := ValidateCurrencyCode(currencyCode); err != nil {
			return err
		}
	} else {
		return NewCustomError("currency code must be a string")
	}

	// Validate source ID
	if sourceID, ok := data["source_id"].(string); ok {
		if strings.HasPrefix(data["source_id"].(string), "tkn") {
			if err := ValidateStringStart(sourceID, "tkn"); err != nil {
				return err
			}
		} else if strings.HasPrefix(data["source_id"].(string), "ype") {
			if err := ValidateStringStart(sourceID, "ype"); err != nil {
				return err
			}
		} else if strings.HasPrefix(data["source_id"].(string), "crd") {
			if err := ValidateStringStart(sourceID, "crd"); err != nil {
				return err
			}
		} else {
			return NewCustomError("Incorrect format. The format must start with tkn, ype or crd")
		}
	} else {
		return NewCustomError("source ID must be a string")
	}

	return nil
}

func ChargeListValidation(data map[string]interface{}) error {
	if _, exists := data["email"]; exists {
		if email, ok := data["email"].(string); ok {
			if !IsValidEmail(email) {
				return NewCustomError("invalid email")
			}
		}
	}
	if _, exists := data["amount"]; exists {
		if amount, ok := data["amount"].(float64); !ok || int(amount) != int(data["amount"].(float64)) {
			return NewCustomError("invalid amount")
		}
	}
	if _, exists := data["min_amount"]; exists {
		if amount, ok := data["min_amount"].(float64); !ok || int(amount) != int(data["min_amount"].(float64)) {
			return NewCustomError("invalid min_amount")
		}
	}
	if _, exists := data["max_amount"]; exists {
		if amount, ok := data["max_amount"].(float64); !ok || int(amount) != int(data["max_amount"].(float64)) {
			return NewCustomError("invalid max_amount")
		}
	}
	if _, exists := data["installments"]; exists {
		if amount, ok := data["installments"].(float64); !ok || int(amount) != int(data["installments"].(float64)) {
			return NewCustomError("invalid installments")
		}
	}
	if _, exists := data["min_installments"]; exists {
		if amount, ok := data["min_installments"].(float64); !ok || int(amount) != int(data["min_installments"].(float64)) {
			return NewCustomError("invalid min_installments")
		}
	}
	if _, exists := data["max_installments"]; exists {
		if amount, ok := data["max_installments"].(float64); !ok || int(amount) != int(data["max_installments"].(float64)) {
			return NewCustomError("invalid max_installments")
		}
	}
	if _, exists := data["card_brand"]; exists {
		allowedDeviceValues := []string{"Visa", "Mastercard", "Amex", "Diners"}
		err := ValidateValue(data["card_brand"].(string), allowedDeviceValues)
		if err != nil {
			return err
		}
	}
	if _, exists := data["card_type"]; exists {
		allowedDeviceValues := []string{"credito", "debito", "internacional"}
		err := ValidateValue(data["card_type"].(string), allowedDeviceValues)
		if err != nil {
			return err
		}
	}
	if _, exists := data["country_code"]; exists {
		allowedDeviceValues := utils.GetCountryCodes()
		err := ValidateValue(data["country_code"].(string), allowedDeviceValues)
		if err != nil {
			return err
		}
	}
	if _, existsFrom := data["creation_date_from"]; existsFrom {
		if _, existsTo := data["creation_date_to"]; existsTo {
			date_from, _ := strconv.ParseInt(data["creation_date_from"].(string), 10, 64)
			date_to, _ := strconv.ParseInt(data["creation_date_to"].(string), 10, 64)
			err := ValidateDateFilter(date_from, date_to)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
