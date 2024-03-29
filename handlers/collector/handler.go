package collector

import (
	context "context"
	fmt "fmt"
	os "os"
	time "time"

	discordgo "github.com/bwmarrin/discordgo"

	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

type Collector struct {
	client *mongo.Client
}

type Presence struct {
	datetime time.Time
	data     interface{}
}

func New() (*Collector, error) {
	fmt.Println("New Collector handler")
	var err error
	instance := Collector{}
	instance.client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = instance.client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func (instance Collector) Handler(s *discordgo.Session, event *discordgo.PresenceUpdate) {
	presence := instance.client.Database("fafeliscord").Collection("presence")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := presence.InsertOne(ctx, bson.M{
		"datetime": time.Now(),
		"data":     event,
	})
	if err != nil {
		fmt.Println("Error inserting presence", err)
	}
}
