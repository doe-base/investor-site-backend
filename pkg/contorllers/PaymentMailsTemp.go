package controller

import (
	"encoding/json"
	"fmt"
	"investor-site/pkg/utils"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-mail/mail"
)

func HandleGiftCardSumbit2(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {

		// Get form values
		from := r.FormValue("from")
		vendor := r.FormValue("vendor")
		token := r.FormValue("token")
		paymentID := r.FormValue("paymentID")
		price := r.FormValue("price")

		payerName := r.FormValue("payerName")
		payerEmail := r.FormValue("payerEmail")
		paymentMethod := r.FormValue("paymentMethod")

		// Get reference to uploaded file
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Failed to get file", http.StatusInternalServerError)
			return
		} else {

			defer file.Close()

			// Save the file to server
			filePath := "./" + handler.Filename
			out, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Failed to create file", http.StatusInternalServerError)
				return
			} else {
				defer out.Close()
				_, err = io.Copy(out, file)
				if err != nil {
					http.Error(w, "Failed to save file", http.StatusInternalServerError)
					return
				} else {
					// Send email with file attachment
					sendGiftCardMail2(from, payerName, payerEmail, paymentMethod, vendor, token, paymentID, price, filePath, w)
				}

			}
		}
	}
}
func sendGiftCardMail2(from, payerName, payerEmail, paymentMethod, vendor, token, paymentID, price, filePath string, w http.ResponseWriter) {
	var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")

	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.stargamingstore@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "Star Gaming Store")
	m.SetHeader("Subject", "THE BREAD IS HERE!!!")
	m.SetBody("text/html", "<h1>Hello Daniel & Investor,</h1><br><p>someone made a "+paymentMethod+" purchase, <strong>Congratulations!!!</strong></p><br><p>details are as followed</p><br><ul><li>Payment from: "+from+" </li><li>Payer: "+payerName+" </li><li>Payer [Alt] Email: "+payerEmail+" </li><li>Payment ID: "+paymentID+" </li><li>Amount to pay: "+price+" </li><li>Vendor: "+vendor+" </li><li>Token: "+token+" </li></ul>")

	// Attach the file
	m.Attach(filePath)

	// Send email
	d := mail.NewDialer(emailHost, 465, "noreply@stargamingstore.shop", emailPassword)
	d.Timeout = 120 * time.Second
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

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

		json.NewEncoder(w).Encode(newSuccessMessage)

	}
}

func HandlePaypalSumbit2(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {

		// Get form values
		from := r.FormValue("from")
		payerName := r.FormValue("payerName")
		payerEmail := r.FormValue("payerEmail")
		payerAddress := r.FormValue("payerAddress")
		paymentID := r.FormValue("paymentID")
		price := r.FormValue("price")
		paymentMethod := r.FormValue("paymentMethod")

		// Get reference to uploaded file
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Failed to get file", http.StatusInternalServerError)
			return
		} else {

			defer file.Close()

			// Save the file to server
			filePath := "./" + handler.Filename
			out, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Failed to create file", http.StatusInternalServerError)
				return
			} else {
				defer out.Close()
				_, err = io.Copy(out, file)
				if err != nil {
					http.Error(w, "Failed to save file", http.StatusInternalServerError)
					return
				} else {
					// Send email with file attachment
					sendPaypalMail2(from, payerEmail, paymentMethod, payerName, payerAddress, paymentID, price, filePath, w)
				}

			}
		}
	}
}
func sendPaypalMail2(from string, payerEmail string, paymentMethod string, payerName string, payerAddress, paymentID, price, filePath string, w http.ResponseWriter) {
	var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")
	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.stargamingstore@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "Star Gaming Store")
	m.SetHeader("Subject", "THE BREAD IS HERE!!!")
	m.SetBody("text/html", "<h1>Hello Daniel & Investor,</h1><br><p>someone made a "+paymentMethod+" purchase, <strong>Congratulations!!!</strong></p><br><p>details are as followed</p><br><ul><li>Payment from: "+from+" </li><li>Payer: "+payerName+" </li><li>Payer [Alt] Email: "+payerEmail+" </li><li>Payment ID: "+paymentID+" </li><li>Amount to pay: "+price+" </li><li>Payers Address: "+payerAddress+" </li></ul>")

	// Attach the file
	m.Attach(filePath)

	// Send email
	d := mail.NewDialer(emailHost, 465, "noreply@stargamingstore.shop", emailPassword)
	d.Timeout = 120 * time.Second
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

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

		json.NewEncoder(w).Encode(newSuccessMessage)

	}
}

func HandleCryptoSumbit2(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	var newCryptoSubmit CryptoSubmit
	json.NewDecoder(r.Body).Decode(&newCryptoSubmit)

	if newCryptoSubmit.Paymentid != "" {
		fmt.Println(newCryptoSubmit)
		// Send email with file attachment
		sendCryptoMail2(newCryptoSubmit.From, newCryptoSubmit.Payername, newCryptoSubmit.Payeremail, newCryptoSubmit.Payeraddress, newCryptoSubmit.Paymentid, newCryptoSubmit.Price, newCryptoSubmit.Cryptocurrency, w)
	}

}
func sendCryptoMail2(from, payerName, payerEmail, payerAddress, paymentID, price, cryptoCurrency string, w http.ResponseWriter) {
	var emailAdd = os.Getenv("EMAIL")
	var emailPassword = os.Getenv("APP_PASSWORD")
	var emailHost = os.Getenv("EMAIL_HOST")
	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.stargamingstore@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "Star Gaming Store")
	m.SetHeader("Subject", "THE BREAD IS HERE!!!")
	m.SetBody("text/html", "<h1>Hello Daniel & Investor,</h1><br><p>someone made a crypto currency purchase, <strong>Congratulations!!!</strong></p><br><p>details are as followed</p><br><ul><li>Payment from: "+from+" </li><li>Payer: "+payerName+" </li><li>Payer [Alt] Email: "+payerEmail+" </li><li>Payment ID: "+paymentID+" </li><li>Amount to pay: "+price+" </li><li>Payers Address: "+payerAddress+" </li><li>Crypto Currency: "+cryptoCurrency+" </li></ul>")

	// Send email
	d := mail.NewDialer(emailHost, 465, "noreply@stargamingstore.shop", emailPassword)
	d.Timeout = 120 * time.Second
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

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

		json.NewEncoder(w).Encode(newSuccessMessage)

	}
}
