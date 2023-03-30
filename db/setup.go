package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"linked-page/config"
	"log"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		config.Env.POSTGRES_USER,
		config.Env.POSTGRES_PASSWORD,
		config.Env.POSTGRES_DB,
		config.Env.POSTGRES_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PgSql")
	return db
}

// DB Client instance
var DB = ConnectDB()

// RedisClient Start Redis on database 1 - it's used to store the JWT, but you can use it for anything else
// Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
// RedisClient ...
// var RedisClient *_redis.Client = InitRedis(1)

// func InitRedis(selectDB ...int) *_redis.Client {

// 	var redisHost = os.Getenv("REDIS_HOST")
// 	//var redisPassword = os.Getenv("REDIS_PASSWORD")

// 	_RedisClient := _redis.NewClient(&_redis.Options{
// 		Addr: redisHost,
// 		//Password: redisPassword,
// 		DB: selectDB[0],
// 		// DialTimeout:        10 * time.Second,
// 		// ReadTimeout:        30 * time.Second,
// 		// WriteTimeout:       30 * time.Second,
// 		// PoolSize:           10,
// 		// PoolTimeout:        30 * time.Second,
// 		// IdleTimeout:        500 * time.Millisecond,
// 		// IdleCheckFrequency: 500 * time.Millisecond,
// 		// TLSConfig: &tls.Config{
// 		// 	InsecureSkipVerify: true,
// 		// },
// 	})
// 	_, err := _RedisClient.Ping().Result()
// 	if err != nil {
// 		fmt.Println("redis connection error ", err)
// 	}
// 	//print("redis test: ", sta, "\n")
// 	return _RedisClient
// }

// // GetRedis ...
// func GetRedis() *_redis.Client {
// 	return RedisClient
// }
