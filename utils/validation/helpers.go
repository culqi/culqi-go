package culqi

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	utils "github.com/culqi/culqi-go/utils/validation"
)

func IsValidCardNumber(number string) bool {
	match, _ := regexp.MatchString(`^\d{13,19}$`, number)
	return match
}

func IsValidEmail(email string) bool {
	match, _ := regexp.MatchString(`^\S+@\S+\.\S+$`, email)
	return match
}

func ValidateCurrencyCode(currencyCode string) error {
	if currencyCode == "" {
		return utils.NewCustomError("Currency code is empty.")
	}

	allowedValues := []string{"PEN", "USD"}
	for _, v := range allowedValues {
		if currencyCode == v {
			return nil
		}
	}
	return utils.NewCustomError("Currency code must be either 'PEN' or 'USD'.")
}

func ValidateStringStart(str string, start string) error {
	if !strings.HasPrefix(str, start+"_test_") && !strings.HasPrefix(str, start+"_live_") {
		return utils.NewCustomError(fmt.Sprintf("Incorrect format. The format must start with %s_test_ or %s_live_", start, start))
	}
	return nil
}

func ValidateValue(value string, allowedValues []string) error {
	for _, v := range allowedValues {
		if value == v {
			return nil
		}
	}
	allowedValuesJSON, _ := json.Marshal(allowedValues)
	return utils.NewCustomError(fmt.Sprintf("Invalid value. It must be %s.", string(allowedValuesJSON)))
}

func IsFutureDate(expirationDate int64) bool {
	expTime := time.Unix(expirationDate, 0)
	return expTime.After(time.Now())
}

func ValidateDateFilter(dateFrom int64, dateTo int64) error {
	if dateTo < dateFrom {
		return utils.NewCustomError("Invalid value. Date_from must be less than date_to")
	}
	return nil
}
