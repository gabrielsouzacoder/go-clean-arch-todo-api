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
	_, _ = r.db.InsertOne(ctx, e)

	var s = new(*entity.ID)

	return *s, nil
}

func (r *MongoDb) List() ([]*entity.Todo, error) {
	var tasks []*entity.Todo

	cur, err := r.db.Find(ctx, bson.D{{}})
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t entity.Todo
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	// once exhausted, close the cursor
	erre := cur.Close(ctx)
	if erre != nil {
		return nil, erre
	}

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

