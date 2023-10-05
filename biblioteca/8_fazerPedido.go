package biblioteca

import(
	"encoding/json"
	"strconv"
	"net/http"
	"io"
	"fmt"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

//Essa função serve para que um pedido seja solicitado
func FazerPedido(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Busco o usuuário pelo ID
	usuario_id, erro := strconv.ParseUint(parametro["usuario_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID", 422)
		return
	}

	//Busco cada dados do usuário
	usuariobuscado, erro := buscandoUMUsuario(int(usuario_id))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para interio", 422)
	}

	//Tratamento de erro caso o nome não seja digitado
	if usuariobuscado.Nome == "" {
		TratandoErros(w, "Usuario não encontrado", 404)
		return
	}

	//Leio todos os dados do corpo da requisição feita
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	//Crio uma variável que vai fazer a conversão de json para struct
	var usuario dados.Pedido
	erro = json.Unmarshal(corpoDaRequisicao, &usuario)
	if erro != nil {
		fmt.Println(erro)
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	//Erro caso o usuário digitado for diferente do cadastrado
	if usuario.User_cadastrado != usuariobuscado.Nome {
		TratandoErros(w, "Usuario inserido é diferente do cadastrado", 422)
		return
	}

	//Função para fazer a conexão com o banco de dados 
	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}

	//Crio o statement que salvar os dados na tabela pedidos do banco de dados
	statement, erro := db.Prepare("insert into pedidos(user_cadastrado) values(?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	//Executo o statement e salvo os dados em uma das linhas da tabela
	_, erro = statement.Exec(usuario.User_cadastrado)
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statement",422)
		return
	}

	//Se não houve nenhum erro durante a execuçãodo código exibo essa mensagem no final
	TratandoErros(w, "Pedido solicitado com sucesso", 200)
}