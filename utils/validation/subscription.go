package culqi

import (
	"strconv"
)

// CulqiValidation contains methods for validating card data
type SubscriptionValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewSubscriptionValidation() *SubscriptionValidation {
	return &SubscriptionValidation{}
}

func (t *SubscriptionValidation) Create(data map[string]interface{}) error {
	if err := ValidateStringStart(data["card_id"].(string), "crd"); err != nil {
		return err
	}
	if err := ValidateStringStart(data["plan_id"].(string), "pln"); err != nil {
		return err
	}
	return nil
}

func SubscriptionListValidation(data map[string]interface{}) error {
	if _, exists := data["plan_id"]; exists {
		if err := ValidateStringStart(data["plan_id"].(string), "pln"); err != nil {
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
