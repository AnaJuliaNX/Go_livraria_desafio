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

	//Transformo os dados recebidos em struct para json, assim facilita o entendimento em outras linguagens
	erro = json.NewEncoder(w).Encode(usuarios)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Usuarios buscados com sucesso", 200)
	return
}
