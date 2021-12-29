package src
import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
}

func (p Persona) Saludar() {
	fmt.Printf("Hola, mi nombre es %s %s \n",
		p.Nombre, p.Apellido)
}

