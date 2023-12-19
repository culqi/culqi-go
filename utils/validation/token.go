package culqi

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	utils "github.com/culqi/culqi-go/utils/validation"
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
		return utils.NewCustomError("Invalid card number.")
	}

	// Validate CVV
	cvv := data["cvv"]
	match, _ := regexp.MatchString(`^\d{3,4}$`, cvv)
	if !match {
		return utils.NewCustomError("Invalid CVV.")
	}

	// Validate email
	email := data["email"]
	if !IsValidEmail(email) {
		return utils.NewCustomError("Invalid email.")
	}

	// Validate expiration month
	expMonth := data["expiration_month"]
	match, _ = regexp.MatchString(`^(0?[1-9]|1[012])$`, expMonth)
	if !match {
		return utils.NewCustomError("invalid expiration month")
	}

	// Validate expiration year
	currentYear := time.Now().Year()
	expYear := data["expiration_year"]
	year, err := strconv.Atoi(expYear)
	if err != nil || year < currentYear {
		return utils.NewCustomError("invalid expiration year")
	}

	// Check if the card is expired
	expDate, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", data["expiration_year"], data["expiration_month"]))
	if err != nil || expDate.Before(time.Now()) {
		return utils.NewCustomError("card has expired")
	}

	return nil
}
