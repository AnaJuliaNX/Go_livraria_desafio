package biblioteca

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListarUMLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	livroencontrado, erro := BuscandoUMLivro(int(ID))
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	erro = json.NewEncoder(w).Encode(livroencontrado)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	TratandoErros(w, "Livro buscado com sucesso", 200)
	return
}
