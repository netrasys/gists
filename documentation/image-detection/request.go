package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

const SUBSCRIPTION_KEY = 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXX'
const CALLBACK_URL = "http://api.webhookinbox.com/i/XXXXXX/"

func main() {
    client := &http.Client{}

    request := map[string]string{
        "image_url": "https://storage.googleapis.com/netra-public-data/async-image-detection/documentation/api-test-image.jpg",
        "callback_url": CALLBACK_URL,
    }

    requestBody, err := json.Marshal(request)
    if err != nil {
        panic(err)
    }

    req, err := http.NewRequest("POST", "https://api.getnetra.com/image-detection/process/brands", bytes.NewBuffer(requestBody))
    if err != nil {
        panic(err)
    }

    req.Header.Add("Ocp-Apim-Subscription-Key", SUBSCRIPTION_KEY)
    req.Header.Add("Content-Type", "application/json")

    response, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    var responseData interface{}
    err = json.NewDecoder(response.Body).Decode(&responseData)
    if err != nil {
        panic(err)
    }

    fmt.Println("Sent request:", response.Status, responseData)
}
