package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

func ItemDoPedido(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Busco o pedido pelo ID
	pedido_id, erro := strconv.ParseUint(parametro["pedido_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID do pedido", 422)
		return
	}

	//Busco cada um dos dados do pedido
	pedidobuscado, erro := BuscandoUMPedido(int(pedido_id))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	//Tratamento de erro caso não tenha sido encontrado o pedido
	if pedidobuscado.ID == 0 {
		TratandoErros(w, "Pedido não encontrado", 404)
		return
	}

	//Busco os dados do livro pelo ID
	livro_id, erro := strconv.ParseUint(parametro["livro_id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID do livro", 422)
		return
	}

	//Busco todos os dados desse livro
	livrobuscado, erro := BuscandoUMLivro(int(livro_id))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	//Tratamento de erro caso o livro não esteja cadastrado
	if livrobuscado.ID == 0 {
		TratandoErros(w, "Livro não encontrado", 404)
		return
	}

	//Leio todos os dados do corpo da requisição feita
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler os dados do corpo", 422)
		return
	}

	//Crio uma variavel para que possa fazer a conversão de json para struct
	var itens dados.Itens
	erro = json.Unmarshal(corpoDaRequisicao, &itens)
	if erro != nil {
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	if itens.Quantidade < 1 {
		TratandoErros(w, "Quantidade minima de um(1) item não atingida", 422)
		return
	}
	if itens.Quantidade > 9999 {
		TratandoErros(w, "Quantidade maxima de itens atingida", 422)
		return
	}
	if livrobuscado.Estoque > itens.Quantidade {
		livrobuscado.Estoque = livrobuscado.Estoque - itens.Quantidade
		itens.Pedido_feito = pedidobuscado.ID
		itens.Livro_cadastrado = livrobuscado.Titulo
		itens.Valor_final = livrobuscado.Valor * float64(itens.Quantidade)

	} else {
		TratandoErros(w, "Quantidade solicitada superior a quantidade em estoque", 422)
		return
	}

	erro = AlterarEstoque(int(livro_id), livrobuscado.Estoque)
	if erro != nil {
		TratandoErros(w, "Erro ao alterar o estoque", 422)
		return
	}

	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into itens(pedido_feito, livro_cadastrado, quantidade, valor_final) values(?, ?, ?, ?)")
	if erro != nil {
		fmt.Println(erro)
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	_, erro = statement.Exec(itens.Pedido_feito, itens.Livro_cadastrado, itens.Quantidade, itens.Valor_final)
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statement", 422)
		return
	}

	TratandoErros(w, "Venda realizada com sucesso", 200)
}
