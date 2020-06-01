package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Started.")
	client := &http.Client{}
	data, err := json.Marshal(map[string]string{
		"exchangeType":   "CURRENCY",
		"baseAmount":     "5",
		"baseCurrency":   "EUR",
		"resultCurrency": "HUF",
	})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://www.otpbank.hu/apps/exchangerate/api/exchangerate/exchange", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Host", "www.otpbank.hu")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Cache-Control", "no-store, no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Expires", "0")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Content-Length", "86")
	req.Header.Add("Origin", "https://www.otpbank.hu")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://www.otpbank.hu/portal/hu/Arfolyamok/OTP")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
