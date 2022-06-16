package initialize

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	info := global.Settings.MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+info.Host+":"+strconv.Itoa(info.Port)))
	if err != nil {
		color.Red("========================================================")
		color.Red("MongoDB connection is not correct, please check settings-dev.yaml file and rerun the server again")
		fmt.Println(err)
		color.Red("========================================================")
		os.Exit(3)
	}

	global.Mongo = client
}
