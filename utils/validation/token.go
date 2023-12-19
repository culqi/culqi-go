package culqi

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	utils "github.com/culqi/culqi-go/utils"
)

// CulqiValidation contains methods for validating card data
type TokenValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewTokenValidation() *TokenValidation {
	return &TokenValidation{}
}

func (t *TokenValidation) Create(data map[string]string) error {
	// Validate card number
	cardNumber := data["card_number"]
	if !IsValidCardNumber(cardNumber) {
		return NewCustomError("Invalid card number.")
	}

	// Validate CVV
	cvv := data["cvv"]
	match, _ := regexp.MatchString(`^\d{3,4}$`, cvv)
	if !match {
		return NewCustomError("Invalid CVV.")
	}

	// Validate email
	email := data["email"]
	if !IsValidEmail(email) {
		return NewCustomError("Invalid email.")
	}

	// Validate expiration month
	expMonth := data["expiration_month"]
	match, _ = regexp.MatchString(`^(0?[1-9]|1[012])$`, expMonth)
	if !match {
		return NewCustomError("invalid expiration month")
	}

	// Validate expiration year
	currentYear := time.Now().Year()
	expYear := data["expiration_year"]
	year, err := strconv.Atoi(expYear)
	if err != nil || year < currentYear {
		return NewCustomError("invalid expiration year")
	}

	// Check if the card is expired
	expDate, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", data["expiration_year"], data["expiration_month"]))
	if err != nil || expDate.Before(time.Now()) {
		return NewCustomError("card has expired")
	}

	return nil
}

func CreateTokenYape(data map[string]interface{}) error {
	// Retrieve and validate the 'amount' field
	amount, exists := data["amount"]
	if !exists {
		return NewCustomError("Amount field is missing.")
	}

	switch amt := amount.(type) {
	case float64:
		if amt != float64(int(amt)) {
			return NewCustomError("Invalid amount.")
		}
	case int:
		// Valid case, no action needed
	default:
		return NewCustomError("Invalid amount type.")
	}

	return nil
}

func TokenListValidation(data map[string]string) error {
	if _, exists := data["device_type"]; exists {
		allowedDeviceValues := []string{"desktop", "mobile", "tablet"}
		err := ValidateValue(data["device_type"], allowedDeviceValues)
		if err != nil {
			return err
		}
	}
	if _, exists := data["card_brand"]; exists {
		allowedDeviceValues := []string{"Visa", "Mastercard", "Amex", "Diners"}
		err := ValidateValue(data["card_brand"], allowedDeviceValues)
		if err != nil {
			return err
		}
	}
	if _, exists := data["card_type"]; exists {
		allowedDeviceValues := []string{"credito", "debito", "internacional"}
		err := ValidateValue(data["card_type"], allowedDeviceValues)
		if err != nil {
			return err
		}
	}
	if _, exists := data["country_code"]; exists {
		allowedDeviceValues := utils.GetCountryCodes()
		err := ValidateValue(data["country_code"], allowedDeviceValues)
		if err != nil {
			return err
		}
	}
	if _, existsFrom := data["creation_date_from"]; existsFrom {
		if _, existsTo := data["creation_date_to"]; existsTo {
			date_from, _ := strconv.ParseInt(data["creation_date_from"], 10, 64)
			date_to, _ := strconv.ParseInt(data["creation_date_to"], 10, 64)
			err := ValidateDateFilter(date_from, date_to)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
