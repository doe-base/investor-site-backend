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
			json.NewEncoder(w).Encode(paymentObject)
		}
	}
}
