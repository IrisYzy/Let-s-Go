package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Stu struct {
	Name string
	Age  int
}

type WhisperRecord struct {
	ClassID      string `json:"class_id,omitempty" bson:"class_id"`
	MessageType  string `bson:"message_type,omitempty"`
	Content      string `json:"content,omitempty" bson:"content"`
	Data         interface{}   `json:"data,omitempty" bson:"data,omitempty" `
	Time         int64  `bson:"time,omitempty" `
	WspId        string `bson:"wspid,omitempty"`
	SessionId    string `bson:"sessionid,omitempty"`
	WhisperType  int    `bson:"whisper_type,omitempty"`
	CourseNumber string `bson:"course_number,omitempty"`
	FromNumber   string `json:"from_number" bson:"from_number,omitempty"`
	ToNumber     string `json:"to_number" bson:"to_number,omitempty"`
}

type Data struct {
	Height int    `bson:"height,omitempty"`
	Type   int    `bson:"type,omitempty"`
	Url    string `bson:"url,omitempty"`
	Width  int    `bson:"width,omitempty"`
}

type From struct {
	Avatar  string `bson:"avatar,omitempty"`
	EndType int    `bson:"end_type,omitempty"`
	Group   int    `bson:"group,omitempty"`
	ID      string `bson:"id,omitempty"`
	Name    string `bson:"name,omitempty"`
	Number  string `bson:"number,omitempty"`
	Type    int    `bson:"type,omitempty"`
}

type ResInfo struct {
	SuccessCnt int `bson:"success_cnt"`
}

type To struct {
	Avatar  string `bson:"avatar"`
	EndType int    `bson:"end_type"`
	Group   int    `bson:"group"`
	ID      string `bson:"id"`
	Name    string `bson:"name"`
	Number  string `bson:"number"`
	Type    int    `bson:"type"`
}

type MessageRelation struct {
	LastId     string `bson:"last_id"`
	ClassId    string `bson:"class_id"`
	Time       int64  `bson:"time"`
	SessionId  string `bson:"sessionid"`
	FromNumber string `bson:"from_number"`
	ToNumber   string `bson:"to_number"`
}

type UserInfo struct {
	ClassId     int64  `bson:"class_id"`
	MessageType string `bson:"message_type"`
	Id          string `bson:"id"`
	Number      string `bson:"number"`
	Type        int64  `bson:"type"`
	Name        string `bson:"name"`
	Avatar      string `bson:"avatar"`
}

func main() {
	db, _ := InitMongoDB()
	var result []WhisperRecord
	_ = db.C("message_whisper_send").Find(bson.M{"from_number": "668669678"}).All(&result)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))

	//_, toNum := "625161488", "123"
	//insertRecord(fromNum, toNum, db)
	//mongo := db.C("message_relation_store")
	//insertRelation := MessageRelation{
	//	LastId:     "925wspId",
	//	Time:       1604650600,
	//	SessionId:  "625161481_625161488",
	//	FromNumber: fromNum,
	//	ToNumber:   toNum,
	//}
	//Insert(mongo, insertRelation)

	//userInfo := UserInfo{
	//	Number: toNum,
	//	Name:   "userName1",
	//	Type:   1,
	//	Avatar: "http://baidu.com",
	//}
	//mongoUser := db.C("message_userinfo")
	//Insert(mongoUser,userInfo)

	//classIds := []string{"10"}
	//
	//query := bson.M{"course_number": bson.M{"$in": classIds}}
	////query["time"] = bson.M{"$gt": 1604388568}
	//query["sessionid"] = "625161486_625161488"
	//
	//err := mongo.Find(query).
	//	All(&result)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//
	//fmt.Println(result)
}

func insertRelation() {

}

func insertRecord(fromNum, toNum string, db *mgo.Database) {
	mongo := db.C("message_whisper_send")
	sessionId := genSessionId(fromNum, toNum)
	whisperRecord := WhisperRecord{
		ClassID:     "6010230990008288",
		MessageType: "message_whisper_send",
		Content:     "私聊",
		Data: Data{
			Url: "http://baidu.com",
		},
		Time:         time.Now().Unix(),
		WspId:        "925wspId",
		SessionId:    sessionId,
		WhisperType:  1,
		CourseNumber: "10",
		FromNumber:   fromNum,
		ToNumber:     toNum,
	}
	Insert(mongo, whisperRecord)
}

func genSessionId(fromNumber, toNumber string) string {
	sessionId := fromNumber + "_" + toNumber
	if fromNumber > toNumber {
		sessionId = toNumber + "_" + fromNumber
	}
	return sessionId
}

func insertUser() {

}

func Insert(mongo *mgo.Collection, InsertInfo interface{}) {

	err := mongo.Insert(InsertInfo)
	if err != nil {
		fmt.Println("ERROR_Insert:", err)
	}

}

func InitMongoDB() (*mgo.Database, *mgo.Session) {
	mgo_url := "mongodb://172.21.138.18:27017" //登录地址 第一步

	session, err := mgo.Dial(mgo_url)
	if err != nil || session == nil {
		fmt.Printf("MongoDB connect error %s", err)
	}

	session.SetMode(mgo.Monotonic, true)

	db := session.DB("chat_message") //数据库名称

	if err != nil {
		fmt.Println("------连接数据库失败------------")
		panic(err)
	}
	fmt.Println("------ConnectionDb-----2-------")
	return db, session
}
