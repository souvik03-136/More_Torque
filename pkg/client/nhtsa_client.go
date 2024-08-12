package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DecodeVIN(vin string) (interface{}, error) {
	// Call NHTSA API
	resp, err := http.Get(fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/decodevin/%s?format=json", vin))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Process the response
	return body, nil
}
