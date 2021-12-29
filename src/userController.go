package src

import (
	"net/http"
	
	"text/template"
)
var deleteSql = "DELETE FROM user WHERE id=?"
var inserUser = "INSERT INTO user(id, name, lastname) VALUES (?, ?, ?)"
var getAllUser = "SELECT * FROM user"
var getUser = "SELECT * FROM user WHERE id=?"
var updateUser = "UPDATE user SET name=?, lastname=? WHERE id=?"

var redirect301 = 301

var plnatilla = template.Must(template.ParseGlob("plantillas/*"))
var bdco = GetConexion()

func Index(w http.ResponseWriter, r *http.Request) {
	
	registros, err:= bdco.Query(getAllUser)
	if err != nil {	panic(err.Error()) }

	user := User{}
	userList := []User{}

	for registros.Next() {
		var id int
		var name, lastname string
		err = registros.Scan(&id, &name, &lastname)

		if err != nil {	panic(err.Error()) }

		user.Id = id
		user.Name = name
		user.Lastname = lastname
		userList=append(userList, user)
	}
	
	plnatilla.ExecuteTemplate(w, "inicio", userList)
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("manehando rutas en go"))
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	plnatilla.ExecuteTemplate(w, "addUser", nil)
}


func EditUser(w http.ResponseWriter, r *http.Request) {
	idparam := r.URL.Query().Get("id")
	registro, err:= bdco.Query(getUser, idparam)
	
	if err != nil {	panic(err.Error()) }
	
	user := User{}
	
	for registro.Next() {
		var id int
		var name, lastname string
		err = registro.Scan(&id, &name, &lastname)

		if err != nil {	panic(err.Error()) }

		user.Id = id
		user.Name = name
		user.Lastname = lastname
	}
	plnatilla.ExecuteTemplate(w, "edit", user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idparam := r.URL.Query().Get("id")
	var bdco = GetConexion()
	res, err := bdco.Prepare(deleteSql)
	if err != nil {
		panic(err.Error())
	}
	
	res.Exec(idparam)

	http.Redirect(w, r, "/", redirect301)
}

func InserUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		lastName := r.FormValue("lastname")
		
		crearregistro, err:= bdco.Prepare(inserUser)
		if err != nil {	panic(err.Error()) }
		
		crearregistro.Exec(id, name, lastName)
	
		http.Redirect(w, r, "/", redirect301)
	}
}

func UdapteUser(w http.ResponseWriter, r *http.Request) {
	idparam := r.URL.Query().Get("id")
	
	if r.Method == "POST" {
		name := r.FormValue("name")
		lastName := r.FormValue("lastname")

		var bdco = GetConexion()
		res, err := bdco.Prepare(updateUser)
		if err != nil {
			panic(err.Error())
		}
		
		res.Exec(name, lastName, idparam)

		http.Redirect(w, r, "/", redirect301)
	}

}