package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	// ! DON'T FORGET TO ENV THIS
	var clientOptions = os.Getenv("MONGODB_CONNECTION_STRING")

	//** CONNECTING TO MONGODB & SERVE
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(clientOptions))
	if err != nil {
		panic(err)
	} else {
		Client = mongoClient
	}
}

func GetProducts() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("products")
	return theCollection
}

func ApexAccountBoosting() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("apexAccountBooting")
	return theCollection
}
func ApexStackedAccount() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("apexStackedAccount")
	return theCollection
}

func CODStackedAccount() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("codStackedAccount")
	return theCollection
}
func CODUnlockAllTools() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("codUnlockAllTools")
	return theCollection
}

func FortniteStackedAccount() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("fortniteStackedAccount")
	return theCollection
}

func GTALevelRpBoosting() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("gtaLevelRpBoosting")
	return theCollection
}
func GTAModMenu() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("gtaModMenu")
	return theCollection
}
func GTAMoneyDrop() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("gtaMoneyDrop")
	return theCollection
}
func GTAStackedAccount() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("gtaStackedAccount")
	return theCollection
}
func R6StackedAccount() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("r6StackedAccount")
	return theCollection
}
func PokemonStackedAccount() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("pokemonStackedAccount")
	return theCollection
}
func CryptoUpdates() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("cryptoupdates")
	return theCollection
}

func InitialPaymentCollection() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("initialPayment")
	return theCollection
}
func SeriousPaymentCollection() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("seriousPayment")
	return theCollection
}
func PromoCodeCollection() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("promoCodes")
	return theCollection
}

func ReviewsCollection() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("reviews")
	return theCollection
}

func VulchersCollection() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("vulchers")
	return theCollection
}

func AddressCollection() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("address")
	return theCollection
}

func CartCollection() *mongo.Collection {
	theCollection := Client.Database("account-tools").Collection("cart")
	return theCollection
}
func InitialPaymentCollectionForCart() *mongo.Collection {
	theCollection := Client.Database("base1").Collection("initialPaymentForCart")
	return theCollection
}
