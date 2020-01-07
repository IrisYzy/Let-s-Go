package main

import (
	"fmt"
	"github.com/oklog/ulid"
	"math/rand"
	"sync"
	"time"
)

var synWait sync.WaitGroup

func main() {
	start := time.Now()
	for i := 1; i <= 10000; i++ {
		synWait.Add(1)
		data := GenUlid()
		fmt.Println(data)
	}
	synWait.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())

}

/**
 * 生成uuid
 * @return 	uuid
 */
func GenUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	//fmt.Println(id)
	synWait.Done()
	return id.String()
}
