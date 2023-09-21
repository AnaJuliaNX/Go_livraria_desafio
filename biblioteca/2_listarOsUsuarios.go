package biblioteca

import (
	"encoding/json"
	"net/http"
)

func ListarOsUsuarios(w http.ResponseWriter, r *http.Request) {

	usuarios, erro := BuscandoOSUsuarios()
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	erro = json.NewEncoder(w).Encode(usuarios)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	TratandoErros(w, "Usuarios buscados com sucesso", 200)
	return
}
