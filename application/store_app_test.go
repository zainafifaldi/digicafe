package application

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zainafifaldi/digicafe/domain/entity"
)

type fakeStoreRepo struct{}

var (
	getStoreRepo    func(uint64) (*entity.Store, error)
	getAllStoreRepo func(map[string]interface{}) ([]entity.Store, error)
)

func (f *fakeStoreRepo) GetStoreDetail(storeId uint64) (*entity.Store, error) {
	return getStoreRepo(storeId)
}
func (f *fakeStoreRepo) GetAllStore(options map[string]interface{}) ([]entity.Store, error) {
	return getAllStoreRepo(options)
}

var storeAppFake StoreAppInterface = &fakeStoreRepo{}

func TestGetStore_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getStoreRepo = func(storeId uint64) (*entity.Store, error) {
		return &entity.Store{
			ID:                storeId,
			Type:              1,
			Name:              "Main Store",
			Description:       "This is main store",
			Address:           "Bandung, Jawa Barat",
			Active:            true,
			LocationLatitude:  1.23456,
			LocationLongitude: 2.34567,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}, nil
	}

	storeId := uint64(1)

	f, err := storeAppFake.GetStoreDetail(storeId)
	assert.Nil(t, err)
	assert.EqualValues(t, f.ID, storeId)
	assert.EqualValues(t, f.Name, "Main Store")
}

func TestGetAllStore_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getAllStoreRepo = func(map[string]interface{}) ([]entity.Store, error) {
		return []entity.Store{
			{
				Type:              1,
				Name:              "Main Store",
				Description:       "This is main store",
				Address:           "Kota Bandung, Jawa Barat",
				Active:            true,
				LocationLatitude:  1.23456,
				LocationLongitude: 2.34567,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
			{
				Type:              1,
				Name:              "Branch Store",
				Description:       "This is branch store",
				Address:           "Kab. Bandung Barat, Jawa Barat",
				Active:            true,
				LocationLatitude:  1.23568,
				LocationLongitude: 3.76543,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			},
		}, nil
	}

	f, err := storeAppFake.GetAllStore(map[string]interface{}{"limit": 10, "offset": 0})
	assert.Nil(t, err)
	assert.EqualValues(t, len(f), 2)
}
