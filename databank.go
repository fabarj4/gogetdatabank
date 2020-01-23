package gogetdatabank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	var result []*Data
	url := "https://www.atmbersama.com/layanan"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	temp := string(body)
	// fmt.Println(temp)

	// out := strings.TrimLeft(strings.TrimRight(temp, "];"), "var bank_code = [")
	// fmt.Println(out)
	startIndex := strings.Index(temp, "bank_code")
	lastIndex := 0
	for i := startIndex; i < len(temp); i++ {
		if string(temp[i]) == "]" {
			lastIndex = i
			break
		}
	}
	if lastIndex == 0 {
		return nil, fmt.Errorf("error get data from atm bersama")
	}
	jsonString := ""
	if startIndex != -1 && lastIndex != -1 {
		jsonString = temp[startIndex+12 : lastIndex+1]
	}
	fmt.Println(jsonString)
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		return result, err
	}
	return result, nil
}
