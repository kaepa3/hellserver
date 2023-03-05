package testmod

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongo(m *testing.M) (error, func(), string) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Printf("Could not construct pool: %s\n", err)
		return err, nil, ""
	}
	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Printf("Could not connect to Docker: %s", err)
		return err, nil, ""
	}
	runOptions := &dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=user",
			"MONGO_INITDB_ROOT_PASSWORD=password",
		},
	}

	resource, err := pool.RunWithOptions(runOptions,
		func(hc *docker.HostConfig) {
			hc.AutoRemove = true
			hc.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)

	port := resource.GetPort("27017/tcp")
	log.Println(port)
	uri := fmt.Sprintf("mongodb://user:password@localhost:%s", port)
	ctx := context.TODO()
	// 起動するまで待機
	var dbClient *mongo.Client
	pool.Retry(func() error {
		dbClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		defer dbClient.Disconnect(ctx)
		if err != nil {
			return err
		}
		return dbClient.Ping(ctx, nil)
	})
	if err != nil {
		log.Printf("Could not connect to docker: %s", err)
		return err, nil, ""
	}
	fmt.Println("start mongo container🐳")
	// mongoクライアントをクローズ
	dbClient.Disconnect(ctx)
	return nil, func() { pool.Purge(resource) }, uri
}
