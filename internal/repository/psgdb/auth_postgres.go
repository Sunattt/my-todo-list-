package psgdb

import (
	"fmt"
	"todo/pkg/todo_list"

	"github.com/jmoiron/sqlx"
)
type AuthPostgres struct{
	db *sqlx.DB
}

func NewAuthPostgres (db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo_list.User) (int, error){
	var id int 

	query := fmt.Sprintf("INSERT INTO %s (username, name, hash_password) values (1$, 2$, 3$) RETURNING id")
	
	row := r.db.QueryRow(query, usersTable, user.Name, user.Username, user.Password) // принимает запрос и произвольный набор аргумнтов которые будут поставлины плейсхолдеры из запроса 
					// заросы можно выполнять и без аргументов вообще ... И возврощает этот метод  объект роу тоесть он хранит информацию о возврощаемой строке в базу (в данный момет это id )
	if err := row.Scan(&id); err != nil{
		return 0, err
	}

	return id,  nil
}