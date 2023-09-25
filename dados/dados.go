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

type RequestDevolverLivro struct {
	Devolvido   int     `json:"devolvido"`
	TaxaCobrada float64 `json:"taxacobrada"`
}
