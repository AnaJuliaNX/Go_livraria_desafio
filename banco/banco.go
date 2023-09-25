package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarNoBanco() (*sql.DB, error) {
	stringdeconexao := "funcionario:trabalho@/biblioteca?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringdeconexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
