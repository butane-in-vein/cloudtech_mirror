package structs

type StatusStruct struct {
	CountryRestApi  string `json:"StatusCodeCountriesAPI"`
	CurrencyRestApi string `json:"StatusCodeCurrenciesAPI"`
	Version         string `json:"Version"`
	Uptime          string `json:"uptime in seconds"`
}

type CurrencyAPIResponse struct {
	BaseCode string             `json:"base_code"`
	Rates    map[string]float64 `json:"rates"`
}

type ExchangeResponse struct {
	Country       string             `json:"country"`
	BaseCurrency  string             `json:"base-currency"`
	ExchangeRates map[string]float64 `json:"exchange-rates"`
}

type ExchangeInfo struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Borders    []string                   `json:"borders"`
	Currencies map[string]CurrencyDetails `json:"currencies"`
}

type CurrencyDetails struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type CountryStruct struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    []string          `json:"capital"`
}

type Rates struct {
	Exchanges map[string]float64 `json:"rates"`
}

type CountryResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		PNG string `json:"png"`
		SVG string `json:"svg"`
	} `json:"flags"`
	Capital    []string       `json:"capital"`
	Currencies map[string]any `json:"currencies"`
}
