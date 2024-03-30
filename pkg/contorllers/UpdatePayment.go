package controller

import (
	"context"
	"encoding/json"
	"investor-site/pkg/config"
	"investor-site/pkg/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type UpdateInput struct {
	PaymentChoice string `json:"paymentchoice"`
	Paymentid     string `json:"paymentid"`
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newUpdateInput UpdateInput
	json.NewDecoder(r.Body).Decode(&newUpdateInput)

	if newUpdateInput.Paymentid != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		theCollection := config.InitialPaymentCollection()

		filter := bson.M{"paymentid": newUpdateInput.Paymentid}
		update := bson.M{"$set": bson.M{"paymentpost.payment": newUpdateInput.PaymentChoice}}
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

	}
}
