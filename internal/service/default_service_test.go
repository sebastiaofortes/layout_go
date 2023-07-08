package service

import (
	"testing"

	"github.com/sebastiaofortes/layout_go/internal/domain"
	impl_mocks "github.com/sebastiaofortes/layout_go/tests/Impl_mocks"
	mock_values "github.com/sebastiaofortes/layout_go/tests/mock_values"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalcularExpectativaDeVida(t *testing.T) {
	pessoaR := impl_mocks.NewPessoaRepository(t)
	paisR := impl_mocks.NewPaisRepository(t)

	listMockPessoas := mock_values.Pessoas

	pessoaR.On("GetPessoasPorIdade", mock.Anything).Return(listMockPessoas, nil)
	paisR.On("GetPais", mock.Anything).Return(domain.Pais{}, nil)

	serv := NewDefaultService(pessoaR, paisR)

	result, err := serv.CalcularExpectativaDeVidaPorIdade(22)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, result)
}
