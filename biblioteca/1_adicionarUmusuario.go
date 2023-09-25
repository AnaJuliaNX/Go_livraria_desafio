package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
)

// Função para adicionar um novo usuário no banco de dados, especificamente na tabela usuário
func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {

	//Aqui faço a leitura dos dados do corpo, ou seja, o que o funcionário da biblioteca escreveu
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	//Converto de json para struct, ou seja, de um jeito que seja legivel para o Go
	var usuario dados.Usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	//Executo a função que vai fazer a conexão com o banco (mais informações no arquivo "comandosBancoErro")
	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	//Crio um statement que vai receber os dados e salvar onde eu mandar, que nesse caso é no nome
	statement, erro := db.Prepare("insert into usuario(nome) values(?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statmente", 422)
		return
	}
	defer statement.Close()

	//Executo o satement, ou seja, salvo os dados inseridos na parte escolhida
	inserir, erro := statement.Exec(usuario.Nome)
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao executar o statmente", 422)
		return
	}

	//Vai retorna um novo Id onde salvei esse livro e por meio dele que vou pode pesquisar o livro depois
	_, erro = inserir.LastInsertId()
	if erro != nil {
		TratandoErros(w, "Erro ao obter o ID inserido", 422)
		return
	}

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Usuário adicionado com sucesso", 200)
	return
}
