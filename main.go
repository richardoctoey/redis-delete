package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

func DeleteFlushDB(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	formatstring := "xxxx:%d"

	for i := 0; i < 100000; i++ {
		err := rdb.Set(fmt.Sprintf(formatstring, i), "longkeykeykeykeykeykeykeykeykeykey", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	start := time.Now()
	rdb.FlushDB()
	elapsed := time.Since(start)
	fmt.Printf("Flushdb took %s\n", elapsed)
}

func DeleteByPattern(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	formatstring := "xxxx:%d"

	for i := 0; i < 100000; i++ {
		err := rdb.Set(fmt.Sprintf(formatstring, i), "longkeykeykeykeykeykeykeykeykeykey", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	startTotal := time.Now()
	iter := rdb.Scan(0, "xxxx:", 0).Iterator()
	var vals = []string{}
	for iter.Next() {
		vals = append(vals, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	startDelete := time.Now()
	rdb.Del(vals...)
	elapsedDelete := time.Since(startDelete)
	fmt.Printf("Delete took %s\n", elapsedDelete)
	elapsed := time.Since(startTotal)
	fmt.Printf("ScanDelete took %s\n", elapsed)

}

func main() {
	wg := &sync.WaitGroup{}
	DeleteFlushDB(wg)
	DeleteByPattern(wg)
	wg.Wait()
	fmt.Println("Done")
}
