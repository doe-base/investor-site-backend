package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"investor-site/pkg/config"
	"investor-site/pkg/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-mail/mail"
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
							var emailPassword = os.Getenv("APP_PASSWORD")
							var emailHost = os.Getenv("EMAIL_HOST")
							var emailAdd = os.Getenv("EMAIL")
							// Create a new mailer
							m := mail.NewMessage()
							m.SetHeader("From", emailAdd)
							m.SetHeader("To", "stargamingstoree@gmail.com")
							m.SetAddressHeader("Cc", emailAdd, "Star Gaming Store")
							m.SetHeader("Subject", "THE BREAD IS HERE!!!")
							m.SetBody("text/html", "Hello <b>"+paymentObject.Name+"</b>,<br>Your order for <b>"+paymentObject.PackageName+"</b> is ready and available for purchase.<br><br></br><br></br> Payment Id: <b>"+paymentObject.PaymentID+"</b>   <br></br> Package: <b>"+paymentObject.PackageName+"</b>    <br></br> Console/PC: <b>"+paymentObject.Game+"</b>   <br></br> Price: <b>"+paymentObject.DisplayPrice+"</b>   <br></br> Delivery Method: <b>"+paymentObject.Delivery+"</b> <br></br> Payment Method: <b>"+paymentObject.Payment+"</b></b> <br></br> Account ID: <b>"+strconv.Itoa(paymentObject.PackageTire)+"</b><br><br>To make the process smooth and convenient, we can generate a payment link for you. Just let us know if you're ready to proceed by replying this mail, and we'll send the link directly to your email. <br/><br/> Thanks again for your order")

							// Send email
							d := mail.NewDialer(emailHost, 465, "noreply@stargamingstore.shop", emailPassword)
							d.Timeout = 120 * time.Second
							d.StartTLSPolicy = mail.MandatoryStartTLS

							if err := d.DialAndSend(m); err != nil {
								var newfailureMessage FailureMessage
								newfailureMessage.Success = false
								newfailureMessage.ErrorNumber = 2
								newfailureMessage.Message = "fail to send mall"

								json.NewEncoder(w).Encode(newfailureMessage)
								fmt.Println(err)
								// panic(err)
							} else {

								var newSuccessMessage SuccessMessage
								newSuccessMessage.Success = true

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
				}

			} else {
				var newVerifySuccessMessage VerifySuccessMessage
				newVerifySuccessMessage.Success = false

				json.NewEncoder(w).Encode(newVerifySuccessMessage)
			}

		}

	}
}
