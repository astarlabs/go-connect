package data

import (
	"fmt"
	"strconv"

	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQLConnection ...
type MySQLConnection struct {
	host string
	port int
	user string
	pass string
	name string
}

// NewMySQLConnection - MySQLConnection construtor
func NewMySQLConnection() *MySQLConnection {

	fileReader, err := ini.InsensitiveLoad("./database.conf")
	if err != nil {
		fmt.Println("database.conf not found")
		return nil
	}

	section := fileReader.Section("DATABASE")

	host := section.Key("database.host").Value()
	port := section.Key("database.port").Value()
	user := section.Key("database.user").Value()
	pass := section.Key("database.pass").Value()
	name := section.Key("database.name").Value()

	mysqlConn := new(MySQLConnection)

	mysqlConn.host = host
	mysqlConn.port, _ = strconv.Atoi(port)
	mysqlConn.user = user
	mysqlConn.pass = pass
	mysqlConn.name = name

	return mysqlConn

}

// Connect - connect
func (connection *MySQLConnection) Connect() (*gorm.DB, error) {

	port := strconv.Itoa(connection.port)

	connectionURL := connection.user + ":" + connection.pass + "@tcp(" + connection.host + ":" + port + ")/" + connection.name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
