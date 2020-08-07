package routes

import (
	"net/http"

	"github.com/grbalmeida/curso-golang-web-alura/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
