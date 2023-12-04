package culqi

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	utils "github.com/culqi/culqi-go/utils"
)

// CulqiValidation contains methods for validating card data
type CulqiValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewCulqiValidation() *CulqiValidation {
	return &CulqiValidation{}
}

// CreateTokenValidation validates the card data
func (cv *CulqiValidation) CreateTokenValidation(data map[string]string) error {
	// Validate card number
	if !cv.IsValidCardNumber(data["card_number"]) {
		return fmt.Errorf("invalid card number")
	}

	// Validate CVV
	match, _ := regexp.MatchString(`^\d{3,4}$`, data["cvv"])
	if !match {
		return fmt.Errorf("invalid CVV")
	}

	// Validate email
	if !cv.IsValidEmail(data["email"]) {
		return fmt.Errorf("invalid email")
	}

	// Validate expiration month
	match, _ = regexp.MatchString(`^(0?[1-9]|1[012])$`, data["expiration_month"])
	if !match {
		return fmt.Errorf("invalid expiration month")
	}

	// Validate expiration year
	currentYear := time.Now().Year()
	year, err := strconv.Atoi(data["expiration_year"])
	if err != nil || year < currentYear {
		return fmt.Errorf("invalid expiration year")
	}

	// Check if the card is expired
	expDate, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", data["expiration_year"], data["expiration_month"]))
	if err != nil || expDate.Before(time.Now()) {
		return fmt.Errorf("card has expired")
	}

	return nil
}

// chargeValidation validates the charge data
func (cv *CulqiValidation) chargeValidation(data map[string]interface{}) error {
	// Validate email
	if email, ok := data["email"].(string); ok {
		if !cv.IsValidEmail(email) {
			return errors.New("invalid email")
		}
	} else {
		return errors.New("email must be a string")
	}

	// Validate amount
	if amount, ok := data["amount"].(float64); !ok || int(amount) != int(data["amount"].(float64)) {
		return errors.New("invalid amount")
	}

	// Validate currency code
	if currencyCode, ok := data["currency_code"].(string); ok {
		if err := validateCurrencyCode(currencyCode); err != nil {
			return err
		}
	} else {
		return errors.New("currency code must be a string")
	}

	// Validate source ID
	if sourceID, ok := data["source_id"].(string); ok {
		if err := validateStringStart(sourceID, "tkn"); err != nil {
			return err
		}
	} else {
		return errors.New("source ID must be a string")
	}

	return nil
}

// refundValidation validates the refund data
func refundValidation(data map[string]interface{}) error {
	// Validate charge ID format
	if chargeID, ok := data["charge_id"].(string); ok {
		if err := validateStringStart(chargeID, "chr"); err != nil {
			return err
		}
	} else {
		return errors.New("charge ID must be a string")
	}

	// Validate reason
	if reason, ok := data["reason"].(string); ok {
		allowedValues := []string{"duplicado", "fraudulento", "solicitud_comprador"}
		if err := validateValue(reason, allowedValues); err != nil {
			return err
		}
	} else {
		return errors.New("reason must be a string")
	}

	// Validate amount
	if amount, ok := data["amount"].(float64); !ok || int(amount) != int(data["amount"].(float64)) {
		return errors.New("invalid amount")
	}

	return nil
}

func (cv *CulqiValidation) customerValidation(data map[string]interface{}) error {
	requiredFields := []string{"first_name", "last_name", "address", "address_city"}
	for _, field := range requiredFields {
		if _, ok := data[field].(string); !ok || data[field] == "" {
			return fmt.Errorf("%s is empty", field)
		}
	}

	// Validate country code
	countryCode, ok := data["country_code"].(string)
	if !ok {
		return errors.New("country code must be a string")
	}
	if err := validateValue(countryCode, utils.GetCountryCodes()); err != nil {
		return err
	}

	// Validate email
	email, ok := data["email"].(string)
	if !ok || !cv.IsValidEmail(email) {
		return errors.New("invalid email")
	}

	return nil
}

// cardValidation validates card data
func cardValidation(data map[string]interface{}) error {
	if err := validateStringStart(data["customer_id"].(string), "cus"); err != nil {
		return err
	}
	if err := validateStringStart(data["token_id"].(string), "tkn"); err != nil {
		return err
	}
	return nil
}

// subscriptionValidation validates subscription data
func subscriptionValidation(data map[string]interface{}) error {
	if err := validateStringStart(data["card_id"].(string), "crd"); err != nil {
		return err
	}
	if err := validateStringStart(data["plan_id"].(string), "pln"); err != nil {
		return err
	}
	return nil
}

// orderValidation validates order data
func (cv *CulqiValidation) orderValidation(data map[string]interface{}) error {
	amount, ok := data["amount"].(float64)
	if !ok || int(amount) != int(data["amount"].(float64)) {
		return errors.New("invalid amount")
	}

	if err := validateCurrencyCode(data["currency_code"].(string)); err != nil {
		return err
	}

	clientDetails, ok := data["client_details"].(map[string]interface{})
	if !ok {
		return errors.New("client details must be a map")
	}

	requiredClientFields := []string{"first_name", "last_name", "phone_number"}
	for _, field := range requiredClientFields {
		if _, ok := clientDetails[field].(string); !ok || clientDetails[field] == "" {
			return fmt.Errorf("%s is empty", field)
		}
	}

	email, ok := clientDetails["email"].(string)
	if !ok || !cv.IsValidEmail(email) {
		return errors.New("invalid email")
	}

	expirationDate, ok := data["expiration_date"].(string)
	if !ok {
		return errors.New("expiration date must be a string")
	}
	isFuture, err := isFutureDate(expirationDate)
	if err != nil || !isFuture {
		return errors.New("expiration_date must be a future date")
	}

	return nil
}

// IsValidCardNumber checks if the card number is valid
func (cv *CulqiValidation) IsValidCardNumber(number string) bool {
	match, _ := regexp.MatchString(`^\d{13,19}$`, number)
	return match
}

// IsValidEmail checks if the email is valid
func (cv *CulqiValidation) IsValidEmail(email string) bool {
	match, _ := regexp.MatchString(`^\S+@\S+\.\S+$`, email)
	return match
}

func validateCurrencyCode(currencyCode string) error {
	if currencyCode == "" {
		return errors.New("currency code is empty")
	}

	allowedValues := []string{"PEN", "USD"}
	for _, v := range allowedValues {
		if currencyCode == v {
			return nil
		}
	}

	return fmt.Errorf("currency code must be either \"PEN\" or \"USD\", got %s", currencyCode)
}

// validateStringStart checks if a string starts with a specific prefix followed by "_test_" or "_live_"
func validateStringStart(s, start string) error {
	if !strings.HasPrefix(s, start+"_test_") && !strings.HasPrefix(s, start+"_live_") {
		return fmt.Errorf("incorrect format. The format must start with %s_test_ or %s_live_", start, start)
	}
	return nil
}

// validateValue checks if a value is in a list of allowed values
func validateValue(value string, allowedValues []string) error {
	for _, v := range allowedValues {
		if value == v {
			return nil
		}
	}

	allowedValuesJSON, _ := json.Marshal(allowedValues)
	return fmt.Errorf("invalid value. It must be one of %s, got %s", string(allowedValuesJSON), value)
}

// isFutureDate checks if a given date is in the future
func isFutureDate(expirationDate string) (bool, error) {
	date, err := time.Parse("2006-01-02", expirationDate)
	if err != nil {
		return false, fmt.Errorf("invalid date format: %v", err)
	}
	return date.After(time.Now()), nil
}
