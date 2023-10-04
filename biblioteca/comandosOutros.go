package biblioteca

import (
	"errors"

	"github.com/AnaJuliaNX/desafio2/dados"
)

// função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai buscar os dados de todos os usuários que tenho no banco e trazer para mim
func BuscandoOSUsuarios(search string, offset int) ([]dados.Usuario, error) {

	//Função que faz a conexão com o banco de dados (mais detalhes no arquivo "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	//dou um limite de três usuários por página
	limit := 15
	//Busco a quantidade de usuários ordenados pelo id que o limite permite
	//e uso o offset para saber por qual id devo começar
	//Ou no caso de página 2 uso o offset para saber por qual id devo continuar
	linhas, erro := db.Query("select id, nome from usuario where nome like ? order by id limit ? offset ?", "%"+search+"%", limit, offset)
	if erro != nil {
		return nil, errors.New("erro ao buscar os usuários")
	}
	defer linhas.Close()

	//Escaneio todos os dados de todos os usuários e retorno esses dados
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

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai buscar os dados de um um usuário especifico pelo sue ID e trazer pra mim,
func buscandoUMUsuario(ID int) (dados.Usuario, error) {

	//Chamo a função que faz a conexão com o banco de dados (mais detalhes no arquivo "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return dados.Usuario{}, erro
	}
	defer db.Close()

	//Busco pelo ID de um usuário especifico
	linhas, erro := db.Query("select * from usuario where id = ?", ID)
	if erro != nil {
		return dados.Usuario{}, erro
	}
	defer linhas.Close()

	//Escaneio os dados desse usuário que foi buscado e retorno os dados dele
	var usuarioencontrado dados.Usuario
	if linhas.Next() {
		erro := linhas.Scan(&usuarioencontrado.ID, &usuarioencontrado.Nome)
		if erro != nil {
			return dados.Usuario{}, errors.New("erro ao buscar o usuário")
		}
	}
	return usuarioencontrado, nil
}

// Função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai fazer a alteração no estoque do livro sempre que eue emprestar ou devolver um
func AlterarEstoque(ID int, estoque int) error {

	//Chamo a função que faz a conexão com o banco de dados (mais detalhes no arquivo "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return erro
	}
	defer db.Close()

	//Crio um statement que vai buscar o estoque de um livro especificado pelo Id
	statement, erro := db.Prepare("Update livro_cadastrado set estoque = ? where id = ?")
	if erro != nil {
		return errors.New("erro ao buscar os livros")
	}
	defer statement.Close()
	//Executo o statement que vai alterar o estoque do livro especificado acima
	if _, erro := statement.Exec(estoque, ID); erro != nil {
		return errors.New("erro ao executar o statement")
	}
	return nil
}
