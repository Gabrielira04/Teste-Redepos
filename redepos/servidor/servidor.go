package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuarios []struct {
	Name        string `json:"Name"`
	Cpf         string `json:"CPF"`
	PhoneNumber string `json:"PhoneNumber"`
	CelNumber   string `json:"CelNumber"`
	FaxNumber   string `json:"FaxNumber"`
	Email       string `json:"email"`
	Cep         string `json:"cep,omitempty"`
}

// inserindo usuarios no BD
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("falha na leitura da requisição"))
		return
	}

	var usuario usuarios

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario"))
		return
	}

	//validando os campos

	for _, usu := range usuario {

		if usu.Name == "" || usu.Cpf == "" || usu.CelNumber == "" || usu.Email == "" {

			fmt.Printf("OS campos obrigatorios devem ser preenchidos")
			w.Write([]byte("OS campos obrigatorios devem ser preenchidos"))
			return

		}
	}

	// Aqui deixei comentado um metodo de validação por cada campo.

	/*for _, usu := range usuario {

		if usu.Cpf == "" {

			fmt.Printf("cpf nao pode ser vazio\n")
			//w.Write([]byte("cpf nao pode ser vazio"))
			w.Write([]byte("O CPF é obrigatorio\n"))
			return

		}

	}

	for _, usu := range usuario {

		if usu.CelNumber == "" {

			fmt.Printf("O número de celular é obrigatorio\n")
			w.Write([]byte("O número de celular é obrigatorio\n"))
			return

		}

	}

	for _, usu := range usuario {

		if usu.Email == "" {

			fmt.Printf("O Email é um campo obrigatorio")
			w.Write([]byte("O Email é um campo obrigatorio"))
			return

		}

	}*/

	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Println("erro")
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return

	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios(Name, Cpf, PhoneNumber, CelNumber, FaxNumber, Email, Cep) values (?, ?, ?, ?, ?, ?, ?)")
	if erro != nil {
		fmt.Println("erro")
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	for _, v := range usuario {

		_, erro := statement.Exec(v.Name, v.Cpf, v.PhoneNumber, " ", v.FaxNumber, v.Email, v.Cep)
		if erro != nil {
			fmt.Println("erro statem")
			fmt.Printf("%#v \n", erro)
			w.Write([]byte("Erro ao executar o statement"))
			return
		}
	}

	// idInserido, erro := insercao.LastInsertId()
	// if erro != nil {
	// 	fmt.Printf("%#v \n", erro)
	// 	w.Write([]byte("Erro ao buscar o cpf inserido"))
	// 	return
	// }
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! %d", 0)))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("falha na leitura da requisição de atualização"))
		return
	}

	var usuario usuarios

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario"))
		return
	}

	//TODO validar os campos

	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Println("erro")
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return

	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set Name = ? where Cpf = ?")
	if erro != nil {
		fmt.Println("erro")
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	for _, v := range usuario {

		_, erro := statement.Exec(v.Name, v.Cpf)
		if erro != nil {
			fmt.Printf("%#v \n", erro)
			w.Write([]byte("Erro ao executar atualizção"))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario(s) atualizados com sucesso! %d", 0)))
}

func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("falha na leitura da requisição de deletar usuario"))
		return
	}

	var usuario usuarios

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return

	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE from usuarios where CPF = ?") // Para deletar um usuario é necessário digitar o cpf do mesmo.
	if erro != nil {
		fmt.Println("erro")
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro na solicitação de deletar usuario(s)"))
		return
	}
	defer statement.Close()

	for _, v := range usuario {

		_, erro := statement.Exec(v.Cpf)
		if erro != nil {
			fmt.Println("erro statem")
			fmt.Printf("%#v \n", erro)
			w.Write([]byte("Erro ao deletar usuario(s)"))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario(s) deletados com sucesso! %d", 0)))
}

//buscar usuarios - vai trazer todos os usuarios

func LocalizarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
	}
	defer db.Close()

	//comando para pegar todos os usuarios salvos no banco

	linhas, erro := db.Query("SELECT * from usuarios") //comando para buscar a tabela no banco
	if erro != nil {
		fmt.Printf("%#v \n", erro) //controle de erros
		w.Write([]byte("Erro ao buscar os usuarios no banco"))
		return
	}
	defer linhas.Close()

	var usuario []usuarios
	for linhas.Next() {
		var usuario usuarios
		for _, v := range usuario {
			if erro := linhas.Scan(&v.Name, &v.Cpf, &v.PhoneNumber, &v.CelNumber, &v.FaxNumber, &v.Email, &v.Cep); erro != nil {
				fmt.Printf("%#v \n", erro)
				w.Write([]byte("Erro ao scanear o usuario"))

				return
			}

			usuario = append(usuario, v)
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao convertet os usuarios para json"))
		return
	}
}

//Função para buscar um usuário especifico, buscando pelo Cpf

func LocalizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	Cpf, erro := strconv.ParseUint(parametros["Cpf"], 10, 64) // O cpf para busca deve ser digitado sem pontuações apenas numeros
	if erro != nil {
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao convertet o usuarios para json"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao convertet o usuarios para json"))
		return
	}
	linha, erro := db.Query("select * from usuarios where Cpf = ?", Cpf)
	if erro != nil {
		fmt.Printf("%#v \n", erro)
		w.Write([]byte("Erro ao buscar usuario"))
		return
	}

	var usuario usuarios
	if linha.Next() {
		for _, v := range usuario {
			fmt.Printf("%#v \n", erro)
			if erro := linha.Scan(v.Name, v.Cpf, v.PhoneNumber, v.CelNumber, v.FaxNumber, v.Email, v.Cep); erro != nil {
				w.Write([]byte("erro ao escanear usuario"))
				return
			}

		}

		if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
			fmt.Printf("%#v \n", erro)
			w.Write([]byte("erro ao converte usuario para json"))
			return
		}
	}
}
