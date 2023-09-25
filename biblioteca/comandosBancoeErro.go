package biblioteca

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/banco"
)

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai fazer o tratamento dos meus erros,
// ele vai exibir um código de status e uma mensagem personalizada escrita por mim
func TratandoErros(w http.ResponseWriter, message string, statuscode int) {

	w.WriteHeader(statuscode)
	w.Write([]byte(message))
}

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai abrir a conexão com o banco e será uma "simplificação" para quando
// for abrir a conexão com o banco em outros arquivos
func ConectandoNoBanco() (*sql.DB, error) {
	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		return nil, errors.New("erro ao se conectar com o banco de dados")
	}
	return db, nil
}
