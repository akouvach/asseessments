package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/akouvach/assessments/models"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
)

const version = "1.0.0"

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Database    string `json:"database"`
	Dsn         string `json:"dsn"`
	Port        int    `json:"port"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

type config struct {
	port int
	env  string
	db   struct {
		dsn      string
		name     string
		database string
	}
	ctx context.Context
	jwt struct {
		secret string
	}
}

// $2a$12$UFTSTMDTByXM1D/HyflPVexkQHLzFM9UsK4u878SmWZXAdl2rtUja
// passwordpope

func main() {

	var cfg config
	// const (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "postgres"
	// 	password = "your-password"
	// 	dbname   = "calhounio_demo"
	//   )

	//   psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//   "password=%s dbname=%s sslmode=disable",
	//   host, port, user, password, dbname)

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env_port, err := strconv.Atoi(os.Getenv("APPPORT"))
	if err != nil {
		log.Fatal("Error en el parámetro de puerto")
	}

	env_env := os.Getenv("ENVIRONMENT")
	if err != nil {
		log.Fatal("Error en el parámetro de environment")
	}

	env_dsn := os.Getenv("DSN")
	if err != nil {
		log.Fatal("Error en el parámetro de dsn")
	}
	env_database := os.Getenv("DATABASE")
	if err != nil {
		log.Fatal("Error en el parámetro de DATABASE")
	}

	env_dbname := os.Getenv("DBNAME")
	if err != nil {
		log.Fatal("Error en el parámetro de DBNAME")
	}

	flag.StringVar(&cfg.db.name, "dbname", env_dbname, "Name of the database")

	flag.StringVar(&cfg.db.database, "dbname", env_database, "Name of the database server")

	flag.IntVar(&cfg.port, "port", env_port, "Server port to listen on")

	// flag.StringVar(&cfg.env, "env", "development", "Application environmente (development|production)")

	flag.StringVar(&cfg.env, "env", env_env, "Application environmente (development|testing|production)")

	// flag.StringVar(&cfg.db.dsn, "dsn", "postgress://tcs@localhost/usca?sslmode=disable", "Postgres connection string")

	// flag.StringVar(&cfg.db.dsn, "dsn", "host=localhost port=5432 user=postgres password=mafalda dbname=usca sslmode=disable", "Postgres connection string")
	flag.StringVar(&cfg.db.dsn, "dsn", env_dsn, "Mongo DB connection string")

	// flag.StringVar(&cfg.jwt.secret, "jwt-secret", "$2a$12$UFTSTMDTByXM1D/HyflPVexkQHLzFM9UsK4u878SmWZXAdl2rtUja", "secret")

	flag.Parse()

	cfg.jwt.secret = os.Getenv("GO_JWT")
	fmt.Println("flag", cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cfg.ctx = ctx

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		if err := db.Disconnect(cfg.ctx); err != nil {
			panic(err)
		}
	}()

	// defer db.Disconnect(ctx)

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port: ", cfg.port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}

}

// func openDb(cfg config) (*mongo.Client, error) {
// 	// Set client options
// 	clientOptions := options.Client().ApplyURI(cfg.db.dsn)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Connect to MongoDB
// 	// db, err := mongo.Connect(context.TODO(), clientOptions)
// 	db, err := mongo.Connect(ctx, clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = db.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	return db, nil

// 	// db, err := sql.Open("postgres", cfg.db.dsn)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel()

// 	// err = db.PingContext(ctx)
// 	// if err != nil {
// 	// 	return nil, err

// 	// }
// }

func openDB(cfg config) (*mongo.Client, error) {

	fmt.Println("running on mongodb")
	clientOptions := options.Client().ApplyURI(cfg.db.dsn)

	db, err := mongo.Connect(cfg.ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping(cfg.ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return db, nil

	// db, err := sql.Open("postgres", cfg.db.dsn)
	// if err != nil {
	// 	return nil, err
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// err = db.PingContext(ctx)
	// if err != nil {
	// 	return nil, err

	// }
}
