package repository

import (
	"context"
	"fmt"
	"github.com/gabrielsouzacoder/clean-new/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ctx = context.TODO()

type MongoDb struct {
	db *mongo.Collection
}

func NewMongoDbRepository(clientOptions *options.ClientOptions) *MongoDb {
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[Database] Mongo Connected")

	collection := client.Database("tasker").Collection("tasks")

	return &MongoDb{
		db: collection,
	}
}

func (r *MongoDb) Create(e *entity.Todo) (*entity.ID, error) {
	_, err := r.db.InsertOne(ctx, e)

	if err != nil {
		return nil, err
	}

	return &e.ID, nil
}

func (r *MongoDb) List() ([]*entity.Todo, error) {
	var tasks []*entity.Todo

	cur, err := r.db.Find(ctx, bson.D{{}})
	if err != nil {
		return tasks, err
	}

	todos, err := addToList(cur, &tasks)

	if err != nil {
		return todos, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func addToList(cur *mongo.Cursor, tasks *[]*entity.Todo) ([]*entity.Todo, error) {

	var list []*entity.Todo

	for cur.Next(ctx) {
		var t entity.Todo
		err := cur.Decode(&t)
		if err != nil {
			return nil, err
		}

		list = append(*tasks, &t)
	}
	return list, nil
}

func (r *MongoDb) Delete(id *entity.ID) error {
	_, err := r.db.DeleteOne(ctx, bson.D{{
		"id", *id,
	}})

	if err != nil {
		return err
	}

	return nil
}

func (r *MongoDb) FindById(id *entity.ID) *entity.Todo {

	var t *entity.Todo

	err := r.db.FindOne(ctx, bson.D{{
		"id", *id,
	}}).Decode(&t)

	if err != nil {
		return nil
	}

	return t
}
