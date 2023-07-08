package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	impl_mocks "github.com/sebastiaofortes/layout_go/tests/Impl_mocks"
	"github.com/stretchr/testify/assert"
)

func TestExpectativaDeVidaPorPaisSucess(t *testing.T) {
	eDv := impl_mocks.NewEsperançaDeVida(t)

	c := NewExpectativaController(eDv)

	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	eDv.On("CalcularExpectativaDeVidaPorPais", int32(1)).Return(float32(22), nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/expectativa-pais?pais=%v", "1"), nil)
	router.GET("/expectativa-pais", c.ExpectativaDeVidaPorPais())
	router.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)

	res := string(body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"22.000\"", res)
}
