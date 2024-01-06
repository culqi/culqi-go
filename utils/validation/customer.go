package culqi

import (
	utils "github.com/culqi/culqi-go/utils"
)

// CulqiValidation contains methods for validating card data
type CustomerValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewCustomerValidation() *CustomerValidation {
	return &CustomerValidation{}
}

func (t *CustomerValidation) Create(data map[string]interface{}) error {
	requiredFields := []string{"first_name", "last_name", "address", "address_city"}
	for _, field := range requiredFields {
		if _, ok := data[field].(string); !ok || data[field] == "" {
			return NewCustomError("field is empty")
		}
	}

	// Validate country code
	countryCode, ok := data["country_code"].(string)
	if !ok {
		return NewCustomError("country code must be a string")
	}
	if err := ValidateValue(countryCode, utils.GetCountryCodes()); err != nil {
		return err
	}

	// Validate email
	email, ok := data["email"].(string)
	if !ok || !IsValidEmail(email) {
		return NewCustomError("invalid email")
	}

	return nil
}

func CustomerListValidation(data map[string]interface{}) error {
	if _, exists := data["email"]; exists {
		if email, ok := data["email"].(string); ok {
			if !IsValidEmail(email) {
				return NewCustomError("invalid email")
			}
		}
	}
	if _, exists := data["country_code"]; exists {
		allowedDeviceValues := utils.GetCountryCodes()
		err := ValidateValue(data["country_code"].(string), allowedDeviceValues)
		if err != nil {
			return err
		}
	}

	return nil
}
