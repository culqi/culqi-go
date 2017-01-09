package main

import (
    "fmt"
    culqi "github.com/brayancruces/culqi-go"
    token "github.com/brayancruces/culqi-go/token"
    "encoding/json"
)

func main() {

  // 1. Configuración
  config := &culqi.Config{
    MerchantCode:   "pk_test_xxx",  // Código de tu Comercio
    ApiKey:   "sk_test_xxx", // API Key
    //ApiVersion: "v2",   // No es un parametro necesario, por defecto es la v2
  }

  // 2. Crea un nuevo cliente
  client := culqi.New(config)

  // 3. Parametros
  params := &token.TokenParams{
    FirstName: "Brayan",
    LastName: "Cruces",
    Email: "brayancruces@gmail.com",
    Currency: "PEN",
    CardNumber: 4111111111111111,
    Cvv: 123,
    ExpMonth: 9,
    ExpYear: 2020,
  }

  // 4. Crear Token
  resp, err := token.Create(params, client)

  if err != nil {
      panic(err.Error())
  }

  fmt.Printf("\nError: %v", err)
  fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
  fmt.Printf("\nResponse Status: %v", resp.Status())
  fmt.Printf("\nResponse Time: %v", resp.Time())
  fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
  fmt.Printf("\nResponse Body: %v", resp)

  // Leer Data de respuesta por parametro
  data := token.TokenResponse{}
  json.Unmarshal([]byte(resp.Body()), &data)

  // Imprimiendo algunos parametros
  fmt.Println("\nToken ID: "+data.Id)
  fmt.Println("\nCardNumber: "+data.CardNumber)
  fmt.Print("\nEmail de cardholder: "+data.CardHolder.Email)


}
