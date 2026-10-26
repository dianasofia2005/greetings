package greetings

import (
	"math/rand"
	"time"
)

type Alumno struct {
	Nombre  string
	Unidad1 string
	Unidad2 string
	Unidad3 string
	Unidad4 string
}

func GenerarCalificaciones(nombres []string) []Alumno {
	valores := []string{"100", "80", "70", "NA"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	alumnos := make([]Alumno, 0, len(nombres))

	for _, nombre := range nombres {
		alumno := Alumno{
			Nombre:  nombre,
			Unidad1: valores[r.Intn(len(valores))],
			Unidad2: valores[r.Intn(len(valores))],
			Unidad3: valores[r.Intn(len(valores))],
			Unidad4: valores[r.Intn(len(valores))],
		}
		alumnos = append(alumnos, alumno)
	}

	return alumnos
}
