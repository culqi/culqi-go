package main

import (
    "fmt"
    culqi "github.com/culqi/culqi-go"
    charge "github.com/culqi/culqi-go/charge"
)

func main() {

  // 1. Configuración
  config := &culqi.Config{
    MerchantCode:   "pk_test_J0BnI4vcidMGdPkF",  // Código de Comercio
    ApiKey:   "sk_test_ujVxc7JMCr0ivKQV", // API Key
    //ApiVersion: "v2",   // No es un parametro necesario, por defecto es la v2
  }

  // 2. Crea un nuevo cliente
  client := culqi.New(config)

  // 3. Parametros
  params := &charge.ChargeParams{
    TokenId:"tkn_test_JsHFDfdc5gfzaSBP",
    FirstName: "William",
    LastName: "Muro",
    Email: "willi@me.com",
    Address: "Avenida Lima 34234",
    AddressCity: "Lima",
    PhoneNumber: 123456787,
    CountryCode: "PE",
    CurrencyCode: "PEN",
    Amount: 1000,
    Installments :0,
    ProductDescription: "Venta de prueba",
  }

  // 4. Crear Cargo
  resp, err := charge.Create(params, client)

  if err != nil {
      panic(err.Error())
  }

  fmt.Printf("\nError: %v", err)
  fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
  fmt.Printf("\nResponse Status: %v", resp.Status())
  fmt.Printf("\nResponse Time: %v", resp.Time())
  fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
  fmt.Printf("\nResponse Body: %v", resp)


}
