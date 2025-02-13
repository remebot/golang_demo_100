package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	coll := client.Database("test").Collection("LkMongo20250213")

	filter := bson.D{{"ReqNonceStr", "c43d7b41-7a84-40b5-bf38-e19e9cc3bec2"}}

	var record LogRecord
	err = coll.FindOne(context.TODO(), filter).Decode(&record)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	// end findOne

	output, err := json.MarshalIndent(record, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}

type ReportInfo struct {
	TimeStart        string `json:"TimeStart"`
	TimeEnd          string `json:"TimeEnd"`
	DefaultTimeStart string `json:"DefaultTimeStart"`
	DefaultTimeEnd   string `json:"DefaultTimeEnd"`
	OpenTimeMemo     string `json:"OpenTimeMemo"`
}

type ResponseResult struct {
	IsSuccess string `json:"IsSuccess"`
}

type BodyJson struct {
	Body struct {
		ReportInfo     ReportInfo     `json:"ReportInfo"`
		ResponseResult ResponseResult `json:"ResponseResult"`
	} `json:"body"`
}

type InputXml struct {
	XmlContent string `json:"InputXml"`
}

type OutputXml struct {
	XmlContent string `json:"OutputXml"`
}

type LogRecord struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	BodyJson      BodyJson           `bson:"BodyJson" json:"BodyJson"`
	MethodName    string             `bson:"MethodName" json:"MethodName"`
	KeyWord       *string            `bson:"KeyWord" json:"KeyWord"`
	BizIdentifier string             `bson:"BizIdentifier" json:"BizIdentifier"`
	Environment   string             `bson:"Environment" json:"Environment"`
	ReqNonceStr   string             `bson:"ReqNonceStr" json:"ReqNonceStr"`
	Day           string             `bson:"Day" json:"Day"`
	Hours         int32              `bson:"Hours" json:"Hours"`
	Date          int32              `bson:"Date" json:"Date"`
	OpenUserID    string             `bson:"OpenUserID" json:"OpenUserID"`
	Platform      int32              `bson:"Platform" json:"Platform"`
	InputXml      InputXml           `bson:"InputXml" json:"InputXml"`
	OutputXml     OutputXml          `bson:"OutputXml" json:"OutputXml"`
	TimeStamp     int32              `bson:"timeStamp" json:"timeStamp"`
	Level         int32              `bson:"Level" json:"Level"`
	Message       string             `bson:"Message" json:"Message"`
	HMACSM3       *string            `bson:"HMACSM3" json:"HMACSM3"`
}
