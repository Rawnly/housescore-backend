package main //server
import (
	"fmt"
	"github.com/gorilla/mux"
	"housescore/configuration"
	"housescore/controllers/users_controller"
	"housescore/database"
	"housescore/database/models"
	"housescore/util"
	"log"
	"net/http"
)

func main() {
	env := util.Env{}
	port := env.Get("port", "8080")

	// Configurazione con valori prederifiniti
	configuration.Default()

	// Sovrascrive la configurazione predefinita se possibile
	configuration.Init()

	// Connessione al database (banca di dati)
	db := database.Init()

	db.AutoMigrate(&models.User{})

	sql, err := db.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer sql.Close()


	r := mux.NewRouter() //crei un router, ti indica la strada

	// Quando arriva una POST su "/auth/login" esegui `users_controller.Login`
	POST(r, "/auth/signup", users_controller.SignUp)

	r.HandleFunc("/ciao/{name}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request) //variaibili nell url
		writer.WriteHeader(http.StatusOK)
		token := request.URL.Query().Get("token")
		fmt.Print(token)

		_, _ = fmt.Fprintf(writer, "ciao %s", vars["name"]) // quello che stampi
	})

	log.Printf("Server running at http://localhost:%s", port)

	//dopo lo / si occupa r e sei in ascolto sulla porta
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}


// METODI
func GET(r *mux.Router, path string, h http.HandlerFunc)  {
	r.HandleFunc(path, h).Methods("GET")
}

func POST(r *mux.Router, path string, h http.HandlerFunc)  {
	r.HandleFunc(path, h).Methods("POST")
}

func PUT(r *mux.Router, path string, h http.HandlerFunc)  {
	r.HandleFunc(path, h).Methods("PUT", "PATCH")
}

func DELETE(r *mux.Router, path string, h http.HandlerFunc)  {
	r.HandleFunc(path, h).Methods("DELETE")
}

/**
terminale:
go mod init : "inizializzi il progetto"
go get -u github.com/gorilla/mux : "scarichi qualsiasi libreria"


*/
