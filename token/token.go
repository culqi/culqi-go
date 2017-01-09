package token

import (
     "gopkg.in/resty.v0"
     culqi "github.com/brayancruces/culqi-go"
     "errors"
)

const (
	tokensBase = "tokens"
)

type TokenParams struct {
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Email string `json:"email"`
  Currency string `json:"currency"`
  CardNumber int `json:"card_number"`
  Cvv int `json:"cvv"`
  ExpMonth int `json:"expiration_month"`
  ExpYear int `json:"expiration_year"`
}


type TokenResponse struct {
  // Object: "token"
  Object string `json:"object"`
  Id string `json:"id"`
  CardNumber string `json:"card_number"`
  CardHolder *CardHolder `json:"cardholder` // opt: map[string]string
  BrandName string `json:"brand_name`
  BankName string `json:"bank_name`
  BankCountry string `json:"bank_country`
}

type CardHolder struct {
  // Object: "cardholder"
  Object string `json:"object"`
  FirstName string `json:"first_name"`
  Email string `json:"email"`
  LastName string `json:"last_name"`
}


func Create(params *TokenParams, configuracion *culqi.Config)  (*resty.Response, error) {

  if configuracion.ApiVersion == "" {
    configuracion.ApiVersion = "v2"
  }

  // Valores de configuración
  //baseurl := configuracion.BaseURL+configuracion.ApiVersion
  codcomercio := configuracion.MerchantCode
  //userAgent := configuracion.UserAgent

  //Validación parámetros
  if len(params.FirstName) == 0 {
     err := errors.New("El campo FirstName se encuentra vacío.")
     return nil, err
  }

  resp, err := resty.R().
      SetHeader("Content-Type","application/json").
      SetHeader("Authorization","Code "+codcomercio).
      SetHeader("User-Agent", culqi.UserAgent).
      SetBody(params).
      Post(culqi.DefaultBaseURL+"v2/"+tokensBase)
  return resp, err

}
