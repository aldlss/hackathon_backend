package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	Driver   neo4j.DriverWithContext
	UserPgDb *gorm.DB
)

func initNeo4j() {
	var err error
	uri := os.Getenv("NEO4J_URL")
	auth := neo4j.BasicAuth(os.Getenv("NEO4J_USERNAME"), os.Getenv("NEO4J_PASSWORD"), "")
	Driver, err = neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		log.Fatal(err.Error())
	}
}

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
}

func Init() {
	//initNeo4j()
	initPgsql()
}
