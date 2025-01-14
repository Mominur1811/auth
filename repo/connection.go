package repo

import (
	"auth-repo/config"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type db struct {
	*sqlx.DB
	psql sq.StatementBuilderType
}

func getConnectionString(dbConfig config.DBConfig) string {

	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Name,
		dbConfig.Host,
		dbConfig.Port,
	)
	fmt.Println(connectionString)
	if !dbConfig.EnableSSLMode {
		connectionString += " sslmode=disable"
	}

	return connectionString
}

func connect(dbConfig config.DBConfig) (*sqlx.DB, error) {
	dbSource := getConnectionString(dbConfig)

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	dbCon.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleTimeInMinute * int(time.Minute)))

	return dbCon, nil
}

func ConnectDb(conf *config.Config) (*db, error) {

	dbCon, err := connect(conf.DB)
	if err != nil {
		return nil, err
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &db{dbCon, psql}, nil
}
