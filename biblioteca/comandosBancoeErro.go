package biblioteca

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/banco"
)

func TratandoErros(w http.ResponseWriter, message string, statuscode int) {

	w.WriteHeader(statuscode)
	w.Write([]byte(message))
}

func ConectandoNoBanco() (*sql.DB, error) {
	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		return nil, errors.New("erro ao se conectar com o banco de dados")
	}
	return db, nil
}
