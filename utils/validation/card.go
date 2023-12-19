package culqi

import (
	"strconv"

	utils "github.com/culqi/culqi-go/utils"
)

// CulqiValidation contains methods for validating card data
type CardValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewCardValidation() *CardValidation {
	return &CardValidation{}
}

func (t *CardValidation) Create(data map[string]interface{}) error {
	if err := ValidateStringStart(data["customer_id"].(string), "cus"); err != nil {
		return err
	}
	if err := ValidateStringStart(data["token_id"].(string), "tkn"); err != nil {
		return err
	}
	return nil
}

func CardListValidation(data map[string]interface{}) error {
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
