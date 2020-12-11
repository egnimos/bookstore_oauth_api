package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	//Connect to Cassandra Cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	//creates the new session for the cassandra DB
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

//GetSession : returns the created session
func GetSession() *gocql.Session {
	return session
}
