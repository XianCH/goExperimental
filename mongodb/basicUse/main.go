package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOption := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	collection := client.Database("basicUse").Collection("basicUse-collection")

	_, err = collection.InsertOne(context.TODO(), bson.M{"name": "pi", "value": 3.14159})

	if err != nil {
		log.Fatal(err)
	}

	// Update
	filter := bson.M{"name": "pi"}
	update := bson.M{"$set": bson.M{"value": 3.14}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// Replace
	replace := bson.M{"name": "e", "value": 2.71828}
	_, err = collection.ReplaceOne(context.TODO(), filter, replace)
	if err != nil {
		log.Fatal(err)
	}

	// Aggregate
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"name", "pi"}}}},
		{{"$group", bson.D{{"_id", "$name"}, {"total", bson.D{{"$sum", "$value"}}}}}},
	}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

}
