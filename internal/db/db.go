package db

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type DatabaseConfig struct {
	Port     string `env:"DATABASE_PORT"`
	Dbname   string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DBAddr   string `env:"POSTGRES_ADDR"`
}

var cfg DatabaseConfig

var ConnectionPool *pgxpool.Pool

func getConfig() *pgxpool.Config {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
		return nil
	}

	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	conf, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.DBAddr, cfg.Port, cfg.Dbname))
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	conf.MaxConns = defaultMaxConns
	conf.MinConns = defaultMinConns
	conf.MaxConnLifetime = defaultMaxConnLifetime
	conf.MaxConnIdleTime = defaultMaxConnIdleTime
	conf.HealthCheckPeriod = defaultHealthCheckPeriod
	conf.ConnConfig.ConnectTimeout = defaultConnectTimeout

	conf.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	conf.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	conf.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return conf
}

func InitDB() error {
	connPool, err := pgxpool.NewWithConfig(context.Background(), getConfig())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}
	defer connection.Release()
	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
	}
	fmt.Println("Connected to the database!!")

	ConnectionPool = connPool

	err = createTables()
	if err != nil {
		return err
	}

	return nil
}

func createTables() error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users(
    	id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    	email text NOT NULL UNIQUE,
    	password varchar(100) NOT NULL
	)`

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
    	id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    	name varchar(30) NOT NULL ,
    	description TEXT NOT NULL,
    	location varchar(30) NOT NULL ,
    	dateTime timestamp NOT NULL,
    	userID int REFERENCES users(id)
	)
	`

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations(
    	id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    	eventID int REFERENCES events(id),
    	userID int REFERENCES users(id)
	)
	`

	connection, err := ConnectionPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}

	_, err = connection.Exec(context.Background(), createUserTable)
	if err != nil {
		return err
	}

	_, err = connection.Exec(context.Background(), createEventsTable)
	if err != nil {
		return err
	}

	_, err = connection.Exec(context.Background(), createRegistrationTable)
	if err != nil {
		return err
	}

	connection.Release()
	return nil
}
