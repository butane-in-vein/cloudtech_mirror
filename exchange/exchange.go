package exchange

import (
	"assignment1/structs"
	"encoding/json"
	"net/http"
)

func GetExchange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET supported", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	countryCode := r.PathValue("code")
	if len(countryCode) != 2 {
		http.Error(w, "Country codes are 2 letters only.", http.StatusBadRequest)
		return
	}
	countryUrl := "http://129.241.150.113:8080/v3.1/alpha/" + countryCode + "?fields=name,borders,currencies"
	countryResp, err := http.Get(countryUrl)
	if err != nil {
		http.Error(w, "Error connecting to RestCountries API", http.StatusBadGateway)
		return
	}
	if countryResp.StatusCode != http.StatusOK {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}
	defer countryResp.Body.Close()

	//borderCountries := structs.ExchangeInfo{}
	var base structs.ExchangeInfo
	if err := json.NewDecoder(countryResp.Body).Decode(&base); err != nil {
		http.Error(w, "Error decoding base country", http.StatusInternalServerError)
		return
	}

	// base country currency
	baseCurrency := ""
	for k := range base.Currencies {
		baseCurrency = k
		break
	}
	if baseCurrency == "" {
		http.Error(w, "Missing base country currency", http.StatusBadGateway)
		return
	}

	// over to rates
	ratesResp, ratesErr := http.Get("http://129.241.150.113:9090/currency/" + baseCurrency)
	if ratesErr != nil {
		http.Error(w, "Error fetching rates from API", http.StatusBadGateway)
		return
	}
	defer ratesResp.Body.Close()

	if ratesResp.StatusCode != http.StatusOK {
		http.Error(w, "Currency API error", http.StatusBadGateway)
		return
	}

	var fetchedRates structs.CurrencyAPIResponse
	if err := json.NewDecoder(ratesResp.Body).Decode(&fetchedRates); err != nil {
		http.Error(w, "Error decoding rates", http.StatusInternalServerError)
		return
	}

	relevantRates := make(map[string]float64)
	// bordering country
	for _, borders := range base.Borders {
		borderResp, err := http.Get("http://129.241.150.113:8080/v3.1/alpha/" + borders + "?fields=currencies")
		if err != nil {
			continue
		}

		func() {
			defer borderResp.Body.Close()

			if borderResp.StatusCode != http.StatusOK {
				return
			}

			var neighbor structs.ExchangeInfo
			if err := json.NewDecoder(borderResp.Body).Decode(&neighbor); err != nil {
				return
			}

			// neighbor currency
			neighborCurrency := ""
			for k := range neighbor.Currencies {
				neighborCurrency = k
				break
			}
			if neighborCurrency == "" || neighborCurrency == baseCurrency {
				return
			}

			//if rate exists
			if rate, ok := fetchedRates.Rates[neighborCurrency]; ok {
				relevantRates[neighborCurrency] = rate
			}
		}()
	}

	exchangeResponse := structs.ExchangeResponse{
		Country:       base.Name.Common,
		BaseCurrency:  baseCurrency,
		ExchangeRates: relevantRates,
	}
	_ = json.NewEncoder(w).Encode(exchangeResponse)

}
