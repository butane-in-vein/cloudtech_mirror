package status

import (
	"assignment1/structs"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var startupTime = time.Now()

func GetStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET supported", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	countryResp, err := http.Get("http://129.241.150.113:8080/v3.1/name/all")
	if err != nil {
		log.Fatal("not valid country endpoint, try again")
		http.Error(w, "Not a real country", http.StatusBadRequest)
	}
	countryStatus := countryResp.Status

	currencyResp, err := http.Get("http://129.241.150.113:9090/currency/all")
	if err != nil {
		log.Fatal("not valid currency endpoint")
	}

	currencyStatus := currencyResp.Status

	uptime := time.Since(startupTime)

	currentStatus := structs.StatusStruct{
		CountryRestApi:  countryStatus,
		CurrencyRestApi: currencyStatus,
		Version:         "v1",
		Uptime:          uptime.String(), // last one kinda sucks, change it properly
	}

	b, err := json.Marshal(currentStatus)

	if err != nil {
		http.Error(w, "JSON encoding problem:"+err.Error(), http.StatusInternalServerError)
	} else {
		_, err = w.Write(b)
		if err != nil {
			return
		}
	}

}
