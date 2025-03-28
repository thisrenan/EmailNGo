package endpoints

import (
	internalerrors "EmailNGo/internal/internal-errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerErro_when_endpoint_returns_internal_errors(t *testing.T) {
	assert := assert.New(t)
	//arrange
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	//act
	handlerFunc.ServeHTTP(res, req)

	//assert
	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}
