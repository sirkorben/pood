package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JokeResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type Response struct {
	Prices []*Oneprice
}
type Oneprice struct {
	Price            float64 `json:"price,omitempty"`
	Article          string  `json:"article,omitempty"`
	Supplier         string  `json:"supplier,omitempty"`
	SupplierPriceNum float64 `json:"supplier_price_num,omitempty"`
	Brand            string  `json:"brand,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	CurrencyRate     string  `json:"currency_rate,omitempty"`
	Delivery         string  `json:"delivery,omitempty"`
	Weight           float64 `json:"weight,omitempty"`
	Name             string  `json:"name,omitempty"`
}

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
		log.Println(err)
		return err
	}

	// taking a pointer variable using
	// := by assigning it with the
	// address of variable models.ApiResponse{}
	// apiResponse := &models.ApiResponse{}
	json.Unmarshal(bodyBytes, &dst)
	if err != nil {
		log.Println("REQUEST BODY UNMARSHALLING err ", err)
		return err
	}

	// returning pointer variable to the address of models.ApiResponse{}
	return nil
}
