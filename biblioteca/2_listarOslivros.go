package biblioteca

import (
	"encoding/json"
	"net/http"
)

func ListarOsLivros(w http.ResponseWriter, r *http.Request) {

	livros, erro := BuscandoOSLivros()
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	erro = json.NewEncoder(w).Encode(livros)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	TratandoErros(w, "Livros buscascados com sucesso", 200)
	return
}
