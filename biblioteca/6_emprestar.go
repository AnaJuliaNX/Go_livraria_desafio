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

func Emprestando(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Busca o usuário pelo ID
	usuario_id, erro := strconv.ParseUint(parametro["usuario_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID", 422)
		return
	}

	//Busco cada dado do usuário
	usuariobuscado, erro := BuscandoUMUsuario(int(usuario_id))
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

	//Busca os dados que estão naquele livro com aquele ID
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

	//Aqui executo apenas se o valor que estou querendo for menor que o do estoque
	//Caso for maior que o que tenho ele vai para o else e da a msg do estoque insuficiente
	if livrobuscado.Estoque > emprestar.Quantidade {
		livrobuscado.Estoque = livrobuscado.Estoque - emprestar.Quantidade
		emprestar.Nome_Usuario = usuariobuscado.Nome
		emprestar.Titulo_livro = livrobuscado.Titulo
		emprestar.Data_Emprestimo = time.Now()
		emprestar.Data_Devolucao = emprestar.Data_Emprestimo.Add(15 * 24 * time.Hour)
		emprestar.Taxa_Emprestimo = float64(emprestar.Quantidade) * 5.50

		//Se tiver tudo certo vai imprimir essas coisas
		fmt.Println("Usuário:", emprestar.Nome_Usuario)
		fmt.Println("Titulo selecionado:", emprestar.Titulo_livro)
		fmt.Println("A taxa cobrada foi de:", emprestar.Taxa_Emprestimo)
		fmt.Println("A data do emprestimo é:", emprestar.Data_Emprestimo.Format("02/01/2006 03:04:05"))
		fmt.Println("A data da devolução será:", emprestar.Data_Devolucao.Format("02/01/2006 03:04:05"))

	} else {
		//Se o estoque for menor do que eu to querendo emprestar ele exibe isso
		fmt.Println("Estoque insuficiente")
	}

	//Aqui eu faço a alteração do estoque, ou seja, reduzindo a quantidade que tenho salvo
	erro = AlterarEstoque(int(livro_id), livrobuscado.Estoque)
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	//Aqui abro o banco para fazer a alteração no estoque
	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar no banco de dados", 422)
		return
	}
	//fecho o banco quando terminar o que preciso
	defer db.Close()

	//Aqui crio um statemen e preparo ele para armazernar as coisas especificadas ali
	statement, erro := db.Prepare("insert into emprestimo_devolucao(nome_usuario, titulo_livro, data_emprestimo, data_devolucao, taxa_emprestimo) values (?, ?, ?, ?, ?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	_, erro = statement.Exec(emprestar.Nome_Usuario, emprestar.Titulo_livro, emprestar.Data_Emprestimo.Format("2006-01-02"), emprestar.Data_Devolucao.Format("2006-01-02"), emprestar.Taxa_Emprestimo)
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statement", 422)
		return
	}

	TratandoErros(w, "Emprestimo realizado com sucesso", 200)
	return
}
