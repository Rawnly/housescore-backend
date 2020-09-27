package main //server
import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter() //crei un router, ti indica la strada

	r.HandleFunc("/ciao/{name}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request) //variaibili nell url
		writer.WriteHeader(http.StatusOK)
		token := request.URL.Query().Get("token")

		_, _ = fmt.Fprintf(writer, "ciao %s, token = %s", vars["name"], token) // quello che stampi
	})

	//dopo lo / si occupa r e sei in ascolto sulla porta
	log.Fatal(http.ListenAndServe(":8080", r))
}

/**
terminale:
go mod init : "inizializzi il progetto"
go get -u github.com/gorilla/mux : "scarichi qualsiasi libreria"


*/
