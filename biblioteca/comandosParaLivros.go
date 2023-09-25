package biblioteca

import (
	"errors"

	"github.com/AnaJuliaNX/desafio2/dados"
)

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai buscar e exibir todos os livros que tenho previamente cadastrados no banco
func BuscandoOSLivros() ([]dados.Livro, error) {

	//Chamo a função que faz a conexão com o banco de dados (mais detalhes no arquivo "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	//Seleciono a minha tabela de livros cadastrado no banco de dados e pego todos os livros
	linhas, erro := db.Query("select * from livro_cadastrado")
	if erro != nil {
		return nil, errors.New("erro ao buscar os livros")
	}
	defer linhas.Close()

	//Escaneio todos os dados sobre os livros que foram encontrados no banco
	var livros []dados.Livro
	for linhas.Next() {
		var livro dados.Livro

		if erro := linhas.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Estoque); erro != nil {
			return nil, errors.New("erro ao escanear os livros")
		}

		livros = append(livros, livro)
	}
	return livros, nil
}

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai buscar pelo ID e exibir apenas um livro que tenho previamente cadastrado no banco de dados
func BuscandoUMLivro(ID int) (dados.Livro, error) {

	//Chamo a função que faz a conexão com o banco de dados (mais detalhes no arquivo "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return dados.Livro{}, erro
	}
	defer db.Close()

	//Seleciono a minha tabela de livros cadastrado no banco de dados e busco pelo ID especifico
	linhas, erro := db.Query("select * from livro_cadastrado where id = ?", ID)
	if erro != nil {
		return dados.Livro{}, errors.New("erro ao buscar o livro")
	}
	defer linhas.Close()

	//Escaneio todos os dados sobre aquele livro em especifico que foi encontrado no banco
	var livroencontrado dados.Livro
	if linhas.Next() {
		erro := linhas.Scan(&livroencontrado.ID, &livroencontrado.Titulo, &livroencontrado.Autor, &livroencontrado.Estoque)
		if erro != nil {
			return dados.Livro{}, errors.New("erro ao escanear o livro")
		}
	}
	return livroencontrado, nil
}
