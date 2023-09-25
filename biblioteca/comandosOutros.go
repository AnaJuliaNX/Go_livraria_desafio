package biblioteca

import (
	"errors"

	"github.com/AnaJuliaNX/desafio2/dados"
)

// função com finalidade de reduzi a repetição do mesmo comando em vários arquivos
// Essa função vai buscar os dados de todos os usuários que tenho no banco  e trazer para mim
func BuscandoOSUsuarios() ([]dados.Usuario, error) {

	//Função que faz a conexão com o banco de dados (mais detalhes no arquivo "comandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	//Busco por todos os usuários que tenho salvo no banco
	linhas, erro := db.Query("select * from usuario")
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

func BuscarDadosEmprestimo() ([]dados.EmprestimoDevolucao, error) {

	db, erro := ConectandoNoBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select * from emprestimo_devolucao")
	if erro != nil {
		return nil, errors.New("erro ao buscar os dados dos empréstimos")
	}
	defer linhas.Close()

	var dadosEmprestimos []dados.EmprestimoDevolucao
	for linhas.Next() {
		var dadosEmprestimo dados.EmprestimoDevolucao

		if erro := linhas.Scan(&dadosEmprestimo.Nome_Usuario, &dadosEmprestimo.Titulo_livro, &dadosEmprestimo.Data_Emprestimo, &dadosEmprestimo.Data_Devolucao); erro != nil {
			return nil, errors.New("erro ao escanear os dados de emprestimo")
		}

		dadosEmprestimos = append(dadosEmprestimos, dadosEmprestimo)
	}
	return dadosEmprestimos, nil
}
