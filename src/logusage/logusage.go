// accesses a mongodb to insdert and retrieve sessions.
package logusage

import (
  "fmt"
  "os"
  "io"
  "time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
  "myconfig"
  "bytes"
  "encoding/hex"
  "crypto/md5"
)

// sesson data record
type Runtimeinfo struct {
        User string
        Navhash string
        Runtime string
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

// record current session
func Logit() {
  config := myconfig.Getconfig()
  var buffer bytes.Buffer
  var navplanfilename string
  // connect to mongo db
  session, err := mgo.Dial(config.Database.Host)
  if err != nil {
          panic(err)
  }
  defer session.Close()
  // Optional. Switching to monotonic session behavior. Data volume
  // very low here, no impact but leave it in as a matter of principle.
  session.SetMode(mgo.Monotonic, true)
  // reference the collection
  c := session.DB(config.Database.Databasename).C(config.Database.Collectionname)
  // construct the navplan data file name to get hash, so we can see what
  // the user was operating with.
  buffer.WriteString(os.Getenv("GOPATH"))
  buffer.WriteString(config.Navdata.Navplan)
  navplanfilename  = buffer.String()
  // get the current date time in UTC
  loc, _ := time.LoadLocation("UTC")
  now := time.Now().In(loc)
  // record it
  if err = c.Insert(&Runtimeinfo{os.Getenv("USER"), getMD5(navplanfilename), now.String()}); err != nil {
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

  // yeah, I could've definded the array in the beginning. left over
  // from my original version
  results := []Runtimeinfo{}
  // find all access records for this user and display to console
  if err = c.Find(bson.M{"user": userlogin}).All(&results); err != nil {
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
          hashchangeflag = " "
        }
        fmt.Printf("%d.\t%s\t%s %s\n", i+1, results[i].Runtime, results[i].Navhash, hashchangeflag)
      }
    }
  }
}
