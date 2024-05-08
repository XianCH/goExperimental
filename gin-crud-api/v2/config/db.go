package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "password"
	dbName   = "test"
)

func DatabaseConnection() (*sql.DB, error) {
	// 构建数据库连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)

	// 连接数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// 测试连接是否成功
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
