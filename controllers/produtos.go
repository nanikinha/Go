package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"br.com.industrial/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	BuscaTodosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", BuscaTodosOsProdutos)
}
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)

		}

		quantidadeConvertida, err := strconv.ParseInt(quantidade, 10, 0)
		if err != nil {
			log.Println("Erro na conversão da quantidade")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)

		}
		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProdutoStr := r.URL.Query().Get("id")
	idDoProduto, err := strconv.ParseInt(idDoProdutoStr, 10, 0)
	if err != nil {
		log.Println("Erro ao converter inteiro")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProdutoStr := r.URL.Query().Get("id")
	idDoProduto, err := strconv.ParseInt(idDoProdutoStr, 10, 0)
	if err != nil {
		log.Println("Erro ao converter inteiro")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		nome := r.FormValue("nome")

		idDoProduto, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			log.Println("Erro ao converter inteiro")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)

		}

		quantidadeConvertida, err := strconv.ParseInt(quantidade, 10, 0)
		if err != nil {
			log.Println("Erro na conversão da quantidade")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)

		}
		models.AtualizaProduto(idDoProduto, nome, descricao, precoConvertido, quantidadeConvertida)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}
