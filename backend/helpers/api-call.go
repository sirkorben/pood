package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ApiCall(article string, dst interface{}) error {
	apiKey := os.Getenv("API_KEY")
	apiLink := os.Getenv("API_LINK")
	interpolatedApiLink := fmt.Sprintf("%vapiKey=%v&article=%v", apiLink, apiKey, article)

	client := http.Client{}
	req, err := http.NewRequest("GET", interpolatedApiLink, nil)
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
