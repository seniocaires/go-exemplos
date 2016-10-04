package main

import "fmt"

func main() {
  estados := make(map[string]Estado, 6)

  estados["GO"] = Estado{"Goias", 6434052, "Goiania"}
  estados["PB"] = Estado{"Paraiba", 3914418, "Joao Pessoa"}
  estados["PR"] = Estado{"Parana", 10997462, "Curitiba"}
  estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
  estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
  estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

  fmt.Println(estados)
}

type Estado struct {
  nome string
  populacao int
  capital string
}

// go run src/cap4-estados/estados.go
