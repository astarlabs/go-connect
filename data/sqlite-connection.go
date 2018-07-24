package data

import (
	"github.com/jinzhu/gorm"
)

// MySQLConnection ...
type SQLiteConnection struct {
	host string
}

func NewSQLiteConnection(connection string) *SQLiteConnection {

	sqliteConn := new(SQLiteConnection)
	sqliteConn.host = connection

	return sqliteConn
}

// Connect - connect
func (connection *SQLiteConnection) Connect() (*gorm.DB, error) {

	db, err := gorm.Open("sqlite3", connection.host)
	if err != nil {
		return nil, err
	}
	return db, err
}
