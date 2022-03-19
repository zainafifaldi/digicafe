package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStore_Success(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	store, err := seedStore(db)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewStoreRepository(db)

	f, getErr := repo.GetStoreDetail(store.ID)

	assert.Nil(t, getErr)
	assert.EqualValues(t, f.Name, store.Name)
}

func TestGetAllStore_Success(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	_, err = seedStores(db)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewStoreRepository(db)
	stores, getErr := repo.GetAllStore(map[string]interface{}{"limit": 10, "offset": 0})

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(stores), 2)
}
