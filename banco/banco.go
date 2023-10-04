package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarNoBanco() (*sql.DB, error) {
	stringdeconexao := "funcionario:trabalho@/livraria?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringdeconexao)
	if erro != nil {
		return nil, erro
	}
	erro = db.Ping()
	if erro != nil {
		return nil, erro
	}
	return db, nil
}
