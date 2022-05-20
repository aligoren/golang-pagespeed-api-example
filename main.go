package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pagespeed_api/models"
)

// API Resource: https://developers.google.com/speed/docs/insights/v5/get-started

var BaseUrl = "https://www.googleapis.com/pagespeedonline/v5/runPagespeed?url=%s&key=%s"
var ApiKey string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file couldn't loaded")
	}

	ApiKey = os.Getenv("API_KEY")
}

func get_pagespeed_result(url string) (models.PageSpeedResponseModel, []byte) {

	newUrl := fmt.Sprintf(BaseUrl, url, ApiKey)

	res, err := http.Get(newUrl)

	if err != nil {
		log.Fatalf("Error: response couldn't read from API source %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error while reading response body %v", err)
	}

	var jsonData models.PageSpeedResponseModel
	err = json.Unmarshal(body, &jsonData)

	if err != nil {
		log.Fatalf("Error while unmarshaling body %v", err)
	}

	jsonForm, err := json.MarshalIndent(jsonData, "", " ")

	if err != nil {
		log.Fatalf("Error while converting structure to json form %v", err)
	}

	return jsonData, jsonForm
}

func main() {

	jsonStruct, _ := get_pagespeed_result("https://aligoren.com")

	// fmt.Printf("Site ID: %s", result.ID)
	fmt.Printf("Full Result: %#v", jsonStruct.LighthouseResult.Audits["unused-css-rules"].Details)
}
