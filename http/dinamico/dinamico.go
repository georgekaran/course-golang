package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hora-certa" {
		paginaNaoEncontrada(w, r)
		return
	}
	// Detalhe: O sistema de formato de hora do Go utiliza uma nomenclatura diferente do padrão ("dd/MM/yyyy" por exemplo)
	// Cada número representa uma parte de um tempo.
	// Por exemplo 02 - Dia, 01 - Mês e assim por diante...
	// Para mais detalhes, https://programming.guide/go/format-parse-string-time-date-example.html
	s := time.Now().Format("02/01/2006 03:04:05")
	// Fprint - Recebe um writter e escreve nele
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func paginaNaoEncontrada(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Página não encontrada!</h1>
<a href="/hora-certa">Voltar para hora certa</a>`)
}

func main() {
	http.HandleFunc("/hora-certa", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
