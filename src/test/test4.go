// accesses a mongodb to insdert and retrieve sessions.
package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"myconfig"
	"os"
	"time"
)

// sesson data record
type Runtimeinfo struct {
	User    string
	Navhash string
	Runtime time.Time
}

type mgoconn struct {
	host string
	db   string
	col  string
}

type mgosession struct {
	mgoconn
	mgosess *mgo.Session
	mgoerr  error
}

func (s *mgosession) connect() {
	s.mgosess, s.mgoerr = mgo.Dial(s.host)
}

type isession interface {
	connect()
}

func getSession(sess isession) {
	sess.connect()
}

func main() {
	flag := os.Args[1]

	config := myconfig.Getconfig()
	var userlogin string = os.Getenv("USER")

	co := mgosession{
		mgoconn{
			config.Database.Host,
			config.Database.Databasename,
			config.Database.Collectionname,
		}, nil, nil,
	}

	if flag == "a" {
		show1(userlogin, co)
	} else {
		show2(userlogin, co, config)
	}
}

func show1(user string, co mgosession) {
	co.connect()
	//getSession(co)
	if co.mgoerr != nil {
		panic(co.mgoerr)
	}
	defer co.mgosess.Close()

	c := co.mgosess.DB(co.db).C(co.col)
	results := []Runtimeinfo{}

	if err := c.Find(bson.M{"user": user}).Sort("-runtime").Limit(1).All(&results); err != nil {
		panic(err)
	} else {
		var hashchangeflag string
		if len(results) == 0 {
			fmt.Println("No previous access found.")
		} else {
			for i := range results {
				if i > 0 && results[i].Navhash != results[i-1].Navhash {
					hashchangeflag = "!"
				} else {
					hashchangeflag = ""
				}
				fmt.Printf("%d.\t%s\t%s %s\n", i+1, results[i].Runtime, results[i].Navhash, hashchangeflag)
			}
		}
	}
}

func show2(user string, co mgosession, conf *myconfig.Config) {
	co.connect()
	//getSession(co)
	if co.mgoerr != nil {
		panic(co.mgoerr)
	}
	defer co.mgosess.Close()

	c := co.mgosess.DB(co.db).C(co.col)
	results := []Runtimeinfo{}
	if err := c.Find(bson.M{"user": user}).Sort("-runtime").Limit(conf.Flags.Histoutputlimit).All(&results); err != nil {
		panic(err)
	} else {
		var hashchangeflag string
		if len(results) == 0 {
			fmt.Println("No previous access found.")
		} else {
			for i := range results {
				if i > 0 && results[i].Navhash != results[i-1].Navhash {
					hashchangeflag = "!"
				} else {
					hashchangeflag = ""
				}
				fmt.Printf("%d.\t%s\t%s %s\n", i+1, results[i].Runtime, results[i].Navhash, hashchangeflag)
			}
		}
	}
}
