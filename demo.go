package main

import (
	"errors"
	"log"
	"singleflight/singleflight"
	"sync"
)

func main() {
	//testNoSingleFlight()
	testSingleFlight()
}

var errorNotExist = errors.New("not exist")

func testNoSingleFlight() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := getDataNoSingleFlight("key")
			if err != nil {
				log.Println("getDataNoSingleFlight 获取失败")
			}
			log.Printf("testNoSingleFlight 开启协程:%d,获取key:%s,对应值data:%s \n", i, "key", data)
		}()
	}
	wg.Wait()
}

func testSingleFlight() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := getDataSingleFlight("key")
			if err != nil {
				log.Println("getDataNoSingleFlight 获取失败")
			}
			log.Printf("testSingleFlight 开启协程:%d,获取key:%s,对应值data:%s \n", i, "key", data)
		}()
	}
	wg.Wait()
}

func getDataNoSingleFlight(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		data, err = getDataFromDB(key)
		if err != nil {
			log.Println("从DB获取数据失败")
			return "", err
		}
		// set cache
	} else if err != nil {
		return "", err
	}
	return data, nil
}

var g singleflight.Group

func getDataSingleFlight(key string) (string, error) {

	data, err := getDataFromCache(key)
	if err == errorNotExist {
		v, err := g.Do(key, func() (i interface{}, err error) {
			return getDataFromDB(key)
		})

		if err != nil {
			log.Println(err)
			return "", err
		}

		data = v.(string)
	} else if err != nil {
		return "", err
	}
	return data, nil
}

func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

func getDataFromDB(key string) (string, error) {
	log.Printf("get %s from database", key)
	return "data from db", nil
}
