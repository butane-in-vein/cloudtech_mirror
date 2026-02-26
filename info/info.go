package info

import (
	"assignment1/structs"
	"encoding/json"
	"log"
	"net/http"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET supported", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//url := "http://129.241.150.113:8080/v3.1/alpha/"
	countryCode := r.PathValue("code")
	if len(countryCode) != 2 {
		http.Error(w, "Country codes are 2 letters only.", http.StatusBadRequest)
		return
	}
	countryUrl := "http://129.241.150.113:8080/v3.1/alpha/" + countryCode + "?fields=name,continents,population,area,languages,borders,flag,capital"
	countryInfo, err := http.Get(countryUrl)
	if len(r.PathValue("code")) != 2 {
		http.Error(w, "Code must be 2 letters. Try "+"no"+" (norway) or "+"us"+" (united states)", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Error connecting to RestCountries API", http.StatusBadGateway)
		return
	}
	defer countryInfo.Body.Close()
	if countryInfo.StatusCode != http.StatusOK {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	var country structs.CountryStruct
	if err := json.NewDecoder(countryInfo.Body).Decode(&country); err != nil {
		log.Println("decode error:", err)
		http.Error(w, "Error decoding information", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(country)

}
