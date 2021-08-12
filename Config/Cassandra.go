package Config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func InitCassandra() {
	cluster := gocql.NewCluster("192.168.100.102:9042", "192.168.100.102:9142", "192.168.100.102:9242")
    cluster.Keyspace = "example"
    cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: os.Getenv("CASSANDRA_USERNAME"), Password: os.Getenv("CASSANDRA_PASSWORD")}
    session, err := cluster.CreateSession()
	
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	
	defer session.Close()
	fmt.Println("Connection success!!!")

 
    if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
        "me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
        log.Fatal(err)
    }
 
    var id gocql.UUID
    var text string
 
    if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
        "me").Consistency(gocql.One).Scan(&id, &text); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Tweet:", id, text)
 
    iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
    for iter.Scan(&id, &text) {
        fmt.Println("Tweet:", id, text)
    }
    if err := iter.Close(); err != nil {
        log.Fatal(err)
    }
}