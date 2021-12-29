package src

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

//Conexion con base de datos
func GetConexion()(conexion *sql.DB) {
	Driver:="mysql"
	Usuario:="root"
	Contrasenia:=""
	Nombre:="test"

	bde, err:= sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)

	if err != nil {
		panic(err.Error)
	}

	return bde
}
