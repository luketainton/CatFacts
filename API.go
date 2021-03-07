package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CatFact is the information provided by the cat-fact API
type CatFact struct {
	Text string `json:"text"`
}

func getFact() CatFact {
	apiEndpoint := "https://cat-fact.herokuapp.com/facts/random?amount=1"
	resp, _ := http.Get(apiEndpoint)
	body, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	var catFact CatFact
	err := json.Unmarshal([]byte(bodyString), &catFact)
	if err != nil {
		fmt.Println(err)
	}
	return catFact
}
