package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/strikersk/go-mongo/src/Entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func GetTodo(ID string) (outputResult Entity.TodoStructure) {
	hexedID, _ := primitive.ObjectIDFromHex(ID)

	tmpResult := getCollection().FindOne(context.TODO(), bson.M{"_id": hexedID})
	_ = tmpResult.Decode(&outputResult)

	return
}

func CreateTodo(inputTask Entity.TodoStructure) (output string) {
	inputTask.Id = uuid.New().String()
	result, _ := getCollection().InsertOne(context.TODO(), inputTask)

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		output = oid.String()
	}

	return
}

func UpdateTodo(ID string, inputTask Entity.TodoStructure) {
	hexedID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Printf("%v", err)
	}

	_, err = getCollection().ReplaceOne(context.TODO(), bson.M{"_id": hexedID}, inputTask)
	if err != nil {
		log.Printf("%v", err)
	}

	return
}
