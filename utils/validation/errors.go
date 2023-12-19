package culqi

import (
	"encoding/json"
)

type CustomError struct {
	Object          string `json:"object"`
	Type            string `json:"type"`
	MerchantMessage string `json:"merchant_message"`
	UserMessage     string `json:"user_message"`
}

func (e *CustomError) Error() string {
	errorJSON, _ := json.Marshal(e)
	return string(errorJSON)
}

func NewCustomError(merchantMessage string) *CustomError {
	return &CustomError{
		Object:          "error",
		Type:            "param_error",
		MerchantMessage: merchantMessage,
		UserMessage:     merchantMessage,
	}
}
