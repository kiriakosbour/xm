package db

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"xm/domain"
)

type CompanyCrudRepo struct {
}
type StorageService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func CompanyCrudRepoInit() *CompanyCrudRepo {
	return &CompanyCrudRepo{}
}
func (r *CompanyCrudRepo) initializeRedis() *StorageService {
	var ctx = context.Background()
	storageService := &StorageService{}
	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	c := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	if err := c.Ping(ctx).Err(); err != nil {
		log.Println("Unable to connect to redis " + err.Error())
		return nil
	}
	storageService.redisClient = c
	storageService.ctx = ctx
	return storageService
}

//GetKey get value from redis
func (r *CompanyCrudRepo) GetKey(id string) (domain.Company, error) {
	company := domain.Company{}
	client := r.initializeRedis().redisClient
	val, err := client.Get(r.initializeRedis().ctx, id).Result()
	log.Println(val)
	if err == redis.Nil || err != nil {
		log.Printf("Error on get key %s", err.Error())
		return company, err
	}
	err = json.Unmarshal([]byte(val), &company)
	if err != nil {
		log.Printf("Error on unmarshalling %s", err.Error())
		return company, err
	}

	return company, nil
}

//SetKey save or upd value on redis
func (r *CompanyCrudRepo) SetKey(value domain.Company, id string, expiration time.Duration) error {
	log.Println(value)
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		log.Printf("Error on marshalling %s", err.Error())
		return err
	}
	client := r.initializeRedis().redisClient
	err = client.Set(r.initializeRedis().ctx, id, cacheEntry, expiration).Err()
	log.Println(id)
	if err != nil {
		log.Printf("Error on initialization %s", err.Error())
		return err
	}
	return nil
}

//DelKey del key
func (r *CompanyCrudRepo) DelKey(key string) error {
	c := r.initializeRedis().redisClient
	err := c.Del(r.initializeRedis().ctx, key).Err()
	if err != nil {
		log.Printf("Error on initialization %s", err.Error())
		return err
	}
	return nil
}

//GetAllValues retrieves all values from redis
func (r *CompanyCrudRepo) GetAllValues() map[string]string {
	client := r.initializeRedis().redisClient
	iterKeys := client.Scan(r.initializeRedis().ctx, 0, "*", 0).Iterator()
	allValues := make(map[string]string)
	for iterKeys.Next(r.initializeRedis().ctx) {
		val, _ := client.Get(r.initializeRedis().ctx, iterKeys.Val()).Result()
		allValues[iterKeys.Val()] = val
	}

	return allValues
}
