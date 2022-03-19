package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zainafifaldi/digicafe/domain/repository"
)

type Repositories struct {
	Store repository.StoreRepository
	// User  repository.UserRepository
	// Food  repository.FoodRepository
	Db *sql.DB
}

func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName, parseTime string) (*Repositories, error) {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s", DbUser, DbPassword, DbHost, DbPort, DbName, parseTime)

	db, err := sql.Open(DbDriver, dbUrl)
	if err != nil {
		return nil, err
	}
	// db.LogMode(true)

	return &Repositories{
		Store: NewStoreRepository(db),
		// User:  NewUserRepository(db),
		// Food:  NewFoodRepository(db),
		Db: db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.Db.Close()
}

//This migrate all tables
// func (s *Repositories) Automigrate() error {
// 	return s.db.AutoMigrate(&entity.User{}, &entity.Food{}).Error
// }
