package domain

type EstadoFisico struct {
	Classificaçao int64  `json:"classificacao"`
	DataAvaliaçao string `json:"data_avaliacao"`
}
