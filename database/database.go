package database

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/go-gin-example-mongodb/config"
	"github.com/go-gin-example-mongodb/utils"
)

var conn *mongo.Database

func ConnectDB(config *config.Config) *mongo.Database {
	dbHost := config.DatabaseHost
	dbPort := config.DatabasePort
	dbUser := config.DatabaseUser
	dbPass := config.DatabasePassword
	dbName := config.DatabaseName
	dbAuth := config.DatabaseAuth

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var connection string
	var err error
	var client *mongo.Client
	if dbHost != "" || dbPort != "" || dbUser != "" || dbPass != "" {
		credential := options.Credential{
			AuthSource: dbAuth,
			Username:   dbUser,
			Password:   dbPass,
		}
		logrus.Info(utils.DatabaseConfigSet)
		connection = fmt.Sprintf("mongodb://%s:%s/?connect=direct", dbHost, dbPort)
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(connection).SetAuth(credential))
		if err != nil {
			logrus.Errorf(utils.FailedOpenDb, err)
		}
	} else {
		logrus.Errorf(utils.DatabaseConfigNotSet)
		connection = fmt.Sprintf("mongodb://%s:%s/?connect=direct", "localhost", "27017")
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(connection))
		if err != nil {
			logrus.Errorf(utils.FailedOpenDb, err)
		}
	}
	var db *mongo.Database
	if dbName != "" {
		db = client.Database(dbName)
	} else {
		logrus.Errorf(utils.DatabaseNameIsNotSet)
		db = client.Database("loan_process")
	}
	SetConnection(db)
	return db
}

//GetConnection : Get Available Connection
func GetConnection() *mongo.Database {
	return conn
}

//SetConnection : Set Available Connection
func SetConnection(connection *mongo.Database) {
	conn = connection
}
