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
)

type UpdateInput struct {
	PaymentChoice string `json:"paymentchoice"`
	Paymentid     string `json:"paymentid"`
	DisplayPrice  string `json:"displayprice"`
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newUpdateInput UpdateInput
	json.NewDecoder(r.Body).Decode(&newUpdateInput)

	if newUpdateInput.Paymentid != "" {
		fmt.Println(newUpdateInput.DisplayPrice)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		theCollection := config.InitialPaymentCollection()

		filter := bson.M{"paymentid": newUpdateInput.Paymentid}
		var paymentObject PaymentObject
		err := theCollection.FindOne(ctx, filter).Decode(&paymentObject)
		if err != nil {
			var newVerifySuccessMessage VerifySuccessMessage
			newVerifySuccessMessage.Success = false

			json.NewEncoder(w).Encode(newVerifySuccessMessage)
		} else {
			filter := bson.M{"paymentid": newUpdateInput.Paymentid}
			update := bson.M{"$set": bson.M{"paymentpost.payment": newUpdateInput.PaymentChoice, "paymentpost.displayprice": newUpdateInput.DisplayPrice}}
			_, err := theCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				var newVerifySuccessMessage VerifySuccessMessage
				newVerifySuccessMessage.Success = false

				json.NewEncoder(w).Encode(newVerifySuccessMessage)
				panic(err)
			} else {
				var newVerifySuccessMessage VerifySuccessMessage
				newVerifySuccessMessage.Success = true

				json.NewEncoder(w).Encode(newVerifySuccessMessage)
			}
			// NO need for condition
			// if newUpdateInput.PaymentChoice == "crypto currency" {

			// 	filter := bson.M{"paymentid": newUpdateInput.Paymentid}
			// 	update := bson.M{"$set": bson.M{"paymentpost.payment": newUpdateInput.PaymentChoice, "paymentpost.displayprice": newUpdateInput.DisplayPrice}}
			// 	_, err := theCollection.UpdateOne(ctx, filter, update)
			// 	if err != nil {
			// 		var newVerifySuccessMessage VerifySuccessMessage
			// 		newVerifySuccessMessage.Success = false

			// 		json.NewEncoder(w).Encode(newVerifySuccessMessage)
			// 		panic(err)
			// 	} else {
			// 		var newVerifySuccessMessage VerifySuccessMessage
			// 		newVerifySuccessMessage.Success = true

			// 		json.NewEncoder(w).Encode(newVerifySuccessMessage)
			// 	}
			// } else {

			// 	filter := bson.M{"paymentid": newUpdateInput.Paymentid}
			// 	update := bson.M{"$set": bson.M{"paymentpost.payment": newUpdateInput.PaymentChoice, "paymentpost.displayprice": paymentObject.Price}}
			// 	_, err := theCollection.UpdateOne(ctx, filter, update)
			// 	if err != nil {
			// 		var newVerifySuccessMessage VerifySuccessMessage
			// 		newVerifySuccessMessage.Success = false

			// 		json.NewEncoder(w).Encode(newVerifySuccessMessage)
			// 		panic(err)
			// 	} else {
			// 		var newVerifySuccessMessage VerifySuccessMessage
			// 		newVerifySuccessMessage.Success = true

			// 		json.NewEncoder(w).Encode(newVerifySuccessMessage)
			// 	}
			// }
		}

	}
}
