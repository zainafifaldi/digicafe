package application

import (
	"github.com/zainafifaldi/digicafe/domain/entity"
	"github.com/zainafifaldi/digicafe/domain/repository"
)

type storeApp struct {
	sr repository.StoreRepository
}

type StoreAppInterface interface {
	GetAllStore(map[string]interface{}) ([]entity.Store, error)
	GetStoreDetail(uint64) (*entity.Store, error)
}

var _ StoreAppInterface = &storeApp{}

func (s *storeApp) GetAllStore(options map[string]interface{}) ([]entity.Store, error) {
	return s.sr.GetAllStore(options)
}

func (s *storeApp) GetStoreDetail(storeID uint64) (*entity.Store, error) {
	return s.sr.GetStoreDetail(storeID)
}
