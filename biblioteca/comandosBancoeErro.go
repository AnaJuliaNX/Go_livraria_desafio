package biblioteca

import (
	"database/sql"
	"encoding/json"
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

	//acima de 200: informa que está ok / acima de 300: redirecionando para outro lugar
	//acima de 400: informando um erro da parte do usuário, front / acima de 500: erro do servidor
	var dataMessage dados.DataMessageError
	dataMessage.Message = message
	dataMessage.Code = int64(statuscode)

	var data dados.ResponseError
	data.Data = dataMessage

	messagem, erro := json.Marshal(data)
	if erro != nil {
		w.WriteHeader(422)            //definindo o status
		w.Write([]byte(erro.Error())) //transformando minha msg em slice of bytes
		return
	}
	w.WriteHeader(statuscode) //definindo o status
	w.Write(messagem)         //transformando minha msg em slice of bytes
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

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
func Paginacao(totalDeDados int64, dadosDoRetorno interface{}) dados.Response {

	var meta dados.Meta
	meta.Current_page = 1     //Página atual
	meta.Total = totalDeDados //Literalmente total de páginas

	totalDePaginas := totalDeDados / 15 //Total de páginas que tenho
	if totalDePaginas > 0 {
		meta.Total_pages = int64(math.Round(float64(totalDePaginas)))

	}
	if meta.Total != 0 && meta.Total_pages == 0 {
		meta.Total_pages = 1
	}
	var response dados.Response
	response.Data = dadosDoRetorno
	response.Meta = meta

	return response
}
