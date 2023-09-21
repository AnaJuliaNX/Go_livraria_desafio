package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
)

func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	var usuario dados.Usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuario(nome) values(?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statmente", 422)
		return
	}
	defer statement.Close()

	inserir, erro := statement.Exec(usuario.Nome)
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao executar o statmente", 422)
		return
	}

	_, erro = inserir.LastInsertId()
	if erro != nil {
		TratandoErros(w, "Erro ao obter o ID inserido", 422)
		return
	}

	TratandoErros(w, "Usuário adicionado com sucesso", 200)
	return
}
