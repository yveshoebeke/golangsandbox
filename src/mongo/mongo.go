package main

import (
  "fmt"
  "os"
  "time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
  "myconfig"
)


type Runtimeinfo struct {
        User string
        Runtime string
}

func main() {
  config := myconfig.Getconfig()
        //session, err := mgo.Dial("server1.example.com,server2.example.com")
        session, err := mgo.Dial(config.Database.Host)
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C(config.Database.Database)

        loc, _ := time.LoadLocation("UTC")
        now := time.Now().In(loc)
        fmt.Println(now) // UTC

        err = c.Insert(&Runtimeinfo{os.Getenv("USER"), now.String()})
        if err != nil {
                fmt.Println(err.Error())
                os.Exit(1)
        }

        result := Runtimeinfo{}
        err = c.Find(bson.M{"user": os.Getenv("USER")}).One(&result)
        if err != nil {
                fmt.Println(err.Error())
                os.Exit(1)
        }

        fmt.Println("User:", result.User,"Run time:", result.Runtime)

}
