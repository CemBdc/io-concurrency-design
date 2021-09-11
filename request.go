package main

import (
	"io/ioutil"
	"net/http"
)

func getRequest(url string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	failOnError(err, "Error on wrapping request")

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	failOnError(err, "Error on while requesting")

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	failOnError(err, "Error on while getting response")

	return body
}
