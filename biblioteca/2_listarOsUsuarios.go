package biblioteca

import (
	"encoding/json"
	"net/http"
)

// Essa função serve para listar todos os usuários do banco de dados previamente cadastrados
func ListarOsUsuarios(w http.ResponseWriter, r *http.Request) {

	//Função que vai executar toda a busca dos usuários no banco (mais informações no arquivo "comandosOutros")
	usuarios, erro := BuscandoOSUsuarios()
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	//Executo a função que vai fazer a conexão com o banco (mais informações em "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar no banco de dados", 422)
		return
	}
	defer db.Close()

	//Selecioso apenas os id da tabela de usuários e conto quantos tem
	linhas, erro := db.Query("select count(id) from usuario")
	if erro != nil {
		TratandoErros(w, "Erro ao fazer a contagem", 422)
		return
	}
	defer linhas.Close()

	//Escaneio o resultado do comando de cima
	var totalDeUsuarios int64
	if linhas.Next() {
		erro := linhas.Scan(&totalDeUsuarios)
		if erro != nil {
			TratandoErros(w, "Erro ao escanear o total de usuários", 422)
			return
		}
	}

	//Se o resualtado do scan que fiz for igual a 0 exibo a mensagem de erro
	if totalDeUsuarios == 0 {
		TratandoErros(w, "Nenhum usuário encontrado", 404)
		return
	}

	//Faço todos os comandos de páginação necessários (mais informações em "ComandosaBancoeErro")
	response := Paginacao(totalDeUsuarios, usuarios)

	//Transformo os dados recebidos em struct para json
	erro = json.NewEncoder(w).Encode(response)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}
}
