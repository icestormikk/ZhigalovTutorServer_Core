package database

import (
	"context"
	"log"
	"zhigalov_tutor_server_core/main/abstract/interfaces"
	"zhigalov_tutor_server_core/main/abstract/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresDatabase struct {
	db *gorm.DB
}

func NewPostgresDatabase(cfg interfaces.Configuration) *PostgresDatabase {
	databaseUrl, err := cfg.Get("POSTGRES_URL")
	if err != nil {
		log.Panicln(err)
	}

	databaseSchema, err := cfg.Get("POSTGRES_SCHEMA")
	if err != nil {
		log.Panicln(err)
	}

	databaseConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseSchema + ".",
			SingularTable: true,
		},
	}

	db, err := gorm.Open(postgres.Open(databaseUrl), databaseConfig)
	if err != nil {
		log.Panicln("Error while connecting to the database: " + err.Error())
	}
	log.Println("Database connection established")

	log.Println("Migrating..")
	err = db.AutoMigrate(&structs.User{})
	if err != nil {
		log.Panicln("Error while migrating the database: " + err.Error())
	}
	log.Println("Database migrated successfully")

	return &PostgresDatabase{db: db}
}

func (pd *PostgresDatabase) SelectUser(query any, args ...any) (*structs.User, error) {
	ctx := context.Background()

	first, err := gorm.G[structs.User](pd.db).Where(query, args).First(ctx)
	if err != nil {
		return nil, err
	}

	return &first, nil
}

func (pd *PostgresDatabase) SelectUsers(query *any, args ...any) (*[]structs.User, error) {
	ctx := context.Background()

	q := gorm.G[structs.User](pd.db)

	if query == nil {
		all, err := q.Find(ctx)
		if err != nil {
			return nil, err
		}
		return &all, nil
	}

	users, err := q.Where(query, args).Find(ctx)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
