package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type YmlConfig struct {
	Database Database
	Redis    Redis
}

var ymlConfig YmlConfig

func InitYml() (*YmlConfig, error) {
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("yaml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	err = viper.Unmarshal(&ymlConfig)
	if err != nil {
		return nil, fmt.Errorf("fatal error unmarshal file: %w", err)
	}

	fmt.Println(ymlConfig.Database.Host)

	return &ymlConfig, nil
}

func InitDatabase(cfg *Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func InitRedis(cfg *Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return rdb, nil
}

func TestViper() {
	config, err := InitYml()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}
	fmt.Println(config.Database.Host)

	// db, err := InitDatabase(&config.Database)
	// if err != nil {
	// 	log.Fatalf("Error initializing database: %v", err)
	// }
	// defer db.Close()

	// rdb, err := InitRedis(&config.Redis)
	// if err != nil {
	// 	log.Fatalf("Error initializing Redis: %v", err)
	// }
	// defer rdb.Close()

	fmt.Println("Successfully connected to database and Redis")
}
