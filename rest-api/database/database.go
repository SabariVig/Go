package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<USER_NAME>:<SUPER_SECRET_PASSWORD_GOES_HERE>@<URI>"))

var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
