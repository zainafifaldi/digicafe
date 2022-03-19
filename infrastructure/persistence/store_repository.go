package persistence

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zainafifaldi/digicafe/domain/entity"
	"github.com/zainafifaldi/digicafe/domain/repository"
)

type StoreRepo struct {
	Db *sql.DB
}

func NewStoreRepository(db *sql.DB) *StoreRepo {
	return &StoreRepo{db}
}

//StoreRepo implements the repository.StoreRepository interface
var _ repository.StoreRepository = &StoreRepo{}

func (s *StoreRepo) GetAllStore(options map[string]interface{}) ([]entity.Store, error) {
	if _, ok := options["limit"]; !ok {
		return nil, errors.New("invalid limit value")
	}

	if _, ok := options["offset"]; !ok {
		return nil, errors.New("invalid offset value")
	}

	whereClause := ""
	if _, ok := options["active"]; ok {
		whereClause = "WHERE active=true"
	}

	results, err := s.Db.Query(fmt.Sprintf("SELECT * FROM stores %s LIMIT %d OFFSET %d", whereClause, options["limit"], options["offset"]))
	if err != nil {
		return nil, err
	}

	var stores []entity.Store
	for results.Next() {
		var store entity.Store

		err = results.Scan(&store.ID, &store.Type, &store.Name, &store.Description, &store.Address, &store.Active, &store.LocationLatitude, &store.LocationLongitude, &store.CreatedAt, &store.UpdatedAt, &store.DeletedAt)
		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	return stores, nil
}

func (s *StoreRepo) GetStoreDetail(storeID uint64) (*entity.Store, error) {
	var store entity.Store

	row := s.Db.QueryRow(fmt.Sprintf("SELECT * FROM stores WHERE id=%d", storeID))
	err := row.Scan(&store.ID, &store.Type, &store.Name, &store.Description, &store.Address, &store.Active, &store.LocationLatitude, &store.LocationLongitude, &store.CreatedAt, &store.UpdatedAt, &store.DeletedAt)
	return &store, err
}
