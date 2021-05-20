package models

type Student struct {
	tableName struct{} `pg:"students"`
	Id        int      `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	Age       int      `json:"age" pg:"age"`
	Class     string   `json:"class" pg:"class"`
}

type User struct {
	tableName struct{} `pg:"users"`
	Id        int      `json:"id" pg:"id"`
	UserName  string   `json:"user_name" pg:"user_name"`
	Password  string   `json:"password" pg:"password"`
}
