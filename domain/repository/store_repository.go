package repository

import "github.com/zainafifaldi/digicafe/domain/entity"

type StoreRepository interface {
	GetAllStore(map[string]interface{}) ([]entity.Store, error)
	GetStoreDetail(uint64) (*entity.Store, error)
}
