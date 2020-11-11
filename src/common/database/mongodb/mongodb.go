package mongodb

import (
	"context"
	"errors"
	"github.com/mongodb/mongo-go-driver/mongo"
	"ipadgrpc/src/common/database"
	"ipadgrpc/src/common/databaseatabase"
	"log"
	"reflect"
	"strconv"
)

type MongoSession struct {
	client *mongo.Client
	database *mongo.Database
	config *database.DatabaseConfig
}
func NewMongoDataBase(config *database.DatabaseConfig) *MongoSession {
	mongoSession := new(MongoSession)
	mongoSession.Init(config)
	return mongoSession
}
func (this *MongoSession) Init(config *database.DatabaseConfig) *MongoSession {
	this.config = config
	return this
}
func (this *MongoSession) Connect() error {
	config := this.config
	url := "mongodb://"
	if config.UserName != "" {
		url += config.UserName + ":" + config.Password + "@"
	}
	url += config.HostName
	if config.HostPort != 0 {
		url += ":" + strconv.Itoa(config.HostPort)
	}
	client, error := mongo.NewClient(url)
	if error != nil {
		return error
	}
	this.client = client
	error = this.client.Connect(context.Background())
	if error != nil {
		return error
	}
	this.database = this.client.Database(config.DataBaseName)
	if this.database == nil {
		return errors.New("连接到"+config.DataBaseName+"失败")
	}
	return nil
}

func (this *MongoSession) collectionNameForObject(object interface{}) *mongo.Collection {
	targetType := reflect.TypeOf(object)
	collectionName := targetType.Elem().Name()
	return this.database.Collection(collectionName)
}

func (this *MongoSession) Insert(object interface{}) error {
	collection := this.collectionNameForObject(object)
	collection.InsertOne(context.Background(), object)
	return nil
}
