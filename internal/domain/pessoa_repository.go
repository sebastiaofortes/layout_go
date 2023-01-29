package domain

//repository lida com a persistência de agregados
//para persistir esses dados ele pode usar um dao ou um ORM
type ImplementsPessoaRepository struct{

}

//Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
//nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewImplementsPessoaRepository() PessoaRepository{
	return ImplementsPessoaRepository{}
}