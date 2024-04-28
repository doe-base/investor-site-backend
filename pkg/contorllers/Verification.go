package controller

import (
	"context"
	"encoding/json"
	"investor-site/pkg/config"
	"investor-site/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type InputedCode struct {
	Paymentid string `json:"paymentid"`
	Code      string `json:"code"`
}
type VerifySuccessMessage struct {
	Success       bool   `json:"success"`
	PaymentID     string `json:"paymentid"`
	Email         string `json:"email"`
	IsVerified    bool   `json:"isverified"`
	PaymentMethod string `json:"paymentmethod"`
	Price         string `json:"price"`
	CurrencyPrice string `json:"currencyprice"`
	DisplayPrice  string `json:"displayprice"`
}

func Verification(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newInputedCode InputedCode
	json.NewDecoder(r.Body).Decode(&newInputedCode)

	if newInputedCode.Paymentid != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		theCollection := config.InitialPaymentCollection()

		filter := bson.M{"paymentid": newInputedCode.Paymentid}
		var paymentObject PaymentObject
		err := theCollection.FindOne(ctx, filter).Decode(&paymentObject)

		if err != nil {
			var newVerifySuccessMessage VerifySuccessMessage
			newVerifySuccessMessage.Success = false

			json.NewEncoder(w).Encode(newVerifySuccessMessage)
			panic(err)
		} else {
			if newInputedCode.Code == strconv.Itoa(paymentObject.PaymentToken) {
				//change isverified to true
				filter := bson.M{"paymentid": newInputedCode.Paymentid}
				update := bson.M{
					"$set": bson.M{
						"verified": true,
					},
				}
				_, err := theCollection.UpdateOne(ctx, filter, update)
				if err != nil {
					var newVerifySuccessMessage VerifySuccessMessage
					newVerifySuccessMessage.Success = false

					json.NewEncoder(w).Encode(newVerifySuccessMessage)
					panic(err)
				} else {
					//put the opbject in verified payment collection
					var paymentObject PaymentObject
					err = theCollection.FindOne(ctx, filter).Decode(&paymentObject)
					if err != nil {
						var newVerifySuccessMessage VerifySuccessMessage
						newVerifySuccessMessage.Success = false

						json.NewEncoder(w).Encode(newVerifySuccessMessage)
						panic(err)
					} else {
						GetSeriousPaymentCollection := config.SeriousPaymentCollection()
						_, err := GetSeriousPaymentCollection.InsertOne(ctx, paymentObject)
						if err != nil {
							var newVerifySuccessMessage VerifySuccessMessage
							newVerifySuccessMessage.Success = false

							json.NewEncoder(w).Encode(newVerifySuccessMessage)
							panic(err)
						} else {
							var newVerifySuccessMessage VerifySuccessMessage
							newVerifySuccessMessage.Success = true
							newVerifySuccessMessage.PaymentID = paymentObject.PaymentID
							newVerifySuccessMessage.Email = paymentObject.Email
							newVerifySuccessMessage.IsVerified = paymentObject.Verified
							newVerifySuccessMessage.PaymentMethod = paymentObject.Payment
							newVerifySuccessMessage.Price = paymentObject.Price
							newVerifySuccessMessage.CurrencyPrice = paymentObject.CurrencyPrice
							newVerifySuccessMessage.DisplayPrice = paymentObject.DisplayPrice

							json.NewEncoder(w).Encode(newVerifySuccessMessage)
						}

					}
				}

			} else {
				var newVerifySuccessMessage VerifySuccessMessage
				newVerifySuccessMessage.Success = false

				json.NewEncoder(w).Encode(newVerifySuccessMessage)
			}

		}

	}
}
