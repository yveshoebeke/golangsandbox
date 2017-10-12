// accesses a mongodb to insdert and retrieve sessions.
package logusage

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
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

// mgo connection to host
type mgoconn struct {
	host string
	db   string
	col  string
}

// mgo session from mgo
type mgosession struct {
	mgoconn
	mgosess *mgo.Session
	mgoerr  error
}

// get a mgo session
func (s *mgosession) connect() {
	s.mgosess, s.mgoerr = mgo.Dial(s.host)
}

// get MD5 checksum of some file
func getMD5(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}

	decoded := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		panic(err)
	}

	return decoded
}

func Logaction(flag string, user string) {
	config := myconfig.Getconfig()

	if user == "" {
		user = os.Getenv("USER")
	}

	co := mgosession{
		mgoconn{
			config.Database.Host,
			config.Database.Databasename,
			config.Database.Collectionname,
		}, nil, nil,
	}

	switch flag {
		case "logit":
			logit(user, co, config)
		case "showit":
			showit(user, co, config)
	}
}

// record current session
func logit(user string, co mgosession, conf *myconfig.Config) {
	var buffer bytes.Buffer
	var navplanfilename string
	co.connect()
	if co.mgoerr != nil {
		panic(co.mgoerr)
	}
	defer co.mgosess.Close()
	// reference the collection
	c := co.mgosess.DB(co.db).C(co.col)
	// construct the navplan data file name to get hash, so we can see what
	// the user was operating with.
	buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString(conf.Navdata.Navplan)
	navplanfilename = buffer.String()
	// record it
	if err := c.Insert(&Runtimeinfo{user, getMD5(navplanfilename), time.Now()}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// get all session data for user and display it.
func showit(user string, co mgosession, conf *myconfig.Config) {
	co.connect()
	if co.mgoerr != nil {
		panic(co.mgoerr)
	}
	defer co.mgosess.Close()
	// reference the collection
	c := co.mgosess.DB(co.db).C(co.col)
	// find all access records for this user.
	// display last one or all depending on flag.
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
