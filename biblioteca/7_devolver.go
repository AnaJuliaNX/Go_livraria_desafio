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

// Função para fazer a devolução de um livro, calcular e aplicar multa se houver atraso
func Devolvendo(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Busco o usuário pelo ID
	usuario_id, erro := strconv.ParseUint(parametro["usuario_id"], 10, 32)
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	usuariobuscado, erro := buscandoUMUsuario(int(usuario_id))
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o usuário", 422)
		return
	}

	//Busco o livro pelo ID
	livro_id, erro := strconv.ParseUint(parametro["livro_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	livrobuscado, erro := BuscandoUMLivro(int(livro_id))
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o livro", 422)
		return
	}

	//Faz a conexão com o banco de dados
	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	//Busca todos os dados de todos os empréstimos para que eu possa achar a que preciso
	linhas, erro := db.Query("select nome_usuario, titulo_livro, data_emprestimo from emprestimo_devolucao where nome_usuario = ? and titulo_livro = ? and data_devolucao is null", usuariobuscado.Nome, livrobuscado.Titulo)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar dados dos emprestimos", 422)
		return
	}
	defer linhas.Close()

	//Crio uma váriavel para escanear todos os dados que busquei
	var emprestado dados.EmprestimoDevolucao
	if linhas.Next() {
		err := linhas.Scan(
			&emprestado.Nome_Usuario,
			&emprestado.Titulo_livro,
			&emprestado.Data_Emprestimo,
		)

		if err != nil {
			fmt.Println(erro)
			TratandoErros(w, "Erro ao escanear os dados dos empréstimos", 422)
			return
		}
	}

	//Leio todos os dados do corpo da requisição feita
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler os dados do corpo", 422)
		return
	}

	//Crio uma váriavel para converter de json para struct
	var devolver dados.EmprestimoDevolucao
	erro = json.Unmarshal(corpoRequisicao, &devolver)
	if erro != nil {
		TratandoErros(w, "Erro ao converter json para struct", 422)
		return
	}

	//Confiro se o usuário da devolução é o mesmo que fez emprestimo
	if usuariobuscado.Nome != emprestado.Nome_Usuario {
		fmt.Println(erro, 1)
		fmt.Println("As credenciais do usuário são incompativeis")
	}

	//Confiro se o livro a ser devolvido é o mesmo que foi emprestado
	if livrobuscado.Titulo != emprestado.Titulo_livro {
		fmt.Println(erro, 1)
		fmt.Println("Titulo do livro diferente do emprestado")
	}

	//Data da devolução em tempo real
	devolucaoHoje := time.Now()
	dataEsperada := emprestado.Data_Emprestimo.Add(15 * 24 * time.Hour)
	//Formato para que consiga fazer a comparação das datas
	hojeFormatada := devolucaoHoje.Unix()
	emprestimoFormatada := dataEsperada.Unix()
	fmt.Println(hojeFormatada, emprestimoFormatada)

	//Se a devolução foi feita até o prazo de quinze dias acabar executo o if, se foi depois executo o else
	if hojeFormatada <= emprestimoFormatada {
		livrobuscado.Estoque = livrobuscado.Estoque + devolver.Quantidade
		devolver.Data_Devolucao = devolucaoHoje
		fmt.Println("Usuário:", emprestado.Nome_Usuario)
		fmt.Println("Livro:", emprestado.Titulo_livro)
		fmt.Println("Devolução feita antes do prazo de 15 dias")

	} else {
		livrobuscado.Estoque = livrobuscado.Estoque + devolver.Quantidade
		devolver.Data_Devolucao = devolucaoHoje
		devolver.Taxa_Emprestimo = float64(devolver.Quantidade) * 5.50
		fmt.Println("Usuário:", emprestado.Nome_Usuario)
		fmt.Println("Livro:", emprestado.Titulo_livro)
		fmt.Println("Devolução feita após o prazo de quinze dias")
		fmt.Println("Multa de: R$", devolver.Taxa_Emprestimo)
	}

	//Executo esse comando para fazer a alteração do estoque
	erro = AlterarEstoque(int(livro_id), livrobuscado.Estoque)
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	//Faço essa série de comandos para dar um update na data de devolução que antes estava como nulo
	//E caso esteja atrasado também coloco uma multa
	linhas1, erro := db.Query("update emprestimo_devolucao set data_devolucao = ?, taxa_emprestimo = ? where nome_usuario = ? and titulo_livro = ? and data_devolucao is null", devolver.Data_Devolucao, devolver.Taxa_Emprestimo, usuariobuscado.Nome, livrobuscado.Titulo)
	if erro != nil {
		fmt.Println(erro)
		TratandoErros(w, "Erro ao fazer a atualização", 422)
		return
	}
	defer linhas1.Close()

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Devolução realizada com sucesso", 200)
	return
}
