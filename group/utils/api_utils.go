package utils

import (
    "io/ioutil"
    "net/http"
)

// FetchDataFromAPI fetches data from the given API URL
func FetchDataFromAPI(endpoint string) ([]byte, error) {
    resp, err := http.Get(endpoint)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}
