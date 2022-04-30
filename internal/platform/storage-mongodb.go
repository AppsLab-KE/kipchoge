package platform

import (
	"appslab-ke/kipchoge-go/configs"
	"appslab-ke/kipchoge-go/internal/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Client struct {
	client *mongo.Client
}

func (c Client) ConnectDb() error {
	panic("implement me")
}

func (c Client) FindAllJokeCategories() ([]models.JokeCategory, error) {
	panic("implement me")
}

func (c Client) Ping() error {
	log.Println("[Ping] Mongodb")

	err := c.client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	log.Println("[Pong] Mongodb")
	return nil
}

func (c Client) Defer() error {
	log.Println("[Disconnecting] Mongodb")
	err := c.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	log.Println("[Disconnected] Mongodb")
	return nil
}

func NewMongoDbConn() (*Client, error) {
	mongoUrl := configs.GetDbConfig()
	if mongoUrl == "" {
		return nil, errors.New("mongodb uri cannot be empty")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}
