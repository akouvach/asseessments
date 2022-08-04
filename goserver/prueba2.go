package main

import (
	"context"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type storage interface {
	open()
	close()
	getAll() interface{}
}

type db_mongo struct {
	database string
	conn     *mongo.Client
}

type db_postgres struct {
	database string
	conn     *sql.DB
}

type db_mysql struct {
	database string
	conn     *sql.DB
}

func (s *db_mongo) open() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	s.conn = client

}

// func (s *db_postgres) open(){

// }

// func (s *db_mysql) open(){

// }

func initDB(DBSERVER string) interface{} {
	if DBSERVER == "MONGODB" {
		return new(db_mongo)
	}
}

func main() {
	var db storage

	// 	const (
	// 		host     = "localhost"
	// 		port     = 5432
	// 		user     = "root"
	// 		password = "mafalda"
	// 		dbname   = "assessments"
	// 	  )
	// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	  "password=%s dbname=%s sslmode=disable",
	// 	  host, port, user, password, dbname)

	// 	  db, err := sql.Open("postgres", psqlInfo)
	// 	  if err != nil {
	// 		panic(err)
	// 	  }
	// 	  defer db.Close()

	// 	  err = db.Ping()
	// if err != nil {
	//   panic(err)
	// }

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	if 1 == 1 {
		db = initDB("MONGODB")
	}

	db.open()

	// coll := client.Database("assessments").Collection("usuarios")
	// title := "Kouvach"

	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{"apellido", title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)
}
