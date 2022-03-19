package interfaces

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zainafifaldi/digicafe/application"
	"github.com/zainafifaldi/digicafe/infrastructure/auth"
	"github.com/zainafifaldi/digicafe/interfaces/serializer"
)

type Store struct {
	StoreApp application.StoreAppInterface
	Ti       auth.TokenInterface
	Ai       auth.AuthInterface
}

/* Store constructor */
func NewStore(storeApp application.StoreAppInterface, ai auth.AuthInterface, ti auth.TokenInterface) *Store {
	return &Store{
		StoreApp: storeApp,
		Ti:       ti,
		Ai:       ai,
	}
}

func (s *Store) GetAllStore(c *gin.Context) {
	var options = make(map[string]interface{})

	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid limit value")
		return
	}

	offset, err := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid offset value")
		return
	}

	options["limit"] = limit
	options["offset"] = offset
	options["active"] = true

	stores, err := s.StoreApp.GetAllStore(options)
	if err != nil {
		serializer.RenderError(c, []string{err.Error()}, 10001, nil)
		return
	}

	serializer.RenderResponseData(c, stores, map[string]interface{}{
		"http_status": http.StatusOK,
	})
}

func (s *Store) GetStoreDetail(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errStatus := http.StatusUnprocessableEntity
		serializer.RenderError(c, []string{err.Error()}, 10002, &errStatus)
		return
	}

	store, err := s.StoreApp.GetStoreDetail(storeID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errStatus := http.StatusNotFound
			serializer.RenderError(c, []string{"Store not found"}, 10003, &errStatus)
			return
		default:
			serializer.RenderError(c, []string{err.Error()}, 10004, nil)
			return
		}
	}

	serializer.RenderResponseData(c, &store, map[string]interface{}{
		"http_status": http.StatusOK,
	})
}
