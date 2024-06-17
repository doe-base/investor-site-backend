package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"investor-site/pkg/config"
	"investor-site/pkg/utils"
	"net/http"
	"time"
)

type MailList struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MailListSub(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	// var newMailList MailList
	// json.NewDecoder(r.Body).Decode(&newMailList)

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name != "" {
		fmt.Println("hitted")

		var newMailList MailList

		newMailList.Name = name
		newMailList.Email = email

		mailListCollection := config.MailList()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := mailListCollection.InsertOne(ctx, newMailList)

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
