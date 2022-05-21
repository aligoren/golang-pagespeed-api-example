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
	"sync"
	"time"
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

func getPagespeedResult(url string) models.PageSpeedResponseModel {

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
		panic(err)
	}

	return jsonData
}

/*func main() {

	jsonStruct, _ := get_pagespeed_result("https://aligoren.com")

	// fmt.Printf("Site ID: %s", result.ID)
	fmt.Printf("Full Result: %#v", jsonStruct.LighthouseResult.Audits["unused-css-rules"].Details)
}*/

func main() {
	urls := []string{"https://github.com", "https://www.enuygun.com", "https://aligoren.com", "https://coderwall.com",
		"https://eksisozluk.com", "https://medium.com", "https://stackoverflow.com", "https://www.youtube.com"}

	results := make(chan models.PageSpeedResponseModel)
	calculateTime := make(chan time.Duration)

	go func() {
		start := time.Now()
		wg := sync.WaitGroup{}

		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				result := getPagespeedResult(url)
				results <- result

			}(url)
		}
		wg.Wait()

		close(results)

		elapsedTime := time.Since(start)
		calculateTime <- elapsedTime

		close(calculateTime)
	}()

	for result := range results {
		fmt.Printf("%s\n", result.ID)
	}

	fmt.Printf("Elapsed time: %s", <-calculateTime)
}
