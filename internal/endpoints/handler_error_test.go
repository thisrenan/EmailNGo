package endpoints

import (
	internalerrors "EmailNGo/internal/internal-errors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerErro_when_endpoint_returns_internal_error(t *testing.T) {
	//arrange
	assert := assert.New(t)
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

func Test_HandlerErro_when_endpoint_returns_domain_error(t *testing.T) {
	//arrange
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("domain error")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	//act
	handlerFunc.ServeHTTP(res, req)

	//assert
	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}

func Test_HandlerErro_when_endpoint_returns_obj_and_status(t *testing.T) {
	//arrange
	assert := assert.New(t)
	type bodyForTest struct {
		Id int
	}
	objExpected := bodyForTest{Id: 2}
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	//act
	handlerFunc.ServeHTTP(res, req)

	//assert
	assert.Equal(http.StatusCreated, res.Code)
	objReturned := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objExpected, objReturned)
}
