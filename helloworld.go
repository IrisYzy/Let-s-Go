package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main()  {
	//请求参数
	requestParams := map[string]interface{}{
		"event":  "123",
		"params": "abc",
	}
	//构建request实例，发送请求
	jsonRequestParams, err := json.Marshal(requestParams)
	if err != nil {
		return
	}
	requestEntity, err := http.NewRequest("httpType", "http://baidu.com", bytes.NewBuffer(jsonRequestParams))
	fmt.Println(string(jsonRequestParams))
	fmt.Println(requestEntity)

}

