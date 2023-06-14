package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// "github.com/gocql/gocql"

	"waroeng_pgn1/internal/bootstrap"
)

// type DBConnection struct {
// 	cluster *gocql.ClusterConfig
// 	session *sql.DB
// }

// var connection DBConnection
var DB *sql.DB

// ConnectToDB creates a connection to the MySQL database and returns a pointer to the database
func ConnectToDB() (*sql.DB, error) {
	// connection.cluster = gocql.NewCluster(bootstrap.NewEnv().CAHost)
	// connection.cluster.Keyspace = bootstrap.NewEnv().DBName
	// connection.cluster.Consistency = gocql.Quorum
	// connection.cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
	// // cluster.ProtoVersion = 4
	// // cluster.ConnectTimeout = 10 * time.Second
	// // cluster.Timeout = 10 * time.Second
	// // cluster.NumConns = 2
	// // cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	// var err error
	// connection.session, err = connection.cluster.CreateSession()
	// if err != nil {
	// 	return nil, err
	// }
	// return connection.session, nil
	db, err := sql.Open("mysql", bootstrap.NewEnv().DBUser+":"+bootstrap.NewEnv().DBPass+"@tcp(127.0.0.1:"+bootstrap.NewEnv().DBPort+")/waroeng_pgn1")
	if err != nil {
		return nil, err
	}
	fmt.Println("WDWDWDW", db)
	if db == nil {
		fmt.Println(bootstrap.NewEnv().DBUser + ":" + bootstrap.NewEnv().DBPass + "@tcp(127.0.0.1:" + bootstrap.NewEnv().DBPort + ")")
	}

	return db, nil
}

// func ConnectToDB() {
// 	connection.cluster = gocql.NewCluster(bootstrap.NewEnv().CAHost)
// 	connection.cluster.Keyspace = bootstrap.NewEnv().DBName
// 	connection.cluster.Consistency = gocql.Quorum
// 	connection.cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
// 	// cluster.ProtoVersion = 4
// 	// cluster.ConnectTimeout = 10 * time.Second
// 	// cluster.Timeout = 10 * time.Second
// 	// cluster.NumConns = 2
// 	// cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

// 	connection.session, _ = connection.cluster.CreateSession()

// }

// func ExecuteQuery(query string, args ...interface{}) {
// 	if err := connection.session.Query(query, args...).Exec(); err != nil {
// 		panic(err)
// 	}
// }
