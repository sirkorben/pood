package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CallForPrices(article string, dst interface{}) error {
	//TODO: hide apiKey in env vars
	apiKey := "f22d9a0a6a65a45e5fe9cd652deb0e98e9051b286d36709ae12d1900da516c10"
	interpolatedString := fmt.Sprintf("https://originalparts.pro/api/search?apiKey=%v&article=%v", apiKey, article)
	client := http.Client{}
	req, err := http.NewRequest("GET", interpolatedString, nil)
	if err != nil {
		log.Println("ERROR IN WRAPPING REQUEST TO API", err)
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR IN MAKING REQUEST TO API", err)
		return err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//handle err
		log.Println("READING BODY BYTES ERR", err)
		return err
	}

	json.Unmarshal(bodyBytes, &dst)
	if err != nil {
		//handle err
		log.Println("REQUEST BODY UNMARSHALLING err ", err)
		return err
	}

	return nil
}
