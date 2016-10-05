package main

import "fmt"

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
  l := *lista
  removido := l[indice]
  *lista = append(l[0 : indice], l[indice+1 : ]...)
  return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
  return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
  return lista.RemoverIndice(len(*lista) - 1)
}

func main() {
  lista := ListaGenerica{1, "Cafe", 42, true, 23, "Bola", 3.14, false}

  fmt.Printf("Lista original:\n%v\n", lista)

  fmt.Printf("Removendo do inicio: %v, apos remocao:\n%v\n", lista.RemoverInicio(), lista)
  fmt.Printf("Removendo do fim: %v, apos remocao:\n%v\n", lista.RemoverFim(), lista)
  fmt.Printf("Removendo do indice 3: %v, apos remocao:\n%v\n", lista.RemoverIndice(3), lista)
  fmt.Printf("Removendo do indice 0: %v, apos remocao:\n%v\n", lista.RemoverIndice(0), lista)
  fmt.Printf("Removendo o ultimo indice: %v, apos remocao:\n%v\n", lista.RemoverIndice(len(lista) - 1), lista)
}

// go run src/cap5-lista-generica/lista.go
