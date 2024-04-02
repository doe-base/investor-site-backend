package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"investor-site/pkg/config"
	"investor-site/pkg/utils"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-mail/mail"
)

type PaymentPost struct {
	Name                string `json:"name"`
	Email               string `json:"email"`
	Country             string `json:"country"`
	Payment             string `json:"payment"`
	Delivery            string `json:"delivery"`
	PromoCode           string `json:"promocode"`
	AccountEmail        string `json:"accountemail"`
	AccountUsername     string `json:"accountusername"`
	AccountActivationId string `json:"accountactivationid"`

	ServiceTitle  string  `json:"servicetitle"`
	PackageName   string  `json:"packagename"`
	Price         int     `json:"price"`
	OrigianlPrice float64 `json:"originalprice"`
	Game          string  `json:"game"`
	SericeCode    string  `json:"servicecode"`
}
type SuccessMessage struct {
	Success    bool   `json:"success"`
	PaymentID  string `json:"paymentid"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isverified"`
}

type FailureMessage struct {
	Success     bool   `json:"success"`
	ErrorNumber int    `json:"errornumber"`
	Message     string `json:"message"`
}

type BillObj struct {
	Id         int32  `json:"id"`
	Completed  bool   `json:"completed"`
	PaymentFor string `json:"paymentfor"`
	Status     string `json:"status"`
}

type PaymentObject struct {
	PaymentPost
	Verified     bool   `json:"verified"`
	PaymentID    string `json:"paymentid"`
	PaymentToken int    `json:"paymenttoken"`

	Status    string `json:"status"`
	Date      string `json:"date"`
	Completed bool   `json:"completed"`

	Bills []BillObj `json:"bills"`
}

func generateRandomNumber(digits int) int {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(digits)
	return randomNumber
}

func paymentIdGenerator(paymentIdNumber int, serviceCode string, game string) string {
	var gameNumber string
	if game == "pc" {
		gameNumber = "1"
	} else if game == "ps" {
		gameNumber = "2"
	} else {
		gameNumber = "3"
	}
	var id = "ORD-" + strconv.Itoa(paymentIdNumber) + "-" + strings.ToUpper(serviceCode) + gameNumber

	return id
}

func HandleChoosePaymentMethodSubmit(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	paymentIdNumber := generateRandomNumber(1000000000)
	paymentToken := generateRandomNumber(1000000)

	var newPayment PaymentPost
	json.NewDecoder(r.Body).Decode(&newPayment)

	if newPayment.Name != "" {
		paymentId := paymentIdGenerator(paymentIdNumber, newPayment.SericeCode, newPayment.Game)
		newDate := time.Now()
		var newPaymentObject PaymentObject
		newPaymentObject.Verified = false
		newPaymentObject.PaymentID = paymentId
		newPaymentObject.PaymentToken = paymentToken
		newPaymentObject.Name = newPayment.Name
		newPaymentObject.Email = newPayment.Email
		newPaymentObject.Country = newPayment.Country
		newPaymentObject.Payment = newPayment.Payment
		newPaymentObject.Delivery = newPayment.Delivery
		newPaymentObject.PromoCode = newPayment.PromoCode
		newPaymentObject.AccountEmail = newPayment.AccountEmail
		newPaymentObject.AccountUsername = newPayment.AccountUsername
		newPaymentObject.AccountActivationId = newPayment.AccountActivationId
		newPaymentObject.ServiceTitle = newPayment.ServiceTitle
		newPaymentObject.PackageName = newPayment.PackageName
		newPaymentObject.Price = newPayment.Price
		newPaymentObject.Status = "Pending"
		newPaymentObject.Completed = false
		newPaymentObject.Date = strconv.Itoa(newDate.Day()) + " " + newDate.Month().String() + ", " + strconv.Itoa(newDate.Year())

		GetInitialPaymentCollection := config.InitialPaymentCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := GetInitialPaymentCollection.InsertOne(ctx, newPaymentObject)
		if err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 1
			newfailureMessage.Message = "fail to insert payment"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {

			m := mail.NewMessage()
			m.SetHeader("From", os.Getenv("EMAIL"))
			m.SetHeader("To", newPayment.Email)
			m.SetAddressHeader("Cc", os.Getenv("EMAIL"), "Star Gaming Store")
			m.SetHeader("Subject", "Payment Confirmation")
			m.SetBody("text/html", "Hello <b>"+newPayment.Name+"</b> <br></br><br></br><br></br> You have initiated a payment for "+newPayment.ServiceTitle+": <br></br><br></br> Payment Id: <b>"+paymentId+"</b>   <br></br> Package: <b>"+newPayment.PackageName+"</b>   <br></br> Price: <b>$"+strconv.Itoa(newPayment.Price)+".00</b>   <br></br> Delivery Method: <b>"+newPayment.Delivery+"</b> <br></br> Payment Method: <b>"+newPayment.Payment+"</b>   <br></br><br></br>  Use the code below to authenticate payment <br></br> <h1>"+strconv.Itoa(paymentToken)+"</h1> <br></br><br></br> If you didn't initiate this payment, kindly delete this message, Thank you. Star Gaming Store support team.")

			d := mail.NewDialer(os.Getenv("EMAIL_HOST"), 465, os.Getenv("EMAIL"), os.Getenv("APP_PASSWORD"))
			d.StartTLSPolicy = mail.MandatoryStartTLS

			if err := d.DialAndSend(m); err != nil {
				var newfailureMessage FailureMessage
				newfailureMessage.Success = false
				newfailureMessage.ErrorNumber = 2
				newfailureMessage.Message = "fail to send mall"

				json.NewEncoder(w).Encode(newfailureMessage)
				fmt.Println(newfailureMessage.Message)
				fmt.Println(err)
				panic(err)
			} else {

				var newSuccessMessage SuccessMessage
				newSuccessMessage.Success = true
				newSuccessMessage.PaymentID = paymentId
				newSuccessMessage.Email = newPayment.Email
				newSuccessMessage.IsVerified = newPaymentObject.Verified

				json.NewEncoder(w).Encode(newSuccessMessage)

			}
		}

	}

}
