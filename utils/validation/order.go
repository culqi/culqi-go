package culqi

import (
	"strconv"
)

// CulqiValidation contains methods for validating card data
type OrderValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewOrderValidation() *OrderValidation {
	return &OrderValidation{}
}

func (t *OrderValidation) Create(data map[string]interface{}) error {
	amount, ok := data["amount"].(float64)
	if !ok || int(amount) != int(data["amount"].(float64)) {
		return NewCustomError("invalid amount")
	}

	if err := ValidateCurrencyCode(data["currency_code"].(string)); err != nil {
		return err
	}

	clientDetails, ok := data["client_details"].(map[string]interface{})
	if !ok {
		return NewCustomError("client details must be a map")
	}

	requiredClientFields := []string{"first_name", "last_name", "phone_number"}
	for _, field := range requiredClientFields {
		if _, ok := clientDetails[field].(string); !ok || clientDetails[field] == "" {
			return NewCustomError("first_name or last_name or phone_number is empty")
		}
	}

	email, ok := clientDetails["email"].(string)
	if !ok || !IsValidEmail(email) {
		return NewCustomError("invalid email")
	}

	expirationDate, ok := data["expiration_date"].(string)

	if !ok {
		return NewCustomError(expirationDate)
	}

	isFuture := IsFutureDate(expirationDate)
	if !isFuture {
		return NewCustomError("expiration_date must be a future date")
	}

	return nil
}

func OrderListValidation(data map[string]interface{}) error {
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
