package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

type AccountRepository interface{
	InquiryAccount(ctx context.Context, accountId string) (AccountData, error)
	Close()
}

type MongoClient struct {
	DB *mongo.Database
}

func NewMongoClient() *MongoClient {
	mc := &MongoClient{}
    mc.CreateConnection()
	return mc
}

func (mc *MongoClient) CloseConnection() {
	mc.DB.Client().Disconnect(context.Background())
}

func (mc *MongoClient) CreateConnection() {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	dbName := os.Getenv("MONGODB_DB_NAME")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	client, err := mongo.NewClient(
		options.Client().ApplyURI(connectionURI),
		options.Client().SetMinPoolSize(100),
		options.Client().SetMaxPoolSize(1000))
	if err != nil {
		logrus.Errorf("Failed to create client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Errorf("Failed to connect to server: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Errorf("Failed to ping cluster: %v", err)
	}

	mc.DB = client.Database(dbName)
	logrus.Info("Success !!")
}

func (mc *MongoClient) InquiryAccount(ctx context.Context, accountId string) (*AccountData, error) {
	if mc.DB == nil {
		return &AccountData{}, fmt.Errorf("Connection to DB not established!")
	}

	var account AccountData
	ctx, cancel := initContext()
	defer cancel()

	// objID, _ := primitive.ObjectIDFromHex(accountId)
	collection := mc.DB.Collection("account")
	err := collection.FindOne(ctx, bson.D{{"_id", accountId}}).Decode(&account)
	if err != nil {
		logrus.Errorf("Error reading accountID '%v' from DB: %v", accountId, err)
		return nil, err
	}

	return &account, nil
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}