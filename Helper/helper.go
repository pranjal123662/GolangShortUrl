package helper

import (
	controller "ShorUrl/Controller"
	model "ShorUrl/Models"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type JSONData struct {
	Code     string `json:"code"`
	ShortUrl string `json:"shorturl"`
}
type GetUrlData struct {
	Url string `json:"url"`
}

func ConvertIntoShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	startTime := time.Now()
	var receivedData model.ShortUrl
	_ = json.NewDecoder(r.Body).Decode(&receivedData)
	Url := receivedData.Url
	if Url == "" {
		json.NewEncoder(w).Encode(JSONData{Code: "204", ShortUrl: ""})
		return
	}
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Intn(100000))
	receivedData.ShortUrl = "https://updateverse.com/shortUrl/" + id
	res := controller.InsertIntoDataBase(receivedData)
	if res != nil {
		jsonData := JSONData{Code: "200", ShortUrl: receivedData.ShortUrl}
		json.NewEncoder(w).Encode(jsonData)
		endTime := time.Now()
		// Calculate the duration
		duration := endTime.Sub(startTime)
		fmt.Printf("Response time: %v\n", duration)
		return
	}
}
func RedirectToOriginalUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Short Url")
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "*")
	res := controller.FetchDataFromDatabase("https://updateverse.com" + r.URL.String())
	if res == "" {
		w.Write([]byte("<h1>Url not Found</h1>"))
		return
	}
	http.Redirect(w, r, res, http.StatusFound)
}
