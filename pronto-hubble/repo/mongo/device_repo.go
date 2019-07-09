package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pronto-hubble/api/v1/models"
	"time"
)

const devicesCollection = "devices"

type DeviceRepo struct {
	client *MongoClient
}

func (r *DeviceRepo) Store(ctx context.Context, device *models.Device) (string, error) {
	id := RandomHex(4)
	result, err := r.client.InsertOne(ctx, devicesCollection,
		bson.M{"id": id, "name": device.Name, "metadata": device.MetaData, "created": time.Now().Unix()})

	if result != nil {
		return id, nil
	}
	return "", err
}

func (r *DeviceRepo) FindOne(ctx context.Context, id string) (*models.Device, error) {
	var result = models.Device{}
	filter := bson.M{"id": id}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err := r.client.db.Collection(devicesCollection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *DeviceRepo) FindAll(ctx context.Context) ([]*models.Device, error) {

	cur, err := r.client.db.Collection(devicesCollection).Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	} //TODO: propagate error message
	defer cur.Close(ctx)

	var devices []*models.Device
	for cur.Next(ctx) {
		var result = models.Device{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		devices = append(devices, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *DeviceRepo) Update(ctx context.Context, id string) error {
	return nil
}

func (r *DeviceRepo) Delete(ctx context.Context, id string) error {
	return nil
}
