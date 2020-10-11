package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersDB       = "mysql_users_db"
)

//SQLClient is the instance of client
var (
	SQLClient SQLClientInterface = &sqlClient{}
	username                     = os.Getenv(mysqlUsersUsername)
	password                     = os.Getenv(mysqlUsersPassword)
	host                         = os.Getenv(mysqlUsersHost)
	db                           = os.Getenv(mysqlUsersDB)
)

type sqlClient struct {
	Client *sql.DB
}

//SQLClientInterface is the interface for Client
type SQLClientInterface interface {
	Init()
	GetClient() *sql.DB
}

func (c *sqlClient) Init() {
	dataSourceConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, db)
	var err error
	c.Client, err = sql.Open("mysql", dataSourceConfig)
	if err != nil {
		panic(err)
	}
	if err = c.Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to Database successfully")
}

func (c *sqlClient) GetClient() *sql.DB {
	return c.Client
}
