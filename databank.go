package gogetdatabank

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func DataBankHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		dataBank, err := DataBank()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := json.Marshal(dataBank)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func DataBank() ([]*Data, error) {
	// var result string
	url := "https://www.atmbersama.com/layanan"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	temp := string(body)
	jsonString := temp[48812:52467]
	var result []*Data
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		return result, err
	}
	return result, nil
}
