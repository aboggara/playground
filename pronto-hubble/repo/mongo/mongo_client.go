package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"pronto-hubble/config"
)

type MongoClient struct {
	db *mongo.Database
}

func Init(ctx context.Context, dbConfig config.DbConfig) *MongoClient {
	dbClient, err := initializeMongo(ctx, dbConfig)

	if err != nil {
		log.Println("Unable to get the mongo client: %s\n", err.Error())
	}
	return &MongoClient{db: dbClient}
}

func initializeMongo(ctx context.Context, dbConfig config.DbConfig) (*mongo.Database, error) {

	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Database,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client couldn't connect with background context: %v", err)
	}
	prontoDb := client.Database(dbConfig.Database)
	return prontoDb, nil
}

func (mongoClient *MongoClient)  InsertOne(ctx context.Context, collection string, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error){
	result, err := mongoClient.db.Collection(collection).InsertOne(ctx, document)
	if err != nil {
		log.Println("Unable to insert the data into the mongo db: %s", err.Error())
		return nil, err
	}
	if result != nil {
		log.Println("Result is ", result)
		return result, nil
	}
	return nil, nil
}
