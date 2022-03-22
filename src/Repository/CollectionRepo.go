package Repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

const (
	databaseName   = "appDatabase"
	collectionName = "appCollection"
)

// Example: mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority
var databaseUrl = os.Getenv("DB_HOST")

type MongoInstance struct {
	appCollection *mongo.Collection
}

func GetCollection() *mongo.Collection {
	return GetDatabaseInstance().appCollection
}

var mongoLock = &sync.Mutex{}
var databaseConnection *MongoInstance

func GetDatabaseInstance() *MongoInstance {
	//To prevent expensive lock operations
	//This means that the databaseConnection field is already populated
	if databaseConnection == nil {
		mongoLock.Lock()
		defer mongoLock.Unlock()

		//Only one goroutine can create the singleton instance.
		if databaseConnection == nil {
			log.Println("Creating Database instance")
			var tmpMongoInstance MongoInstance

			client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseUrl))
			if err != nil {
				panic(err)
			}

			log.Printf("Collection %s: created\n", collectionName)
			tmpMongoInstance.appCollection = client.Database(databaseName).Collection(collectionName)

			databaseConnection = &tmpMongoInstance
			log.Println("All collections initialization: done")
		} else {
			log.Println("Application database instance already created")
		}
	} else {
		//log.Println("Application Database instance already created.")
	}

	return databaseConnection
}
