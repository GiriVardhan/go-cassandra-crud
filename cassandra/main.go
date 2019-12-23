package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
        "database/sql"
     
)

// Session holds our connection to Cassandra
var Session *gocql.Session
var db *sql.DB

func init() {
	var err error
         cluster := gocql.NewCluster("172.100.0.10")
         Session, err = cluster.CreateSession()
         CreateDb()
         cluster.Keyspace = "clusterdb"
         Session, err = cluster.CreateSession()
	 if err != nil {
		panic(err)
	  }
         fmt.Println("cassandra init done")        
}


func CreateDb(){
     var err error
     err = Session.Query(`CREATE KEYSPACE IF NOT EXISTS clusterdb with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };`).Exec()
	if err != nil {
		panic(err)
	} else {
               err = Session.Query("CREATE TABLE IF NOT EXISTS clusterdb.emps ( empid text PRIMARY KEY, first_name text, last_name text, age int );").Exec()
	       if err != nil {
		  panic(err)
	       }

        }

        fmt.Println("cassandra keyspace and tables created")        
      
}
