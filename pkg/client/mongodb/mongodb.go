package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDb string) (db *mongo.Database, err error) {
	var mongodbURL string
	var isAuth bool
	if username == "" && password == "" {
		mongodbURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongodbURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clientOprtions := options.Client().ApplyURI(mongodbURL)
	if isAuth {
		if authDb == "" {
			authDb = database
		}
		clientOprtions.SetAuth(options.Credential{
			AuthSource: authDb,
			Username: username,
			Password: password,
		})
	}

	// Connect
	client, err := mongo.Connect(ctx, clientOprtions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	// Ping
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}

	return client.Database(database), nil
}
