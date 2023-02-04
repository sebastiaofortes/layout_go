package repository

import "github.com/sebastiaofortes/layout_go/internal/domain"

//repository lida com a persistência de agregados
//para persistir esses dados ele pode usar um dao ou um ORM
type ImplementsPaisRepository struct{

}

func (p *ImplementsPaisRepository) GetPais(int32)(domain.Pais, error) {
	return domain.Pais{}, nil
}

//Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
//nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewImplementsPaisRepository() domain.PaisRepository{
	return &ImplementsPaisRepository{}
}