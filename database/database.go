package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"ToDoList/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func Connect() error {
    var err error
    var connectionString string

    ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancelFn()

    // Load credentials
    dbHost := config.Config("DB_HOST")
    dbPort := config.Config("DB_PORT")
    dbName := config.Config("DB_NAME")
    collName := config.Config("COLL_NAME")
    dbUser := config.Config("DB_USER")
    dbPass := config.Config("DB_PASS")
    isSrv := config.Config("SRV")

    if (isSrv == "yes") {
        connectionString = fmt.Sprintf("mongodb+srv://%s:%s@%s",dbUser,dbPass,dbHost)
    } else {
        connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s",dbUser,dbPass,dbHost,dbPort)
    }

    log.Println("Connection String: " + connectionString)

    client, err:= mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

    if err != nil {
        return err
    }
    defer client.Disconnect(ctx)

    db := client.Database(dbName)
    collection = db.Collection(collName)

    log.Println("Opened database connection and loaded collection")
    
    return nil
}