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

type SuccessObjectData struct {
	PaymentAddressObject PaymentAddressObject `json:"paymentaddressobject"`
	TempMailObject       TempMailObject       `json:"tempmailobject"`
}
type TempMailObject struct {
	TempPaypal []primitive.M `json:"temppaypal"`
}

// type InputedPaymentId struct {
// 	PaymentId string `json:"paymentid"`
// }

func PaymentCheckerFroPersonal(w http.ResponseWriter, r *http.Request) {
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

					theCollection := config.GetTempPaypal()
					cursor, err := theCollection.Find(ctx, bson.M{})
					if err != nil {
						//Error

						var newSuccessObjectData SuccessObjectData
						newSuccessObjectData.PaymentAddressObject = newPaymentAddressObject

						json.NewEncoder(w).Encode(newSuccessObjectData)
					} else {
						var content []bson.M
						if err = cursor.All(ctx, &content); err != nil {
							//Error

							var newSuccessObjectData SuccessObjectData
							newSuccessObjectData.PaymentAddressObject = newPaymentAddressObject

							json.NewEncoder(w).Encode(newSuccessObjectData)
						} else {
							var newTempMailObject TempMailObject
							newTempMailObject.TempPaypal = content

							var newSuccessObjectData SuccessObjectData
							newSuccessObjectData.PaymentAddressObject = newPaymentAddressObject
							newSuccessObjectData.TempMailObject = newTempMailObject

							json.NewEncoder(w).Encode(newSuccessObjectData)
						}
					}
				}
			}
		}
	}
}
