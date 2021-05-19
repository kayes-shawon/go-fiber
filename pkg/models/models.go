package models

type Student struct {
	tableName struct{} `pg:"students"`
	Id        int      `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	Age       int      `json:"age" pg:"age"`
	Class     string   `json:"class" pg:"class"`
}
