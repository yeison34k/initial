package src

import (
	"net/http"
	"log"
)

func HttpCommonServer()  {
	http.HandleFunc("/", Index)
	http.HandleFunc("/test", Test)
	http.HandleFunc("/add-user", AddUser) //route
	http.HandleFunc("/insert", InserUser) //Action Insert
	http.HandleFunc("/edit", EditUser) //route
	http.HandleFunc("/update", UdapteUser) //Action update
	http.HandleFunc("/delete", DeleteUser) //Action delete
	log.Println("Ejecutando servidor")
	log.Fatal(http.ListenAndServe(":9090", nil))
}


	