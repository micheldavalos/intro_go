package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	nombre string
	edad uint64
}

type Bynombre []Persona

func (a Bynombre) Len() int           { return len(a) }
func (a Bynombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Bynombre) Less(i, j int) bool { return a[i].nombre < a[j].nombre }

type ByEdad []Persona

func (a ByEdad) Len() int           { return len(a) }
func (a ByEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEdad) Less(i, j int) bool { return a[i].edad < a[j].edad }

func main()  {
	ps := []Persona{
		Persona{nombre: "michel", edad: 36},
		Persona{nombre: "davalos", edad: 20},
		Persona{nombre: "boites", edad: 50},
	}

	sort.Sort(Bynombre(ps))
	fmt.Println(ps)

	sort.Sort(ByEdad(ps))
	fmt.Println(ps)

}