package biblioteca

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	usuario_id, erro := strconv.ParseUint(parametro["usuario_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	usuariobuscado, erro := buscandoUMUsuario(int(usuario_id))
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o usuário", 422)
		return
	}

	fmt.Println(usuariobuscado)
}
