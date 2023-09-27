package biblioteca

import (
	"database/sql"
	"errors"
	"math"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
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

func Paginacao(totalDeDados int64, dadosDoRetorno interface{}) dados.Response {

	var meta dados.Meta
	meta.Total_pages = 1
	meta.Current_page = 1
	meta.Total = totalDeDados

	totalDePaginas := totalDeDados / 15
	if totalDePaginas > 0 {
		meta.Total_pages = int64(math.Round(float64(totalDePaginas)))
	}

	var response dados.Response
	response.Data = dadosDoRetorno
	response.Meta = meta

	return response
}
