package data

import (
	"fmt"
	"strconv"

	"github.com/go-ini/ini"
	"gopkg.in/mgo.v2"
)

// MongoConnection ...
type MongoConnection struct {
	host string
	port int
	user string
	pass string
	name string
}

// NewMongoConnection - constructor
func NewMongoConnection() *MongoConnection {

	fileReader, err := ini.InsensitiveLoad("database.conf")
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

	mongoConn := new(MongoConnection)

	mongoConn.host = host
	mongoConn.port, _ = strconv.Atoi(port)
	mongoConn.user = user
	mongoConn.pass = pass
	mongoConn.name = name

	return mongoConn

}

// Connect ...
func (connection *MongoConnection) Connect() (*mgo.Session, error) {

	connectionURL := connection.host + ":" + strconv.Itoa(connection.port)

	dialInfo := new(mgo.DialInfo)
	dialInfo.Username = connection.user
	dialInfo.Password = connection.pass
	dialInfo.Addrs = []string{connectionURL}
	dialInfo.Database = connection.name

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
