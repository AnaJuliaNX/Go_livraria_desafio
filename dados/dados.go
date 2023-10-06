package dados

import "time"

// Struct para os dados dos usuários
type Usuario struct {
	ID   int    `json:"id"`   //Automático pelo banco
	Nome string `json:"nome"` //Not null | Max: 20 | Postman: "nome"
}

// Struct para exibir os dados do usuário padronizados em um data
type DataUsuario struct {
	Data interface{}
}

// Struct para os dados dos livros
type Livro struct {
	ID      int     `json:"id"`      //Automatico pelo banco
	Titulo  string  `json:"titulo"`  //Not null | Max: 20 | Postman "titulo"
	Autor   string  `json:"autor"`   //Not null | Max: 20 | Postman "autor"
	Estoque int     `json:"estoque"` //Min: 1 | Postman "estoque"
	Valor   float64 `json:"valor"`   //Min: 1 | Postman "valor"
}

// Struct para exibir os dados do livro padronizado em um data
type DataLivros struct {
	Data interface{} `json:"data"`
}

// Struct para os dados do empréstimo
type EmprestimoDevolucao struct {
	Nome_Usuario    string    `json:"nome_usuario"`    //Usuário previamente criado
	Titulo_livro    string    `json:"titulo_livro"`    //Livro previamente criado
	Quantidade      int       `json:"quantidade"`      //No postman "quantidade"
	Data_Emprestimo time.Time `json:"data_emprestimo"` //Calculada no momento do empréstimno
	Data_Devolucao  time.Time `json:"data_devolucao"`  //Calculada no momento da devolução
	Taxa_Emprestimo float64   `json:"taxa_emprestimo"` //A taxa é cobrada quando tem atraso na devolução
}

// Struct para fazer a páginação dos dados obtidos
type Meta struct {
	Total        int64 `json:"total"`        //Total de itens cadastrados
	Current_page int64 `json:"current_page"` //Informa a página atual
	Total_pages  int64 `json:"total_pages"`  //Informa a quantidade de páginas
}

// Strut para exibir os dados da paginação padronizados em um data
type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"` //Passo todos os dados do meta
}

// Struct para padronizar os dados do TratandoErros
type DataMessageError struct {
	Message string `json:"message"` //Mensagem de erro que vai ser exibida
	Code    int64  `json:"code"`    //Código de erro que vai ser exibido
}

// Struct para exibir os dados padronizar em um data
type ResponseError struct {
	Data DataMessageError `json:"data"`
}

// Struct Para fazer o pedido padronizado em um data
type Pedido struct {
	ID              int    `json:"id"`              //Automático pelo banco
	User_cadastrado string `json:"user_cadastrado"` //Usuário previamente cadastrado
}

// Struxt para selecionar os itens do pedido padronizad em um data
type Itens struct {
	Pedido_feito     int     `json:"pedido_feito"`     //Pedido preciamente feito
	Livro_cadastrado string  `json:"livro_cadastrado"` //Livro previamente cadastrado
	Quantidade       int     `json:"quantidade"`       //Min: 1 | Max: 9999 | Postman: "quantidade"
	Valor_final      float64 `json:"valor_final "`     //Previamente feito, multiplicado pela quantidade
}
