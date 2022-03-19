package interfaces

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zainafifaldi/digicafe/domain/entity"
	"github.com/zainafifaldi/digicafe/interfaces/serializer"
)

func TestGetAllStore_Success(t *testing.T) {
	//application.FoodApp = &fakeFoodApp{} //make it possible to change real method with fake

	//Return Food to check for, with our mock
	storeApp.GetAllStoreFn = func(options map[string]interface{}) ([]entity.Store, error) {
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

	req, err := http.NewRequest(http.MethodGet, "/stores", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/stores", s.GetAllStore)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response serializer.SuccessResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}
	assert.Equal(t, rr.Code, 200)
	assert.EqualValues(t, len(response.Data.([]entity.Store)), 2)
}

func TestGetStore_Success(t *testing.T) {
	//Return Store to check for, with our mock
	storeApp.GetStoreFn = func(storeId uint64) (*entity.Store, error) {
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

	storeId := strconv.Itoa(1)
	req, err := http.NewRequest(http.MethodGet, "/stores/"+storeId, nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/stores/:store_id", s.GetStoreDetail)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response serializer.SuccessResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	expectedStore := entity.Store{}
	bodyBytes, _ := json.Marshal(response.Data)
	json.Unmarshal(bodyBytes, &expectedStore)

	// fmt.Println()
	// fmt.Println("**************")
	// fmt.Println("==============")
	// fmt.Println(storeApp.GetStoreFn)
	// fmt.Println("==============")
	// fmt.Println("**************")
	// fmt.Println()

	assert.Equal(t, 200, rr.Code)

	assert.EqualValues(t, storeId, expectedStore.ID)
	assert.EqualValues(t, "Main store", expectedStore.Name)
}
