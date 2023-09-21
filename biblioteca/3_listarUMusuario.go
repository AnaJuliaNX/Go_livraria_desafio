package biblioteca

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListarUMUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	usuariobuscado, erro := buscandoUMUsuario(int(ID))
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID do usuario", 422)
		return
	}

	erro = json.NewEncoder(w).Encode(usuariobuscado)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	TratandoErros(w, "Usuário buscado com sucesso", 200)
	return
}
