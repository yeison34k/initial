package src
import (
	"net/http"
	"github.com/gorilla/mux"
	"log"

)

func GetPathParameters(str string, r *http.Request) string{
	vars := mux.Vars(r)
    param, ok := vars[str]
    if !ok {
        log.Println("is missing in parameters")
    }

	return  string(param)
}