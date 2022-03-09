package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/strikersk/go-mongo/src/Entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
