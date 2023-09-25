package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

// Essa função serve para emprestar um livro e diminuir a quantidade dele salva no estoque
func Emprestando(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Busco o usuário pelo ID
	usuario_id, erro := strconv.ParseUint(parametro["usuario_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID", 422)
		return
	}

	//Busco cada dado do usuário
	usuariobuscado, erro := buscandoUMUsuario(int(usuario_id))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
	}

	//Se o nome do usuário buscado estiver vazio ele retorna a mensagem
	if usuariobuscado.Nome == "" {
		fmt.Println("Usuário não encontrado")
	}

	//Busca o livro pelo ID
	livro_id, erro := strconv.ParseUint(parametro["livro_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID", 422)
		return
	}

	//Busca os dados que estão naquele livro com o ID especificado
	livrobuscado, erro := BuscandoUMLivro(int(livro_id))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	//Leio todos os dados do corpo da requisição feita
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler os dados do corpo", 422)
		return
	}

	//Crio uma variavel para que possa fazer a conversão de json para struct
	var emprestar dados.EmprestimoDevolucao
	erro = json.Unmarshal(corpoRequisicao, &emprestar)
	if erro != nil {
		TratandoErros(w, "Erro ao converter json para struct", 422)
		return
	}

	//Aqui executo apenas se o valor que estou querendo for menor que o salvo no estoque
	//Caso for maior que o que tenho ele vai para o else e da a mensasgem do estoque insuficiente
	if livrobuscado.Estoque > emprestar.Quantidade {
		livrobuscado.Estoque = livrobuscado.Estoque - emprestar.Quantidade
		emprestar.Nome_Usuario = usuariobuscado.Nome
		emprestar.Titulo_livro = livrobuscado.Titulo
		emprestar.Data_Emprestimo = time.Now()
		emprestar.Data_Devolucao = emprestar.Data_Emprestimo.Add(15 * 24 * time.Hour)
		emprestar.Taxa_Emprestimo = float64(emprestar.Quantidade) * 5.50

		fmt.Println("Usuário:", emprestar.Nome_Usuario)
		fmt.Println("Titulo selecionado:", emprestar.Titulo_livro)
		fmt.Println("A taxa cobrada foi de:", emprestar.Taxa_Emprestimo)
		fmt.Println("A data do emprestimo é:", emprestar.Data_Emprestimo.Format("02/01/2006 03:04:05"))
		fmt.Println("A data da devolução será:", emprestar.Data_Devolucao.Format("02/01/2006 03:04:05"))

	} else {
		fmt.Println("Estoque insuficiente")
	}

	//Faço a alteração do estoque
	erro = AlterarEstoque(int(livro_id), livrobuscado.Estoque)
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	//Executo a função que vai fazer a conexão com o banco (mais informações no arquivo "comandosBancoErro")
	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar no banco de dados", 422)
		return
	}
	defer db.Close()

	//Crio o statement que vai fazer a alteração e salvar os dados na tabela do banco
	statement, erro := db.Prepare("insert into emprestimo_devolucao(nome_usuario, titulo_livro, data_emprestimo, data_devolucao, taxa_emprestimo) values (?, ?, ?, ?, ?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	//Executo o statement e salvo os novos dados inseridos
	_, erro = statement.Exec(emprestar.Nome_Usuario, emprestar.Titulo_livro, emprestar.Data_Emprestimo.Format("2006-01-02"), emprestar.Data_Devolucao.Format("2006-01-02"), emprestar.Taxa_Emprestimo)
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statement", 422)
		return
	}

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Emprestimo realizado com sucesso", 200)
	return
}
