package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const BASE_ADDR = "https://apck8ii297.execute-api.eu-central-1.amazonaws.com/dev"

type Data struct {
	Login 		string 	`json:"login"`
	Password 	string 	`json:"password"`
	Type 		int 	`json:"type"`
	OrderID 	string 	`json:"orderId"`
	MerchantId 	string 	`json:"merchantId"`
	Amount		int 	`json:"amount"`
	Description string  `json:"description"`
}

type Cashout struct {
	Type        int    	`json:"type"`
	OrderID     string 	`json:"orderId"`
	MerchantId  string 	`json:"merchantId"`
	Amount      int    	`json:"amount"`
	Description string 	`json:"description"`
	CallbackURL string 	`json:"callbackUrl"`
	ReturnURL string 	`json:"callbackUrl"`
}

type Callback struct {
	OrderID string	`json:"orderId"`
	ID	string	`json:"id"`
	Account	string	`json:"account"`
	Amount int	`json:"amount"`
	Commision int `json:"commision"`
	CommisionIncluded bool `json:"commisionIncluded"`
	Attempt int	`json:"attempt"`
	ReturnURL string	`json:"returnUrl"`
	CallbackUrl string	`json:"callbackUrl"`
	Date string	`json:"date"`
	DateOut string `json:"dateOut"`
	Status int `json:"status"`
	ErrCode int `json:"errCode"`
	ErrMessage string `json:"errMessage"`
	Metadata MetaData `json:"metadata"`
}

type MetaData struct {
	session string `json:"session"`
}

type Result struct {
	URL string `json:"url"`
	ID string `json:"id"`
}

// Code for refactoring //

func Payment(w http.ResponseWriter, r *http.Request) {
	var data Data
	log.Println(data)
	_ = json.NewDecoder(r.Body).Decode(&data)
	log.Println(data)
	res := CreateCashout(data)
	log.Println(res)
	json.NewEncoder(w).Encode(res)
}

func CreateCashout(data Data) Result {
	cashout := Cashout{
		Type: data.Type,
		OrderID: data.OrderID,
		MerchantId: data.MerchantId,
		Amount: data.Amount,
		Description: data.Description,
		CallbackURL: "http://localhost:8081/callback",
		}
	log.Println(cashout)
	c , err := json.Marshal(cashout)
	if err != nil {
		fmt.Println("shit")
	}

	req, _ := http.NewRequest("POST", BASE_ADDR + "/payment/create", bytes.NewBuffer(c))
	authHeader := base64.StdEncoding.EncodeToString([]byte(data.Login + ":" + data.Password))
	req.Header.Add("Authorization", "Basic " + authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println(resp)
	log.Println("------------------")
	var result Result
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

func GetCallback(w http.ResponseWriter, r *http.Request) {
	var callback Callback
	log.Println(callback)
	_ = json.NewDecoder(r.Body).Decode(&callback)
	log.Println(callback)
	log.Println("------------------")
}