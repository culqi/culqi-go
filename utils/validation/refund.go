package culqi

import (
	utils "github.com/culqi/culqi-go/utils/validation"
)

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
		return utils.NewCustomError("charge ID must be a string")
	}

	// Validate reason
	if reason, ok := data["reason"].(string); ok {
		allowedValues := []string{"duplicado", "fraudulento", "solicitud_comprador"}
		if err := ValidateValue(reason, allowedValues); err != nil {
			return err
		}
	} else {
		return utils.NewCustomError("reason must be a string")
	}

	// Validate amount
	if amount, ok := data["amount"].(float64); !ok || int(amount) != int(data["amount"].(float64)) {
		return utils.NewCustomError("invalid amount")
	}

	return nil
}
