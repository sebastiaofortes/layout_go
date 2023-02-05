package dto

type Enfermidade struct {
	Id         int32  `gorm:"primary_key;not null" json:"-"`
	Nome       string `json:"nome"`
	Severidade int64  `json:"severidade"`
}
type EstadoFisico struct {
	Id            int32  `gorm:"primary_key;not null" json:"-"`
	Classificaçao int64  `json:"classificacao"`
	DataAvaliaçao string `json:"data_avaliacao"`
}
type Pessoa struct {
	Id           int32         `gorm:"primary_key;not null" json:"-"`
	Nome         string        `json:"nome"`
	Idade        int32         `json:"idade"`
	Enfermidades []Enfermidade `gorm:"ForeignKey:Id;AssociationForeignKey:Id" json:"enfermidades"`
	EstadoFisico EstadoFisico  `gorm:"ForeignKey:Id;AssociationForeignKey:Id" json:"estado_fisico"`
	Pais         int64         `json:"pais"`
}
