package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	healthColName = "health"
)

type HealthCollection struct {
	col *mongo.Collection
}

// 健康情報の取得
func GetHealthCollection(ctx context.Context) (*HealthCollection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	col := client.Database(dbName).Collection(healthColName)

	return &HealthCollection{col: col}, err
}
