package culqi

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
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
		return NewCustomError("Currency code is empty.")
	}

	allowedValues := []string{"PEN", "USD"}
	for _, v := range allowedValues {
		if currencyCode == v {
			return nil
		}
	}
	return NewCustomError("Currency code must be either 'PEN' or 'USD'.")
}

func ValidateStringStart(str string, start string) error {
	if !strings.HasPrefix(str, start+"_test_") && !strings.HasPrefix(str, start+"_live_") {
		return NewCustomError(fmt.Sprintf("Incorrect format. The format must start with %s_test_ or %s_live_", start, start))
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
	return NewCustomError(fmt.Sprintf("Invalid value. It must be %s.", string(allowedValuesJSON)))
}

func IsFutureDate(expirationDate string) bool {
	expTimeToConvert, _ := strconv.ParseInt(expirationDate, 10, 64)
	expTime := time.Unix(expTimeToConvert, 0)
	return expTime.After(time.Now())
}

func ValidateDateFilter(dateFrom int64, dateTo int64) error {
	if dateTo < dateFrom {
		return NewCustomError("Invalid value. Date_from must be less than date_to")
	}
	return nil
}
