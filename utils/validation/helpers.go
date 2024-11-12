package culqi

import (
	"encoding/json"
	"errors"
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

func additionalValidation(data map[string]interface{}, requiredFields []string, message ...string) error {
	for _, field := range requiredFields {
		value, ok := data[field]
		if !ok || value == nil {
			if len(message) > 0 && message[0] != "" {
				return errors.New(fmt.Sprintf("el campo '%s.%s' es requerido.", message[0], field))
			}
			return errors.New(fmt.Sprintf("el campo '%s' es requerido.", field))
		}
	}

	return nil
}

func validateValueInArray(value interface{}, allowedValues []interface{}) bool {
	for _, v := range allowedValues {
		if value == v {
			return true
		}
	}
	return false
}

func validateFloat64InArray(value float64, allowedValues []float64) bool {
	for _, v := range allowedValues {
		if value == v {
			return true
		}
	}
	return false
}

func validateInRange(value float64, minValue float64, maxValue float64) bool {
	if value < minValue || value > maxValue {
		return false
	}
	return true
}

func validateIsInteger(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return false
	default:
		return true
	}
}
func validateEnumCurrency(str string) error {
	allowedValues := []string{"PEN", "USD"}
	for _, v := range allowedValues {
		if v == str {
			return nil
		}
	}

	return NewCustomError("El campo 'currency' es inválido o está vacío, el código de la moneda en tres letras (Formato ISO 4217). Culqi actualmente soporta las siguientes monedas: ['PEN','USD'].")
}

func validateInitialCycles(initialCycles map[string]interface{}, currency string) error {
	hasInitialCharge, okInitialCyclesHasInitialCharge := initialCycles["has_initial_charge"].(bool)
	if !okInitialCyclesHasInitialCharge {
		return NewCustomError("El campo 'initial_cycles.has_initial_charge' es inválido o está vacío. El valor debe ser un booleano (true o false).")
	}

	count, okInitialCyclesCount := initialCycles["count"].(float64)
	typeCount := validateIsInteger(count)
	if !typeCount || !okInitialCyclesCount {
		return NewCustomError("El campo 'initial_cycles.count' es inválido o está vacío, debe tener un valor numérico.")
	}

	payAmount, okInitialCyclesAmount := initialCycles["amount"].(float64)
	typePayAmount := validateIsInteger(payAmount)
	if !typePayAmount || !okInitialCyclesAmount {
		return NewCustomError("El campo 'initial_cycles.amount' es inválido o está vacío, debe tener un valor numérico.")
	}

	valuesIntervalUnitTime := []float64{1, 2, 3, 4, 5, 6}
	interval_unit_time, okInitialCyclesInterval_unit_time := initialCycles["interval_unit_time"].(float64)
	typInterval_unit_time := validateIsInteger(interval_unit_time)
	if !typInterval_unit_time || !okInitialCyclesInterval_unit_time ||
		!validateFloat64InArray(interval_unit_time, valuesIntervalUnitTime) {
		return NewCustomError("El campo 'initial_cycles.interval_unit_time' tiene un valor inválido o está vacío. Estos son los únicos valores permitidos: [1,2,3,4,5,6]")
	}

	if hasInitialCharge {

		if count < 1 || count > 9999 {
			return NewCustomError("El campo 'initial_cycles.count' solo admite valores numéricos en el rango 1 a 9999.")
		}

	} else {
		if count < 0 || count > 9999 {
			return NewCustomError("El campo 'initial_cycles.count' solo admite valores numéricos en el rango 0 a 9999.")
		}

	}

	return nil
}

func validateMetadataSchema(objMetadata map[string]interface{}) error {
	if len(objMetadata) == 0 {
		return nil
	}

	if err := validateKeyAndValueLength(objMetadata); err != nil {
		return err
	}

	if _, err := json.Marshal(objMetadata); err != nil {
		return err
	}

	return nil
}

func validateKeyAndValueLength(objMetadata map[string]interface{}) error {
	for key, value := range objMetadata {
		keyStr := key
		if len(keyStr) < 1 || len(keyStr) > 30 {
			return NewCustomError("El objeto 'metadata' es inválido, límite key (1 - 30), value (1 - 200)")
		}

		valueStr := fmt.Sprintf("%v", value)
		if len(valueStr) < 1 || len(valueStr) > 2000 {
			return NewCustomError("El objeto 'metadata' es inválido, límite key (1 - 30), value (1 - 200)")
		}
	}

	return nil
}
