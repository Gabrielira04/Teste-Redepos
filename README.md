# Teste-Redepos
 This repository contains the evaluative test for Redepos.


Para a criação das rotas utilizei o framework Gorilla-Mux.

Rotas:

router := mux.NewRouter()
	router.HandleFunc("/Usuario", servidor.CriarUsuario).Methods(http.MethodPost) - Rota para criação
	router.HandleFunc("/Usuario", servidor.AtualizarUsuario).Methods(http.MethodPut) - Rota para update
	router.HandleFunc("/Usuario/{Cpf}", servidor.ApagarUsuario).Methods(http.MethodDelete) - Rota para Delete
	router.HandleFunc("/Usuario", servidor.LocalizarUsuarios).Methods(http.MethodGet) - Rota para busca em massa
	router.HandleFunc("/Usuario/{Cpf}", servidor.LocalizarUsuario).Methods(http.MethodGet) Rota para busca unitária

O banco de dados rodando em localhost/3306.

Os scripts para criação do banco estão dentro do projeto.
