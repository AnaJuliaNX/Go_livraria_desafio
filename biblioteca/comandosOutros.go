package biblioteca

import (
	"errors"

	"github.com/AnaJuliaNX/desafio2/dados"
)

func BuscandoOSUsuarios() ([]dados.Usuario, error) {

	db, erro := ConectandoNoBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuario")
	if erro != nil {
		return nil, errors.New("erro ao buscar os usuários")
	}
	defer linhas.Close()

	var usuarios []dados.Usuario
	for linhas.Next() {
		var usuario dados.Usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome); erro != nil {
			return nil, errors.New("erro ao escanear os usuários")
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func BuscandoUMUsuario(ID int) (dados.Usuario, error) {

	db, erro := ConectandoNoBanco()
	if erro != nil {
		return dados.Usuario{}, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuario where id = ?", ID)
	if erro != nil {
		return dados.Usuario{}, erro
	}
	defer linhas.Close()

	var usuarioencontrado dados.Usuario
	if linhas.Next() {
		erro := linhas.Scan(&usuarioencontrado.ID, &usuarioencontrado.Nome)
		if erro != nil {
			return dados.Usuario{}, errors.New("erro ao buscar o usuário")
		}
	}
	return usuarioencontrado, nil
}

func AlterarEstoque(ID int, estoque int) error {

	db, erro := ConectandoNoBanco()
	if erro != nil {
		return erro
	}
	defer db.Close()

	statement, erro := db.Prepare("Update livro_cadastrado set estoque = ? where id = ?")
	if erro != nil {
		return errors.New("erro ao buscar os livros")
	}
	defer statement.Close()

	if _, erro := statement.Exec(estoque, ID); erro != nil {
		return errors.New("erro ao executar o statement")
	}
	return nil
}
