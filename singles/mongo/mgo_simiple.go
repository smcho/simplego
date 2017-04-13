package main

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Phone string
}

var connUrl string = "mongodb://localhost:27017"
var useSSL bool

func init() {
	log.SetOutput(os.Stdout)

	args := os.Args[1:]
	if len(args) > 0 {
		connUrl = args[0]

		if len(args) > 1 {
			useSSL, _ = strconv.ParseBool(args[1])
		}
	}

	//mgo.SetDebug(true)
	mgo.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
}

func main() {
	if connUrl == "" {
		panic("the connection url are passed as the first program argument")
	} else {
		log.Printf("connecting to %s", connUrl)
	}

	log.Printf("useSSL = %v", useSSL)

	dialInfo, diErr := mgo.ParseURL(connUrl)
	if diErr != nil {
		panic(diErr)
	}

	if useSSL {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	log.Printf("DialInfo: %v", dialInfo)

	session, sErr := mgo.DialWithInfo(dialInfo)
	if sErr != nil {
		panic(sErr)
	}
	defer session.Clone()

	c := session.DB("").C("people")

	//iErr := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//	&Person{"Cla", "+55 53 8402 8510"})
	//if iErr != nil {
	//	log.Fatal(iErr)
	//}

	result := Person{}
	fErr := c.Find(bson.M{"name": "Ale"}).One(&result)
	if fErr != nil {
		log.Fatal(fErr)
	}

	log.Printf("found: %v", result.Phone)
}