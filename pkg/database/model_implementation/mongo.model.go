package modelimplementation

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoModel struct {
	BaseModel
	Collection *mongo.Collection
}

func NewMongoModel(client *mongo.Client, dbName, collectionName string) *MongoModel {
	return &MongoModel{
		BaseModel: BaseModel{
			Client: client,
			Name:   collectionName,
		},
		Collection: client.Database(dbName).Collection(collectionName),
	}
}

func (m *MongoModel) FindById(id string) (interface{}, error) {
	var result interface{}
	err := m.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	return result, err
}

func (m *MongoModel) FindOne(criteria interface{}) (interface{}, error) {
	var result interface{}
	err := m.Collection.FindOne(context.Background(), criteria).Decode(&result)
	return result, err
}

func (m *MongoModel) FindAll(criteria interface{}) ([]interface{}, error) {
	cursor, err := m.Collection.Find(context.Background(), criteria)
	if err != nil {
		return nil, err
	}
	var results []interface{}
	err = cursor.All(context.Background(), &results)
	return results, err
}

func (m *MongoModel) UpdateOne(criteria interface{}, update interface{}) error {
	_, err := m.Collection.UpdateOne(context.Background(), criteria, bson.M{"$set": update})
	return err
}

func (m *MongoModel) UpdateById(id string, update interface{}) error {
	_, err := m.Collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (m *MongoModel) UpdateAll(criteria interface{}, update interface{}) error {
	_, err := m.Collection.UpdateMany(context.Background(), criteria, bson.M{"$set": update})
	return err
}

func (m *MongoModel) DeleteById(id string) error {
	_, err := m.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (m *MongoModel) DeleteOne(criteria interface{}) error {
	_, err := m.Collection.DeleteOne(context.Background(), criteria)
	return err
}

func (m *MongoModel) DeleteAll(criteria interface{}) error {
	_, err := m.Collection.DeleteMany(context.Background(), criteria)
	return err
}

func (m *MongoModel) Create(data interface{}) error {
	_, err := m.Collection.InsertOne(context.Background(), data)
	return err
}
