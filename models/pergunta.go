package models

// Pergunta representa uma pergunta na aplicação.
// @Description Representa uma pergunta com um conjunto de respostas.
// @ID Pergunta
// @Property id uint `json:"id"`
// @Property pergunta string `json:"pergunta"`
// @Property respostas array[Resposta] `json:"respostas"` // Lista de respostas associadas
type Pergunta struct {
	ID        uint       `gorm:"primaryKey" json:"id"`                   // ID da pergunta
	Pergunta  string     `json:"pergunta"`                               // Texto da pergunta
	Respostas []Resposta `gorm:"foreignKey:PerguntaID" json:"respostas"` // Respostas associadas
}

var Perguntas []Pergunta
