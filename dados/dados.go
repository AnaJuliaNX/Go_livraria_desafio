package dados

import "time"

// Struct para os dados dos usuários
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// Struct para exibir os dados do usuário padronizados em um data
type DataUsuario struct {
	Data interface{}
}

// Struct para os dados dos livros
type Livro struct {
	ID      int    `json:"id"`
	Titulo  string `json:"titulo"`
	Autor   string `json:"autor"`
	Estoque int   `json:"estoque"`
	Valor float64 `json:"valor"`
}

// Struct para exibir os dados do livro padronizado em um data
type DataLivros struct {
	Data interface{} `json:"data"`
}

// Struct para os dados do empréstimo
type EmprestimoDevolucao struct {
	Nome_Usuario    string    `json:"nome_usuario"`
	Titulo_livro    string    `json:"titulo_livro"`
	Quantidade      int       `json:"quantidade"`
	Data_Emprestimo time.Time `json:"data_emprestimo"`
	Data_Devolucao  time.Time `json:"data_devolucao"`
	Taxa_Emprestimo float64   `json:"taxa_emprestimo"` //A taxa é cobrada quando tem atraso na devolução
}

// Struct para fazer a páginação dos dados obtidos
type Meta struct {
	Total        int64 `json:"total"`        //total de itens cadastrados
	Current_page int64 `json:"current_page"` //informa a página atual
	Total_pages  int64 `json:"total_pages"`  //informa a quantidade de páginas
}

// Strut para exibir os dados da paginação padronizados em um data
type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"` //passo todos os dados do meta
}

// Struct para padronizar os dados do TratandoErros
type DataMessageError struct {
	Message string `json:"message"` //mensagem de erro que vai ser exibida
	Code    int64  `json:"code"`    //código de erro que vai ser exibido
}

// Struct para exibir os dados padronizar em um data
type ResponseError struct {
	Data DataMessageError `json:"data"`
}

type Pedido struct {
	ID int `json:"id"`
	User_cadastrado string `json:"user_cadastrado"`
}

type Itens struct {
	Pedido_feito int `json:"pedido_feito"`
	Livro_cadastrado string `json:"livro_cadastrado"`
	Quantidade int `json:"quantidade"`
	Valor_final float64 `json:"valor_final "`
}