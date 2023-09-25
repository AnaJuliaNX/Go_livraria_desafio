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
	linhas, erro := db.Query("select nome_usuario, titulo_livro, data_emprestimo, data_devolucao, taxa_emprestimo from emprestimo_devolucao where nome_usuario = ? and titulo_livro = ?", usuariobuscado.Nome, livrobuscado.Titulo)
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
			&emprestado.Data_Devolucao,
			&emprestado.Taxa_Emprestimo,
		)

		if err != nil {
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

	//data pre determinada, para fins de teste apenas
	//devolucaoHoje := time.Date(2023, 10, 21, 20, 10, 0, 0, time.Local)

	//Data da devolução em tempo real
	devolucaoHoje := time.Now()

	//Formato para que consiga fazer a comparação das datas
	hojeFormatada := devolucaoHoje.Unix()
	devolucaoFormatada := emprestado.Data_Devolucao.Unix()

	//Se a devolução foi feita até o prazo de quinze dias acabar executo o if, se foi depois executo o else
	if hojeFormatada < devolucaoFormatada {
		livrobuscado.Estoque = livrobuscado.Estoque + devolver.Quantidade
		fmt.Println("Usuário:", emprestado.Nome_Usuario)
		fmt.Println("Livro:", emprestado.Titulo_livro)
		fmt.Println("Devolução feita antes do prazo de 15 dias")

	} else {
		livrobuscado.Estoque = livrobuscado.Estoque + emprestado.Quantidade
		fmt.Println("Usuário:", emprestado.Nome_Usuario)
		fmt.Println("Livro:", emprestado.Titulo_livro)
		fmt.Println("Devolução feita após o prazo de quinze dias")
		fmt.Println("Multa de: R$", emprestado.Taxa_Emprestimo)

	}

	//Executo esse comando para fazer a alteração do estoque
	erro = AlterarEstoque(int(livro_id), livrobuscado.Estoque)
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	//Dados que vou registrar no banco
	devolver.Nome_Usuario = emprestado.Nome_Usuario
	devolver.Titulo_livro = emprestado.Titulo_livro
	devolver.Data_Emprestimo = emprestado.Data_Emprestimo
	devolver.Data_Devolucao = devolucaoHoje

	//Crio o statement para fazer a alteração e salvar os dados na tabela do banco
	statement1, erro := db.Prepare("insert into emprestimo_devolucao(nome_usuario, titulo_livro, data_emprestimo, data_devolucao) values (?, ?, ?, ?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement1.Close()

	//Executo o statment e salvo os dados alterados
	_, erro = statement1.Exec(devolver.Nome_Usuario, devolver.Titulo_livro, devolver.Data_Emprestimo.Format("2006-01-02"), devolver.Data_Devolucao.Format("2006-01-02"))
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statement", 422)
		return
	}

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Devolução realizada com sucesso", 200)
	return
}
