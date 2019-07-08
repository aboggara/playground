package mongo

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pronto-hubble/api/v1/models"
	"time"
)

const usersCollection = "users"

type UserRepo struct {
	client *MongoClient
}

func (r *UserRepo) Store(ctx context.Context, user *models.User) (string, error) {
	id := randomHex(4)
	result, err := r.client.InsertOne(ctx, usersCollection, bson.M{"id": id, "name": user.Name, "role": user.Role,
		"created" : time.Now().Unix()})

	if result != nil {
		return id, nil
	}
	return "", err
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func (r *UserRepo) FindOne(ctx context.Context, id string) (*models.User, error) {
	var result = models.User{}
	filter := bson.M{"id": id}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err := r.client.db.Collection(usersCollection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepo) FindAll(ctx context.Context) ([]*models.User, error) {

	cur, err := r.client.db.Collection(usersCollection).Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	} //TODO: propagate error message
	defer cur.Close(ctx)

	var users []*models.User
	for cur.Next(ctx) {
		var result = models.User{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		users = append(users, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) Update(ctx context.Context, id string) error {
	return nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	return nil
}
