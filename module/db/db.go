package db

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"log"
	//"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

type Config struct {
	Host string
}

type DataStore struct {
	Config
	Session *mgo.Session
}

var (
	MongoHost DataStore
	RedisHost DataStore
)

func init() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &MongoHost)
	if err != nil {
		panic(err)
	}
	log.Printf("[+] Database Host : %s", MongoHost.Host)
	MongoHost.Session, err = mgo.Dial(MongoHost.Host)
	if err != nil {
		panic(err)
	}
}

func NewSession() (*DataStore, error) {
	return &DataStore{
		Config{
			Host: MongoHost.Host,
		},
		Session: MongoHost.Session,
	}, nil
}
