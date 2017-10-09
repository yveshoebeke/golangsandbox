// accesses a mongodb to insdert and retrieve sessions.
package logusage

import (
  "fmt"
  "os"
  "time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
  "myconfig"
)

// sesson data record
type Runtimeinfo struct {
        User string
        Runtime string
}

// record current session
func Logit() {
  config := myconfig.Getconfig()
  session, err := mgo.Dial(config.Database.Host)
  if err != nil {
          panic(err)
  }
  defer session.Close()

  // Optional. Switching to monotonic session behavior. Data volume
  // very low here, no impact but leave it in as a matter of principle.
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(config.Database.Databasename).C(config.Database.Collectionname)

  loc, _ := time.LoadLocation("UTC")
  now := time.Now().In(loc)

  if err = c.Insert(&Runtimeinfo{os.Getenv("USER"), now.String()}); err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
}

// get all session data for user and display it.
func Showall(userlogin string) {
  config := myconfig.Getconfig()
  session, err := mgo.Dial(config.Database.Host)
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Optional. Switching to monotonic session behavior. Data volume
  // very low here, no impact but leave it in as a matter of principle.
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(config.Database.Databasename).C(config.Database.Collectionname)

  results := []Runtimeinfo{}

  if err = c.Find(bson.M{"user": userlogin}).All(&results); err != nil {
    panic(err)
  } else {
      if len(results) == 0 {
        fmt.Println("No previous access.")
      } else {}
        for i := range results {
          fmt.Printf("%d.\t%s\n",i+1,results[i].Runtime)
      }
  }
}
