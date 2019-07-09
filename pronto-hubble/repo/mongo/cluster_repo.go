package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pronto-hubble/api/v1/models"
	"time"
)

const clustersCollection = "clusters"

type ClusterRepo struct {
	client *MongoClient
}

func (r *ClusterRepo) Store(ctx context.Context, cluster *models.Cluster) (string, error) {
	id := RandomHex(4)
	deviceIds := make([]string, 0)
	if cluster.DeviceIds != nil {
		deviceIds = cluster.DeviceIds
	}
	result, err := r.client.InsertOne(ctx, clustersCollection,
		bson.M{"id": id, "name": cluster.Name, "deviceIds": deviceIds, "created": time.Now().Unix()})

	if result != nil {
		return id, nil
	}
	return "", err
}

func (r *ClusterRepo) FindOne(ctx context.Context, id string) (*models.Cluster, error) {
	var result = models.Cluster{}
	filter := bson.M{"id": id}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err := r.client.db.Collection(clustersCollection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ClusterRepo) FindAll(ctx context.Context) ([]*models.Cluster, error) {

	cur, err := r.client.db.Collection(clustersCollection).Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	} //TODO: propagate error message
	defer cur.Close(ctx)

	var clusters []*models.Cluster
	for cur.Next(ctx) {
		var result = models.Cluster{}
		var deviceIds []string
		if err := cur.Current.Lookup("deviceIds").Unmarshal(&deviceIds); err != nil {
			panic(err)
		}
		err := cur.Decode(&result)
		if deviceIds != nil { //TODO: somehow decode does not unmarshal the array
			result.DeviceIds = deviceIds
		}
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return clusters, nil
}

func (r *ClusterRepo) Update(ctx context.Context, id string) error {
	return nil
}

func (r *ClusterRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *ClusterRepo) AddDevices(ctx context.Context, id string, deviceIds []string) error {
	filter := bson.D{{"id", id}}
	update := bson.D{
		{"$addToSet", bson.D{
			{"deviceIds", bson.D{{ "$each" , deviceIds}}},
		}},
	}

	_, err := r.client.db.Collection(clustersCollection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClusterRepo) RemoveDevices(ctx context.Context, id string, deviceIds []string) error {
	filter := bson.D{{"id", id}}
	update := bson.D{
		{"$pull", bson.D{
			{"deviceIds", bson.D{{ "$in", deviceIds }}},
		}},
	}

	_, err := r.client.db.Collection(clustersCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}