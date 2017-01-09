// Package culqi implementa metodos para inicializar cliente
//
// Inicia cliente con datos de condifuración de comercio.
package culqi

const (
	libraryVersion = "0.1.0"
	DefaultBaseURL = "https://api.culqi.com/"
	UserAgent      = "culqi-go/" + libraryVersion
	mediaType      = "application/json"
	format         = "json"
	headerCulqiTrackId  = "x-culqi-tracking-id"
	headerEnvironment = "x-culqi-environment"

)

type Config struct {
  MerchantCode   string
	ApiKey   string
	ApiVersion string
  BaseURL string
  UserAgent string
}

func New(config *Config) *Config {
    // set valores por defecto
    return &Config{
			  MerchantCode: config.MerchantCode,
				ApiKey: config.ApiKey,
				ApiVersion: config.ApiVersion,
        //BaseURL: defaultBaseURL,
        //UserAgent: userAgent,
    }
}
