package main

import (
	"fmt"
	"investor-site/pkg/config"
	controller "investor-site/pkg/contorllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

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
	router.HandleFunc("/subit-form-request", controller.HandleChoosePaymentMethodSubmit).Methods("POST", "OPTIONS")
	router.HandleFunc("/verification", controller.Verification).Methods("POST", "OPTIONS")
	router.HandleFunc("/promo-code-checker", controller.PromoCodeChecker).Methods("POST", "OPTIONS")
	router.HandleFunc("/payment-checker", controller.PaymentChecker).Methods("POST", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.Connect()
	fmt.Printf("Starting server on port.............. %d \n", 8080)
	http.ListenAndServe(":"+port, router)

}
