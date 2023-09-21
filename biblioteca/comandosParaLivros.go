package biblioteca

import (
	"errors"

	"github.com/AnaJuliaNX/desafio2/dados"
)

func BuscandoOSLivros() ([]dados.Livro, error) {

	db, erro := ConectandoNoBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select * from livro_cadastrado")
	if erro != nil {
		return nil, errors.New("erro ao buscar os livros")
	}
	defer linhas.Close()

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

func BuscandoUMLivro(ID int) (dados.Livro, error) {
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return dados.Livro{}, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select * from livro_cadastrado where id = ?", ID)
	if erro != nil {
		return dados.Livro{}, errors.New("erro ao buscar o livro")
	}
	defer linhas.Close()

	var livroencontrado dados.Livro
	if linhas.Next() {
		erro := linhas.Scan(&livroencontrado.ID, &livroencontrado.Titulo, &livroencontrado.Autor, &livroencontrado.Estoque)
		if erro != nil {
			return dados.Livro{}, errors.New("erro ao escanear o livro")
		}
	}
	return livroencontrado, nil
}
