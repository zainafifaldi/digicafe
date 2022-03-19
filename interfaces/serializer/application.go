package serializer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const DEFAULT_STATUS = http.StatusOK

type MetaData struct {
	HttpStatus int `json:"http_status"`
	Limit      int `json:"limit,omitempty"`
	Offset     int `json:"offset,omitempty"`
	Total      int `json:"total,omitempty"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
	Meta MetaData    `json:"meta"`
}

type ErrorData struct {
	Code    int      `json:"code,omitempty"`
	Message []string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorData `json:"error"`
}

func RenderResponseData(c *gin.Context, data interface{}, options map[string]interface{}) {
	metaData := MetaData{HttpStatus: DEFAULT_STATUS}

	if _, ok := options["http_status"]; ok {
		metaData.HttpStatus, ok = options["http_status"].(int)
		if !ok {
			metaData.HttpStatus = DEFAULT_STATUS
		}
	}

	if _, ok := options["total"]; ok {
		metaData.Total, _ = options["total"].(int)
	}

	if _, ok := options["limit"]; ok {
		metaData.Limit, _ = options["limit"].(int)
	}

	if _, ok := options["offset"]; ok {
		metaData.Offset, _ = options["offset"].(int)
	}

	c.JSON(metaData.HttpStatus, SuccessResponse{Data: data, Meta: metaData})
}

func RenderError(c *gin.Context, message []string, code int, httpStatus *int) {
	errorData := ErrorResponse{
		Error: ErrorData{
			Code:    code,
			Message: message,
		},
	}

	httpStatusResponse := http.StatusInternalServerError
	if httpStatus != nil {
		httpStatusResponse = *httpStatus
	}

	c.JSON(httpStatusResponse, errorData)
}
