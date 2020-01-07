package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() () {
	initRedis()
}

var redisCoon redis.Conn

//const (
//	SCRIPT_IN = `
//local score = redis.call('ZSCORE',KEYS[1],ARGV[1] )
//if score = nil
//then
//	return 0
//else if score < ARGV[2]
//then
//	return 1
//end`
//)
const (
	SCRIPT_IN = `
local score = redis.call('zcard', 'yzy')
if score == nil
then
	return 0
else 
	return score
end
local code = redis.call('ZREVRANGEBYSCORE','yzy', ARGV[2], ARGV[3])
if (code == nil)
then
	return -1
else 
	return 1
end
`
)

//初始化redis连接，redis在本地docker中运行
func initRedis() () {
	redisClient := &redis.Pool{
		// 最大空闲链接
		MaxIdle: 10,
		// 最大激活链接
		MaxActive: 10,
		// 最大空闲链接等待时间
		IdleTimeout: 5 * time.Second,
		Dial: func() (redis.Conn, error) {
			rc, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			rc.Do("SELECT", 0)
			//fmt.Println("USE DB", 0)

			return rc, nil
		},
	}
	redisCoon = redisClient.Get()
}

//生成六位随机验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

func SendSmsCode(phoneNum, code string) () {
	//phoneNum := "13521400524"

	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	timeFiveminStr := strconv.FormatInt(time.Now().Add(-time.Minute * 5).Unix(),10)
	//fmt.Println(timeStr,timeFiveminStr)
	lua := redis.NewScript(1, SCRIPT_IN)
	//in, err := redis.Int(lua.Do(redisCoon, phoneNum, code, timeStr,timeFiveminStr))
	in, err := lua.Do(redisCoon, phoneNum, code, timeStr,timeFiveminStr)
	fmt.Println(in, "--", err)
}

func main() {
	code := GenValidateCode(6)
	//fmt.Println("验证码：", code)
	SendSmsCode("13521400524", code)
}
