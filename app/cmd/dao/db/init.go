package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"time"
)

var (
	UserPgDb    *gorm.DB
	ContentPgDb *gorm.DB
	random      *rand.Rand
)

func initPgsql() {
	dsn := fmt.Sprintf("host=%s port='%s' user=%s password=%s dbname=%s TimeZone=Asia/Shanghai connect_timeout=10",
		os.Getenv("PGSQL_HOST"), os.Getenv("PGSQL_PORT"),
		os.Getenv("PGSQL_USER"), os.Getenv("PGSQL_PASSWORD"), os.Getenv("PGSQL_DBNAME"))
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	tableName := os.Getenv("USER_TABLE_NAME")
	if tableName == "" {
		log.Fatal("USER_TABLE_NAME can't be empty")
	}
	UserPgDb = db.Table(tableName)

	tableName = os.Getenv("CONTENT_TABLE_NAME")
	if tableName == "" {
		log.Fatal("CONTENT_TABLE_NAME can't be empty")
	}
	ContentPgDb = db.Table(tableName)
}

func Init() {
	initPgsql()
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}
