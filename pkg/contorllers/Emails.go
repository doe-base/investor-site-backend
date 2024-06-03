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
	"go.mongodb.org/mongo-driver/bson"
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
	Price         string  `json:"price"`
	OrigianlPrice float64 `json:"originalprice"`

	PackageTire int    `json:"packagetire"`
	Game        string `json:"game"`
	SerivceCode string `json:"servicecode"`

	CurrencyChoose string `json:"currencychoose"`
	CurrencyPrice  string `json:"currencyprice"`
	DisplayPrice   string `json:"displayprice"`
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
	Price      string `json:"price"`
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

type ResendVerficationCodeObject struct {
	PaymentID string `json:"paymentid"`
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
	} else if game == "xbox" {
		gameNumber = "3"
	} else if game == "mobile" {
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
		paymentId := paymentIdGenerator(paymentIdNumber, newPayment.SerivceCode, newPayment.Game)
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
		newPaymentObject.SerivceCode = newPayment.SerivceCode
		newPaymentObject.PackageName = newPayment.PackageName
		newPaymentObject.Game = newPayment.Game
		newPaymentObject.PackageTire = newPayment.PackageTire
		newPaymentObject.Price = newPayment.Price
		newPaymentObject.OrigianlPrice = newPayment.OrigianlPrice
		newPaymentObject.CurrencyChoose = newPayment.CurrencyChoose
		newPaymentObject.CurrencyPrice = newPayment.CurrencyPrice
		newPaymentObject.DisplayPrice = newPayment.DisplayPrice

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
			m.SetHeader("From", "noreply@stargamingstore.shop")
			m.SetHeader("To", newPayment.Email)
			m.SetAddressHeader("Cc", "noreply@stargamingstore.shop", "Star Gaming Store")
			m.SetHeader("Subject", "Payment Confirmation")
			m.SetBody("text/html", "Hello <b>"+newPayment.Name+"</b> <br></br><br></br><br></br> You have initiated a payment for "+newPayment.ServiceTitle+": <br></br><br></br> Payment Id: <b>"+paymentId+"</b>   <br></br> Package: <b>"+newPayment.PackageName+"</b>    <br></br> Console/PC: <b>"+newPayment.Game+"</b>   <br></br> Price: <b>"+newPayment.DisplayPrice+"</b>   <br></br> Delivery Method: <b>"+newPayment.Delivery+"</b> <br></br> Payment Method: <b>"+newPayment.Payment+"</b>   <br></br><br></br>  Use the code below to authenticate payment <br></br> <h1>"+strconv.Itoa(paymentToken)+"</h1> <br></br><br></br> If you didn't initiate this payment, kindly delete this message, Thank you. Star Gaming Store support team.")

			d := mail.NewDialer(os.Getenv("EMAIL_HOST"), 465, "noreply@stargamingstore.shop", os.Getenv("APP_PASSWORD"))
			d.StartTLSPolicy = mail.MandatoryStartTLS

			if err := d.DialAndSend(m); err != nil {
				var newfailureMessage FailureMessage
				newfailureMessage.Success = false
				newfailureMessage.ErrorNumber = 2
				newfailureMessage.Message = "fail to send mall"

				json.NewEncoder(w).Encode(newfailureMessage)
				fmt.Println(newfailureMessage.Message)
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

func HandleResendVerificationCode(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	paymentToken := generateRandomNumber(1000000)

	var newResendVerificationCode ResendVerficationCodeObject
	json.NewDecoder(r.Body).Decode(&newResendVerificationCode)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	theCollection := config.InitialPaymentCollection()
	filter := bson.M{"paymentid": newResendVerificationCode.PaymentID}
	update := bson.M{"$set": bson.M{"paymenttoken": paymentToken}}
	_, err := theCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 3
		newfailureMessage.Message = "fail to send verification code"

		json.NewEncoder(w).Encode(newfailureMessage)
		fmt.Println(newfailureMessage.Message)
	} else {

		filter := bson.M{"paymentid": newResendVerificationCode.PaymentID}
		var newPaymentObject PaymentObject
		err2 := theCollection.FindOne(ctx, filter).Decode(&newPaymentObject)
		if err2 != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 4
			newfailureMessage.Message = "fail to get payment obhect"

			json.NewEncoder(w).Encode(newfailureMessage)
			fmt.Println(newfailureMessage.Message)
		} else {
			m := mail.NewMessage()
			m.SetHeader("From", "noreply@stargamingstore.shop")
			m.SetHeader("To", newPaymentObject.Email)
			m.SetAddressHeader("Cc", "noreply@stargamingstore.shop", "Star Gaming Store")
			m.SetHeader("Subject", "Payment Confirmation")
			m.SetBody("text/html", "Hello <b>"+newPaymentObject.Name+"</b> <br></br><br></br><br></br> You have initiated a payment for "+newPaymentObject.ServiceTitle+": <br></br><br></br> Payment Id: <b>"+newResendVerificationCode.PaymentID+"</b>   <br></br> Package: <b>"+newPaymentObject.PackageName+"</b>    <br></br> Console/PC: <b>"+newPaymentObject.Game+"</b>   <br></br> Price: <b>"+newPaymentObject.DisplayPrice+"</b>   <br></br> Delivery Method: <b>"+newPaymentObject.Delivery+"</b> <br></br> Payment Method: <b>"+newPaymentObject.Payment+"</b>   <br></br><br></br>  Use the code below to authenticate payment <br></br> <h1>"+strconv.Itoa(paymentToken)+"</h1> <br></br><br></br> If you didn't initiate this payment, kindly delete this message, Thank you. Star Gaming Store support team.")

			d := mail.NewDialer(os.Getenv("EMAIL_HOST"), 465, "noreply@stargamingstore.shop", os.Getenv("APP_PASSWORD"))
			d.StartTLSPolicy = mail.MandatoryStartTLS

			if err3 := d.DialAndSend(m); err3 != nil {
				var newfailureMessage FailureMessage
				newfailureMessage.Success = false
				newfailureMessage.ErrorNumber = 5
				newfailureMessage.Message = "fail to send mall"

				json.NewEncoder(w).Encode(newfailureMessage)
				fmt.Println(newfailureMessage.Message)
				panic(err3)
			} else {

				var newSuccessMessage SuccessMessage
				newSuccessMessage.Success = true
				newSuccessMessage.PaymentID = newResendVerificationCode.PaymentID
				newSuccessMessage.Email = newPaymentObject.Email
				newSuccessMessage.IsVerified = newPaymentObject.Verified

				json.NewEncoder(w).Encode(newSuccessMessage)

			}
		}

	}
}
