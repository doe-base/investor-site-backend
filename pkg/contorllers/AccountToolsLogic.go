package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"investor-site/pkg/config"
	"investor-site/pkg/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentID struct {
	Paymentid string `json:"paymentid"`
}
type ValidatedPaymentID struct {
	IsValidated bool `json:"isvalidated"`
}
type CartItem struct {
	Id     int32  `json:"id"`
	Incart bool   `json:"incart"`
	Name   string `json:"name"`
	Price  int32  `json:"price"`
	Game   string `json:"game"`
}
type CartSubmitInput struct {
	TempPaymentID  string     `json:"temppaymentid"`
	PaymentidInput string     `json:"paymentidinput"`
	Promocode      string     `json:"promocode"`
	CartList       []CartItem `json:"cartlist"`
	TotalAmount    int32      `json:"totalamount"`
	DiscountAmount string     `json:"discountamount"`
	IsPaid         bool       `json:"ispaid"`
}
type TempPaymentID struct {
	TempPaymentID string `json:"temppaymentid"`
}
type GetOrderDataResult struct {
	PaymentObject   PaymentObject
	CartSubmitInput CartSubmitInput
	CryptoRates     []primitive.M
}

func ValidatePaymentID(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newInputedCode PaymentID
	json.NewDecoder(r.Body).Decode(&newInputedCode)

	if newInputedCode.Paymentid != "" {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		theCollection := config.InitialPaymentCollection()

		filter := bson.M{"paymentid": newInputedCode.Paymentid}
		var paymentObject PaymentObject
		err := theCollection.FindOne(ctx, filter).Decode(&paymentObject)
		var newValidatedPaymentID ValidatedPaymentID
		if err != nil {
			newValidatedPaymentID.IsValidated = false

			json.NewEncoder(w).Encode(newValidatedPaymentID)
			fmt.Println(err)
		} else {
			newValidatedPaymentID.IsValidated = true

			json.NewEncoder(w).Encode(newValidatedPaymentID)
		}
	}
}

func CartSubmit(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newCartSubmitInput CartSubmitInput
	json.NewDecoder(r.Body).Decode(&newCartSubmitInput)

	if newCartSubmitInput.TempPaymentID != "" {
		var newCartDBSubmit CartSubmitInput

		newCartDBSubmit.TempPaymentID = newCartSubmitInput.TempPaymentID
		newCartDBSubmit.PaymentidInput = newCartSubmitInput.PaymentidInput
		newCartDBSubmit.Promocode = newCartSubmitInput.Promocode
		newCartDBSubmit.CartList = newCartSubmitInput.CartList
		newCartDBSubmit.TotalAmount = newCartSubmitInput.TotalAmount
		newCartDBSubmit.DiscountAmount = newCartSubmitInput.DiscountAmount
		newCartDBSubmit.IsPaid = newCartSubmitInput.IsPaid

		cartCollection := config.CartCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := cartCollection.InsertOne(ctx, newCartDBSubmit)

		var newValidatedPaymentID ValidatedPaymentID
		if err != nil {
			newValidatedPaymentID.IsValidated = false
			json.NewEncoder(w).Encode(newValidatedPaymentID)
			fmt.Println(err)
		} else {
			newValidatedPaymentID.IsValidated = true
			json.NewEncoder(w).Encode(newValidatedPaymentID)
		}
	}
}

func GetOrderData(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newTempPaymentID TempPaymentID
	json.NewDecoder(r.Body).Decode(&newTempPaymentID)

	if newTempPaymentID.TempPaymentID != "" {
		cartCollection := config.CartCollection()
		initialPaymentCollection := config.InitialPaymentCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{"temppaymentid": newTempPaymentID.TempPaymentID}
		var cartSubmitInputOpject CartSubmitInput
		err := cartCollection.FindOne(ctx, filter).Decode(&cartSubmitInputOpject)
		if err != nil {
			fmt.Println(err)
		} else {
			filter := bson.M{"paymentid": cartSubmitInputOpject.PaymentidInput}
			var newPaymentObject PaymentObject
			err := initialPaymentCollection.FindOne(ctx, filter).Decode(&newPaymentObject)
			if err != nil {
				fmt.Println(err)
			} else {
				theCollection := config.CryptoUpdates()
				cursor, err := theCollection.Find(ctx, bson.M{})
				if err != nil {
					fmt.Println(err)
				} else {
					var content []bson.M
					if err = cursor.All(ctx, &content); err != nil {
						var newfailureMessage FailureMessage
						newfailureMessage.Success = false
						newfailureMessage.ErrorNumber = 03
						newfailureMessage.Message = "could not get crypto updates"

						json.NewEncoder(w).Encode(newfailureMessage)
						panic(err)
					} else {
						var newGetOrderDataResult GetOrderDataResult
						newGetOrderDataResult.CartSubmitInput = cartSubmitInputOpject
						newGetOrderDataResult.PaymentObject = newPaymentObject
						newGetOrderDataResult.CryptoRates = content
						json.NewEncoder(w).Encode(newGetOrderDataResult)
					}
				}
			}
		}
	}
}
