package models

type Task struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    int    `json:"status"`
}
