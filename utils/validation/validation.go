package culqi

import (
	"errors"
	"fmt"
)

// CulqiValidation contains methods for validating card data
type CulqiValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewCulqiValidation() *CulqiValidation {
	return &CulqiValidation{}
}

// cardValidation validates card data
func cardValidation(data map[string]interface{}) error {
	if err := ValidateStringStart(data["customer_id"].(string), "cus"); err != nil {
		return err
	}
	if err := ValidateStringStart(data["token_id"].(string), "tkn"); err != nil {
		return err
	}
	return nil
}

// subscriptionValidation validates subscription data
func subscriptionValidation(data map[string]interface{}) error {
	if err := ValidateStringStart(data["card_id"].(string), "crd"); err != nil {
		return err
	}
	if err := ValidateStringStart(data["plan_id"].(string), "pln"); err != nil {
		return err
	}
	return nil
}

// orderValidation validates order data
func (cv *CulqiValidation) orderValidation(data map[string]interface{}) error {
	amount, ok := data["amount"].(float64)
	if !ok || int(amount) != int(data["amount"].(float64)) {
		return errors.New("invalid amount")
	}

	if err := ValidateCurrencyCode(data["currency_code"].(string)); err != nil {
		return err
	}

	clientDetails, ok := data["client_details"].(map[string]interface{})
	if !ok {
		return errors.New("client details must be a map")
	}

	requiredClientFields := []string{"first_name", "last_name", "phone_number"}
	for _, field := range requiredClientFields {
		if _, ok := clientDetails[field].(string); !ok || clientDetails[field] == "" {
			return fmt.Errorf("%s is empty", field)
		}
	}

	email, ok := clientDetails["email"].(string)
	if !ok || !IsValidEmail(email) {
		return errors.New("invalid email")
	}

	expirationDate, ok := data["expiration_date"].(int64)
	if !ok {
		return errors.New("expiration date must be a string")
	}
	isFuture := IsFutureDate(expirationDate)
	if !isFuture {
		return errors.New("expiration_date must be a future date")
	}

	return nil
}
