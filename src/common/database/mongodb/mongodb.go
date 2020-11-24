package mongodb

import (
	"context"
	"errors"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"qcengine/src/common/database"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type MongoHeader struct {
	database.DataBase
	client *mongo.Client
	database *mongo.Database
	config *database.DatabaseConfig
}

type MongoSession struct {
	*MongoHeader
}
func NewMongoDataBase(config *database.DatabaseConfig) *MongoSession {
	mongoSession := new(MongoSession)
	mongoSession.Init(config)
	return mongoSession
}
func (mgs *MongoHeader) Init(config *database.DatabaseConfig) *MongoHeader {
	mgs.config = config
	return mgs
}
func (mgs *MongoHeader) Connect() error {
	config := mgs.config
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
	mgs.client = client
	error = mgs.client.Connect(context.Background())
	if error != nil {
		return error
	}
	mgs.database = mgs.client.Database(config.DataBaseName)
	if mgs.database == nil {
		return errors.New("连接到"+config.DataBaseName+"失败")
	}
	return nil
}

func (mgs *MongoHeader) collectionNameForObject(object interface{}) *mongo.Collection {
	targetType := reflect.TypeOf(object)
	modelName := targetType.Elem().Name()
	collectionName := strings.ToLower(modelName)
	return mgs.database.Collection(collectionName)
}

func (mgs *MongoHeader) makeContext() context.Context {
	content, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return content
}

func (mgs *MongoSession) Insert(object interface{}) (interface{}, error) {
	collection := mgs.collectionNameForObject(object)
	ctx := mgs.makeContext()
	res, err := collection.InsertOne(ctx, object)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (mgs *MongoSession) Delete(object interface{}) error {
	collection := mgs.collectionNameForObject(object)
	ctx := mgs.makeContext()
	data, err := mgs.structToMap(object)
	if err != nil {
		return err
	}
	_, err = collection.DeleteMany(ctx, data)
	return err
}

func (mgs *MongoSession) Update(object interface{}) error {
	collection := mgs.collectionNameForObject(object)
	ctx := mgs.makeContext()
	data, err := mgs.structToMap(object)
	if err != nil {
		return err
	}
	id := data["_id"]
	delete(data, "_id")
	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": data})
	if err != nil {
		return err
	}
	return nil
}

func (mgs *MongoSession) Find(object interface{}) (interface{}, error) {
	collection := mgs.collectionNameForObject(object)
	ctx := mgs.makeContext()
	data, err := mgs.structToMap(object)
	if err != nil {
		return nil, err
	}
	if id, ok := data["_id"]; ok {
		// 根据ID查找
		res := collection.FindOne(ctx, bson.M{"_id": id})
		output := reflect.New(reflect.TypeOf(object).Elem())
		err = res.Decode(output.Interface())
		return output.Interface(), err
	} else {
		// 查找多个
		res, err := collection.Find(ctx, data)
		if err != nil {
			return nil, err
		}
		s := make([]interface{}, 0)
		for res.Next(ctx) {
			if res.Err() != nil {
				return nil, res.Err()
			}
			output := reflect.New(reflect.TypeOf(object).Elem())
			err = res.Decode(output.Interface())
			if err != nil {
				return nil, err
			}
			s = append(s, output.Interface())
		}
		return s, nil
	}
}

func (mgs *MongoHeader) structToMap(object interface{}) (map[string]interface{}, error) {
	output := make(map[string]interface{}, 0)
	bytes, err := bson.Marshal(object)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(bytes, &output)
	if err != nil {
		return nil, err
	}
	if id, ok := output["_id"]; ok {
		if idString, ok := id.(string); ok {
			objectID, error := primitive.ObjectIDFromHex(idString)
			if error == nil {
				output["_id"] = objectID
			} else {
				return nil, err
			}
		}
	}
	return output, nil
}