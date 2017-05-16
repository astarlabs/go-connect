package data

import "testing"

func TestConnect(t *testing.T) {

	conn := NewMySQLConnection()
	_, err := conn.Connect()

	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("Sucessfuly connected")
	}

}
