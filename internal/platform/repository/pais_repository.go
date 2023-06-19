package repository

import (
	"github.com/sebastiaofortes/layout_go/internal/domain"
	"github.com/sebastiaofortes/layout_go/internal/platform/dao"
)

//repository lida com a persistência de agregados
//para persistir esses dados ele pode usar um dao ou um ORM
type ImplementsPaisRepository struct{
	paisDao dao.PaisDao
}

func (p *ImplementsPaisRepository) GetPais(i int32)(domain.Pais, error) {
	r := p.paisDao.GetPais(i)
	return domain.Pais{ExpectativaVida: float32(r.ExpectativaVida),
		IDH: float32(r.IDH),
		BaseDeCalculo: float32(r.BaseDeCalculo),
		}, nil
}

//Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
//nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewImplementsPaisRepository() PaisRepository{
	return &ImplementsPaisRepository{}
}