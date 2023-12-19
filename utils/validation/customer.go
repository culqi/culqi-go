package culqi

import (
	utils "github.com/culqi/culqi-go/utils/validation"
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
			return utils.NewCustomError("%s is empty", field)
		}
	}

	// Validate country code
	countryCode, ok := data["country_code"].(string)
	if !ok {
		return utils.NewCustomError("country code must be a string")
	}
	if err := ValidateValue(countryCode, utils.GetCountryCodes()); err != nil {
		return err
	}

	// Validate email
	email, ok := data["email"].(string)
	if !ok || !IsValidEmail(email) {
		return utils.NewCustomError("invalid email")
	}

	return nil
}
