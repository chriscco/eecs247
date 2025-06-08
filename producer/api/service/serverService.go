package service

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"
	"wordCountServer/api/interface/client"
	"wordCountServer/common/utils"
	"wordCountServer/global"
)

// ServerImpl
type ServerImpl struct {} 

// NewServerImpl 
//	@return ServerImpl 
func NewServerImpl() ServerImpl {
	return ServerImpl{}  
}

const (
	MAX_MESSAGE_LEN = 2e6 
)

// mutex 
var mutex sync.Mutex  

// Count 
//	@param c 
//	@param text 
//	@return int 
func (ss *ServerImpl) Count(c context.Context, text string) (ans int) {
	var wg sync.WaitGroup

	log.Println("parsing input text...")
	sentences := splitByLength(text, MAX_MESSAGE_LEN) 
	for _, st := range sentences {
		key := utils.KeyGen(st) 
		if checkRedis(key) {
			log.Println("found in redis")
			return getRedisVal(key)
		}

		wg.Add(1)
		go func(key string) {
			log.Println("asynchronously checking redis...")
			for {
				if checkRedis(key) {
					mutex.Lock()
					ans += getRedisVal(key)
					mutex.Unlock()
					wg.Done()
					break 
				}
				time.Sleep(time.Second)
			}
		}(key)
		log.Printf("sending sentence to channel\n")
		client.SenderChannel <- st 
	}
	wg.Wait() 

	return 
}

// getRedisVal 
//	@param key 
//	@return int 
func getRedisVal(key string) (cntAll int) {
	val, err := global.Redis.HGetAll(key).Result() 
	if err != nil {
		log.Fatal("unable to get key: ", err) 
		return 0 
	}
	for _, cnt := range val {
		cntInt, _ := strconv.Atoi(cnt) 
		cntAll += cntInt 
	}
	return
}

// checkRedis 
//	@param key 
//	@return bool 
func checkRedis(key string) bool {
	exist, err := global.Redis.Exists(key).Result() 
	if err != nil {
		log.Fatal("unable to check redis: ", err) 
		return false 
	}
	if exist > 0 {
		log.Printf("key exists, key: %s\n", key) 
		return true 
	}
	return false 
}

// splitByLength 
//	@param s 
//	@param length 
//	@return []string 
func splitByLength(s string, length int) []string {
    var result []string
    for i := 0; i < len(s); i += length {
        end := min(i + length, len(s))
		for end < len(s) && s[end] != ' ' {
			end += 1 
		}
        result = append(result, s[i : end])
    }
    return result
}