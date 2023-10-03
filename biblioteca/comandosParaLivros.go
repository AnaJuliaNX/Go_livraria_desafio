package biblioteca

import (
	"errors"
	"fmt"

	"github.com/AnaJuliaNX/desafio2/dados"
)

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai buscar e exibir todos os livros que tenho previamente cadastrados no banco
// Os livros serão exibidos em páginas contendo 15 itens em cada uma pra não sobrecarregar meu banco
func BuscandoOSLivros(search string, offset int) ([]dados.Livro, error) {

	//Chamo a função que faz a conexão com o banco de dados (mais detalhes no arquivo "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	//Limito a quantidade de livros que vou receber por busca no banco
	limit := 15
	//Seleciono a minha tabela de livros cadastrado no banco e pego todos os dados ordenando pelo id
	//até o tanto que o limit permite
	linhas, erro := db.Query("select id, titulo, autor, estoque from livro_cadastrado where (titulo like ? or autor like ?) order by id limit ? offset ?", "%"+search+"%", "%"+search+"%", limit, offset)
	if erro != nil {
		fmt.Println(erro)
		return nil, errors.New("erro ao buscar os livros")
	}
	defer linhas.Close()

	//Escaneio todos os dados sobre os livros que foram encontrados no banco
	var livros []dados.Livro
	var livro dados.Livro
	for linhas.Next() {
		//Escaneio todas as linhas buscando os dados dos livros, tais como ID, titulo, autor e estoque
		erro := linhas.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Estoque)
		if erro != nil {
			return nil, errors.New("erro ao escanear os livros")
		}

		livros = append(livros, livro)
	}
	//No fim retorno todos os dados dos livros que fiz o scan
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
	linhas, erro := db.Query("select id, titulo, autor, estoque from livro_cadastrado where id = ?", ID)
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
