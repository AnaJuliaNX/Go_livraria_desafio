package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

func Devolver(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	usuario_id, erro := strconv.ParseUint(parametro["usuario_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID do usuario", 422)
		return
	}

	usuariobuscado, erro := BuscandoUMUsuario(int(usuario_id))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	livro_id, erro := strconv.ParseUint(parametro["livro_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID do livro", 422)
		return
	}

	livrobuscado, erro := BuscandoUMLivro(int(livro_id))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler os dados do corpo", 422)
		return
	}

	var devolver dados.EmprestimoDevolucao
	erro = json.Unmarshal(corpoRequisicao, &devolver)
	if erro != nil {
		TratandoErros(w, "Erro ao converter json para struct", 422)
		return
	}

	//Pensar nisso aqui amanhã, preciso mesmo dessas informações? Serão uteis de alguma forma?
	//O que vou fazer com elas? Como vou comparar elas se já estou buscando por aquele ID?
	if usuariobuscado.Nome != devolver.Nome_Usuario {
		fmt.Println("O usuário não é compativel com o cadastrado")
	}

	if livrobuscado.Titulo != devolver.Titulo_livro {
		fmt.Println("O livro devolvido não é o mesmo que o cadastrado")
	} 

	dataDaDevolucao := time.Now()

	if 


}
