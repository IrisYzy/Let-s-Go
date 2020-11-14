package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTIme   int64 `bson:"endTime"`
}

// 存储在mongodb中的内容
type RecordLog struct {
	JobName   string    `bson:"jobName"`
	Command   string    `bson:"command"`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	Timepoint TimePoint `bson:"timepoint"`
}

type LogRecord struct {
	JobName string `bson:"jobName"`
}

func InsertRecord(client *mongo.Client, collect *mongo.Collection) (insertID primitive.ObjectID) {

	collect = client.Database("corn").Collection("jobs")
	record := &RecordLog{
		JobName: "job1",
		Command: "main.go",
		Err:     "",
		Content: "Hello_World",
		Timepoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTIme:   time.Now().Unix() + 10,
		},
	}
	insertRest, err := collect.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return
	}
	insertID = insertRest.InsertedID.(primitive.ObjectID)
	return insertID
}

func FindLog(collect *mongo.Collection) {
	// 创建需要过滤的条件
	logred := &LogRecord{
		JobName: "job1",
	}
	var skip int64 = 0  //从那个开始
	var limit int64 = 2 //炼制几个输出字段
	cursor, err := collect.Find(context.TODO(), logred, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		//创建需要反序列化成什么样子的结构体对象
		records := &RecordLog{}
		//反序列化
		err = cursor.Decode(records)
		if err != nil {
			fmt.Println(err)
			return
		}
		//打印
		fmt.Println(*records)
	}
}

func main() {
	//创建上下文
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(10*time.Second, func() {
		cancelFunc()
	})
	//常见数据库连接
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.21.138.18:27017").SetMaxPoolSize(20))
	if err != nil {
		fmt.Println(err)
		return
	}
	collect := client.Database("chat_message").Collection("message_whisper_send")
	//插入数据
	//InsertRecord(client, collect)
	//查询数据
	//var result WhisperRecord
	//_ = collect.FindOne(context.TODO(), bson.M{"from.name": "yzy"}).Decode(&result)
	//fmt.Println(result)

	countResult, err := collect.CountDocuments(context.TODO(), bson.M{"from_number": "791617018"})
	fmt.Println(countResult)

}

//type WhisperRecord struct {
//	ClassID     string  `json:"class_id,omitempty" bson:"class_id"`
//	Content     string  `json:"content,omitempty" bson:"content"`
//	Data        Data    `bson:"data,omitempty" `
//	From        From    `bson:"from,omitempty"`
//	MessageType string  `bson:"message_type,omitempty"`
//	ResInfo     ResInfo `bson:"res_info,omitempty"`
//	Time        int64   `bson:"time,omitempty" `
//	To          To      `bson:"to,omitempty" `
//	UUID        string  `bson:"wspid,omitempty"`
//	WhisperType int     `bson:"whisper_type,omitempty"`
//	//Type        int     `json:"type"`
//	//UserList    userList `json:"user_list,omitempty"`
//}
//
//type Data struct {
//	Height int    `bson:"height"`
//	Type   int    `bson:"type"`
//	Url    string `bson:"url"`
//	Width  int    `bson:"width"`
//}

//type From struct {
//	Avatar  string `bson:"avatar,omitempty"`
//	EndType int    `bson:"end_type,omitempty"`
//	Group   int    `bson:"group,omitempty"`
//	ID      string `bson:"id,omitempty"`
//	Name    string `bson:"name,omitempty"`
//	Number  string `bson:"number,omitempty"`
//	Type    int    `bson:"type,omitempty"`
//}
//
//type ResInfo struct {
//	SuccessCnt int `bson:"success_cnt"`
//}
//
//type To struct {
//	Avatar  string `bson:"avatar"`
//	EndType int    `bson:"end_type"`
//	Group   int    `bson:"group"`
//	ID      string `bson:"id"`
//	Name    string `bson:"name"`
//	Number  string `bson:"number"`
//	Type    int    `bson:"type"`
//}

