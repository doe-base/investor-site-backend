package main

import (
	"fmt"
	"investor-site/pkg/config"
	controller "investor-site/pkg/contorllers"

	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Only for local
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// 	panic(err)
	// }

	var router *mux.Router = mux.NewRouter()

	router.HandleFunc("/get-products", controller.GetAvailableProducts).Methods("GET")
	router.HandleFunc("/apex-account-boosting", controller.GetApexAccountBoostingProducts).Methods("GET")
	router.HandleFunc("/apex-stacked-account", controller.GetApexStackedAccountProducts).Methods("GET")
	router.HandleFunc("/call-of-duty-stacked-account", controller.GetCODStackedAccountProducts).Methods("GET")
	router.HandleFunc("/call-of-duty-unlock-all-tools", controller.GetCODUnlockAllToolsProducts).Methods("GET")
	router.HandleFunc("/fortnite-stacked-account", controller.GetFortniteStackedAccountProducts).Methods("GET")
	router.HandleFunc("/grand-theft-auto-level-rp-boosting", controller.GetGTALevelRpBoostingProducts).Methods("GET")
	router.HandleFunc("/grand-theft-auto-mod-menu", controller.GetGTAModMenuProducts).Methods("GET")
	router.HandleFunc("/grand-theft-auto-money-drop", controller.GetGTAMoneyDropProducts).Methods("GET")
	router.HandleFunc("/grand-theft-auto-stacked-account", controller.GetGTAStackedAccountProducts).Methods("GET")
	router.HandleFunc("/rainbow-six-stacked-account", controller.GetR6StackedAccountProducts).Methods("GET")
	router.HandleFunc("/pokemon-go-stacked-account", controller.GetPokemonStackedAccountProducts).Methods("GET")
	router.HandleFunc("/get-reviews", controller.GetReviews).Methods("GET")
	router.HandleFunc("/get-vulchers", controller.GetVulchers).Methods("GET")

	router.HandleFunc("/gift-card-payment", controller.HandleGiftCardSumbit).Methods("POST", "OPTIONS")
	router.HandleFunc("/paypal-payment", controller.HandlePaypalSumbit).Methods("POST", "OPTIONS")
	router.HandleFunc("/crypto-currency-payment", controller.HandleCryptoSumbit).Methods("POST", "OPTIONS")

	router.HandleFunc("/subit-form-request", controller.HandleChoosePaymentMethodSubmit).Methods("POST", "OPTIONS")
	router.HandleFunc("/verification", controller.Verification).Methods("POST", "OPTIONS")
	router.HandleFunc("/resend-verification", controller.HandleResendVerificationCode).Methods("POST", "OPTIONS")

	router.HandleFunc("/promo-code-checker", controller.PromoCodeChecker).Methods("POST", "OPTIONS")
	router.HandleFunc("/payment-checker", controller.PaymentChecker).Methods("POST", "OPTIONS")
	router.HandleFunc("/update-payment", controller.UpdatePayment).Methods("POST", "OPTIONS")

	router.HandleFunc("/get-crypto-updates", controller.GetCryptoUpdates).Methods("GET")
	router.HandleFunc("/validate-paymentid", controller.ValidatePaymentID).Methods("POST", "OPTIONS")
	router.HandleFunc("/cart-submit", controller.CartSubmit).Methods("POST", "OPTIONS")
	router.HandleFunc("/get-checkout-info", controller.GetOrderData).Methods("POST", "OPTIONS")
	router.HandleFunc("/get-checkout-info-fot-account-tools", controller.GetOrderDataForAccountTools).Methods("POST", "OPTIONS")

	router.HandleFunc("/get-initial-payment-data", controller.GetInitialPaymentData).Methods("POST", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.Connect()
	fmt.Printf("Starting server on port.............. %d \n", 8080)
	http.ListenAndServe(":"+port, router)

}
