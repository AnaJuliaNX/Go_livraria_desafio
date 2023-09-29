package dados

import "time"

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type Livro struct {
	ID      int    `json:"id"`
	Titulo  string `json:"titulo"`
	Autor   string `json:"autor"`
	Estoque int    `json:"estoque"`
}

type EmprestimoDevolucao struct {
	Nome_Usuario    string    `json:"nome_usuario"`
	Titulo_livro    string    `json:"titulo_livro"`
	Quantidade      int       `json:"quantidade"`
	Data_Emprestimo time.Time `json:"data_emprestimo"`
	Data_Devolucao  time.Time `json:"data_devolucao"`
	Taxa_Emprestimo float64   `json:"taxa_emprestimo"`
}

type RequestEmprestarLivro struct {
	Emprestado int `json:"emprestado"`
}

type Meta struct {
	Total_De_Itens int64 `json:"total_De_Itens"`
	Current_page   int64 `json:"current_page"` //informa a página atual
	Total_pages    int64 `json:"total_pages"`  //informa a quantidade de páginas
}

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"` //passo todos os dados do meta
}
