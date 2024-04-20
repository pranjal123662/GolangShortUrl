package controller

import (
	model "ShorUrl/Models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectString = "mongodb+srv://Pranjal:Pranjal%40123@cluster0.sc7ucqz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "ShortUrlBucket"
const colName = "storedData"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
}

func InsertIntoDataBase(shortUrlData model.ShortUrl) *mongo.InsertOneResult {
	insertData, err := collection.InsertOne(context.Background(), shortUrlData)
	if err != nil {
		log.Fatal(err)
	}
	return insertData
}
func FetchDataFromDatabase(url string) string {
	filter := bson.M{"shorturl": url}
	var result model.ShortUrl
	collection.FindOne(context.Background(), filter).Decode(&result)
	return result.Url
}
