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

type PaymentAddressObject struct {
	PaymentObject PaymentObject `json:"paymentobject"`
	Address       []primitive.M `json:"address"`
}

type InputedPaymentId struct {
	PaymentId string `json:"paymentid"`
}

func PaymentChecker(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newInputedPaymentId InputedPaymentId
	json.NewDecoder(r.Body).Decode(&newInputedPaymentId)

	if newInputedPaymentId.PaymentId != "" {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		theCollection := config.InitialPaymentCollection()

		filter := bson.M{"paymentid": newInputedPaymentId.PaymentId}
		var paymentObject PaymentObject
		err := theCollection.FindOne(ctx, filter).Decode(&paymentObject)
		var newVerifySuccessMessage VerifySuccessMessage
		if err != nil {
			newVerifySuccessMessage.Success = false

			json.NewEncoder(w).Encode(newVerifySuccessMessage)
			fmt.Println(err)
		} else {
			var newPaymentAddressObject PaymentAddressObject
			newPaymentAddressObject.PaymentObject = paymentObject

			theCollection := config.AddressCollection()
			cursor, err := theCollection.Find(ctx, bson.M{})

			if err != nil {
				newVerifySuccessMessage.Success = false

				json.NewEncoder(w).Encode(newVerifySuccessMessage)
				fmt.Println(err)
			} else {
				var content []bson.M
				if err = cursor.All(ctx, &content); err != nil {
					newVerifySuccessMessage.Success = false

					json.NewEncoder(w).Encode(newVerifySuccessMessage)
					fmt.Println(err)
				} else {
					newPaymentAddressObject.Address = content
					json.NewEncoder(w).Encode(newPaymentAddressObject)
				}
			}
		}
	}
}
