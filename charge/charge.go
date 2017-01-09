package charge

import (
     "gopkg.in/resty.v0"
     culqi "github.com/brayancruces/culqi-go"
)

const (
	chargesBase = "charges"
)


type ChargeParams struct {
  TokenId string `json:"token_id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Email string `json:"email"`
  Address string `json:"address"`
  AddressCity string `json:"address_city"`
  PhoneNumber int `json:"phone_number"`
  CountryCode string `json:"country_code"`
  CurrencyCode string `json:"currency_code"`
  Amount int `json:"amount"`
  Installments int `json:"installments"`
  ProductDescription string `json:"product_description"`
}

func Create(params *ChargeParams, configuracion *culqi.Config)  (*resty.Response, error) {

  if configuracion.ApiVersion == "" {
    configuracion.ApiVersion = "v2"
  }

  // Valores de configuración
  apikey := configuracion.ApiKey

  resp, err := resty.R().
      SetHeader("Content-Type","application/json").
      SetHeader("Authorization","Bearer "+apikey).
      SetHeader("User-Agent", culqi.UserAgent).
      SetBody(params).
      Post(culqi.DefaultBaseURL+"v2/"+chargesBase)
  return resp, err

}
