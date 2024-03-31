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

type InputedPropmoCode struct {
	PromoCode string `json:"promocode"`
}

func PromoCodeChecker(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newInputedPropmoCode InputedPropmoCode
	json.NewDecoder(r.Body).Decode(&newInputedPropmoCode)

	if newInputedPropmoCode.PromoCode != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		theCollection := config.PromoCodeCollection()

		filter := bson.M{newInputedPropmoCode.PromoCode: newInputedPropmoCode.PromoCode}
		var promoCode InputedPropmoCode
		err := theCollection.FindOne(ctx, filter).Decode(&promoCode)
		var newVerifySuccessMessage VerifySuccessMessage
		if err != nil {
			newVerifySuccessMessage.Success = false

			json.NewEncoder(w).Encode(newVerifySuccessMessage)
			fmt.Println(err)
		} else {
			newVerifySuccessMessage.Success = true

			json.NewEncoder(w).Encode(newVerifySuccessMessage)
		}
	}
}
