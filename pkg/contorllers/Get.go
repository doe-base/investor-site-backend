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

func GetAvailableProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.GetProducts()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 0
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 0
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetApexAccountBoostingProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.ApexAccountBoosting()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 01
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 01
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}
func GetApexStackedAccountProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.ApexStackedAccount()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 02
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 02
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetCODStackedAccountProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.CODStackedAccount()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 03
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 03
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}
func GetCODUnlockAllToolsProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.CODUnlockAllTools()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 04
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 04
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetGTALevelRpBoostingProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.GTALevelRpBoosting()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 05
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 05
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetGTAModMenuProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.GTAModMenu()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 06
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 06
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}
func GetGTAMoneyDropProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.GTAMoneyDrop()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 07
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 07
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}
func GetGTAStackedAccountProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.GTAStackedAccount()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 8
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 8
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}
func GetFortniteStackedAccountProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.FortniteStackedAccount()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 9
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 9
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

// Get Reviews
func GetReviews(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.ReviewsCollection()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 9
		newfailureMessage.Message = "could not get reviews"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 9
			newfailureMessage.Message = "could not get reviews"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

// Get Vulchers
func GetVulchers(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.VulchersCollection()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 9
		newfailureMessage.Message = "could not get valchers"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 9
			newfailureMessage.Message = "could not get valchers"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetR6StackedAccountProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.R6StackedAccount()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 02
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 02
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetPokemonStackedAccountProducts(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.PokemonStackedAccount()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 02
		newfailureMessage.Message = "could not get products"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 02
			newfailureMessage.Message = "could not get products"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetCryptoUpdates(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.CryptoUpdates()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		var newfailureMessage FailureMessage
		newfailureMessage.Success = false
		newfailureMessage.ErrorNumber = 03
		newfailureMessage.Message = "could not get crypto updates"

		json.NewEncoder(w).Encode(newfailureMessage)
		panic(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 03
			newfailureMessage.Message = "could not get crypto updates"

			json.NewEncoder(w).Encode(newfailureMessage)
			panic(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetInitialPaymentData(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.InitialPaymentCollection()

	var newPaymentId ResendVerficationCodeObject
	json.NewDecoder(r.Body).Decode(&newPaymentId)

	if newPaymentId.PaymentID != "" {

		filter := bson.M{"paymentid": newPaymentId.PaymentID}
		var newPaymentObject PaymentObject
		err2 := theCollection.FindOne(ctx, filter).Decode(&newPaymentObject)

		if err2 != nil {
			var newfailureMessage FailureMessage
			newfailureMessage.Success = false
			newfailureMessage.ErrorNumber = 4
			newfailureMessage.Message = "fail to get payment obhect"

			json.NewEncoder(w).Encode(newfailureMessage)
			fmt.Println(newfailureMessage.Message)
		} else {

			theCollection := config.CryptoUpdates()
			cursor, err := theCollection.Find(ctx, bson.M{})
			if err != nil {
				fmt.Println(err)
			} else {
				var content []bson.M
				if err = cursor.All(ctx, &content); err != nil {
					var newfailureMessage FailureMessage
					newfailureMessage.Success = false
					newfailureMessage.ErrorNumber = 03
					newfailureMessage.Message = "could not get crypto updates"

					json.NewEncoder(w).Encode(newfailureMessage)
					panic(err)
				} else {
					var newGetOrderDataResult GetOrderDataResult
					newGetOrderDataResult.PaymentObject = newPaymentObject
					newGetOrderDataResult.CryptoRates = content
					json.NewEncoder(w).Encode(newGetOrderDataResult)
				}
			}

		}

	}

}
