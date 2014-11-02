package db

import (
	"testing"
)

func Testdb(t *testing.T) {
	session, err := NewSession()
	if err != nil {
		t.Error(err)
	}
	t.Log(session.Host)
}
