package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/AnaJuliaNX/desafio2/banco"
)

// Função para adicionar um novo usuário no banco de dados, especificamente na tabela usuário
func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {

	//Aqui faço a leitura dos dados do corpo, ou seja, o que o funcionário da biblioteca escreveu
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	var body map[string]interface{}
	//Converto de json para struct, ou seja, de um jeito que seja legivel para o Go
	if erro = json.Unmarshal(corpoDaRequisicao, &body); erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	//Verifica se o campo nome foi preenchido, exibe uma mensagem de erro caso esteja vazio
	if body["nome"] == nil {
		TratandoErros(w, "O campo nome é obrigatório", 422)
		return
	}
	//verifico se o que foi digitado é uma string, se não for exibo a mensagem
	if reflect.TypeOf(body["nome"]).Kind() != reflect.String {
		TratandoErros(w, "Nome inválido", 422)
		return
	}
	//Se o campo do nome estiver vazio exibo a mensagem de erro
	if body["nome"] == "" {
		TratandoErros(w, "O campo nome não pode estar vazio", 422)
		return
	}
	//Limito o tanto de carcteres para 20 e se atingir exibe a mensagem de erro
	limiteDeCaracter := 20
	if len(body["nome"].(string)) > limiteDeCaracter {
		TratandoErros(w, "Limite de caracteres superior a vinte (20) no campo de nome", 422)
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
	inserir, erro := statement.Exec(body["nome"])
	if erro != nil {
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
}
