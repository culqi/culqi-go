package culqi

import "strconv"

// CulqiValidation contains methods for validating card data
type RefundValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewRefundValidation() *RefundValidation {
	return &RefundValidation{}
}

func (t *RefundValidation) Create(data map[string]interface{}) error {
	// Validate charge ID format
	if chargeID, ok := data["charge_id"].(string); ok {
		if err := ValidateStringStart(chargeID, "chr"); err != nil {
			return err
		}
	} else {
		return NewCustomError("charge ID must be a string")
	}

	// Validate reason
	if reason, ok := data["reason"].(string); ok {
		allowedValues := []string{"duplicado", "fraudulento", "solicitud_comprador"}
		if err := ValidateValue(reason, allowedValues); err != nil {
			return err
		}
	} else {
		return NewCustomError("reason must be a string")
	}

	// Validate amount
	if amount, ok := data["amount"].(float64); !ok || int(amount) != int(data["amount"].(float64)) {
		return NewCustomError("invalid amount")
	}

	return nil
}

func RefundListValidation(data map[string]interface{}) error {
	if _, exists := data["reason"]; exists {
		allowedValues := []string{"duplicado", "fraudulento", "solicitud_comprador"}
		err := ValidateValue(data["reason"].(string), allowedValues)
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
