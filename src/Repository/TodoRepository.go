package Repository

import (
	"context"
	"github.com/strikersk/go-mongo/src/Entity"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func ReadTodo(ID string) (outputResult Entity.TodoStructure) {
	tmpResult := GetCollection().FindOne(context.Background(), bson.M{"id": ID})
	_ = tmpResult.Decode(&outputResult)
	return
}

func FindAll() []Entity.TodoStructure {
	outputTasks := make([]Entity.TodoStructure, 0)
	tmpResult, err := GetCollection().Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("%v\n", err)
	}

	for tmpResult.Next(context.Background()) {
		var tmp Entity.TodoStructure
		if err = tmpResult.Decode(&tmp); err != nil {
			log.Printf("%v\n", err)
		}
		outputTasks = append(outputTasks, tmp)
	}

	return outputTasks
}

func CreateTodo(inputTask Entity.TodoStructure) (err error) {
	_, err = GetCollection().InsertOne(context.Background(), inputTask)
	return
}

func UpdateTodo(ID string, inputTask Entity.TodoStructure) {
	_, err := GetCollection().ReplaceOne(context.Background(), bson.M{"id": ID}, inputTask)
	if err != nil {
		log.Printf("%v", err)
	}

	return
}
