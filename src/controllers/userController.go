package controllers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar Usuários"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

func SearchById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário"))
}
