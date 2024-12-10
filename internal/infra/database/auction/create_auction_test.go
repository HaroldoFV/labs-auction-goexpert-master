package auction

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionClosingRoutine(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/auctions?authSource=admin"))
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.TODO())

	database := client.Database("auctions_test")
	collection := database.Collection("auctions")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.Drop(ctx)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewAuctionRepository(database)

	auction, _ := auction_entity.CreateAuction("Test Product", "Category", "Description", auction_entity.New)
	_ = repo.CreateAuction(context.TODO(), auction)

	interval, _ := time.ParseDuration("24s")
	log.Printf("Sleeping for %v", interval+time.Second)
	time.Sleep(interval + time.Second)
	log.Println("Woke up from sleep")

	var result auction_entity.Auction
	err = collection.FindOne(context.TODO(), bson.M{"_id": auction.Id}).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, auction_entity.Completed, result.Status)
}
