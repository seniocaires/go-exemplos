package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {

  entrada := os.Args[1:]
  numeros := make([]int, len(entrada))

  for i, n := range entrada {
    numero, err := strconv.Atoi(n)
    if err != nil {
      fmt.Printf("%s nao e um numero valido!", n)
      os.Exit(1)
    }
    numeros[i] = numero
  }

  fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
  if len(numeros) <= 1 {
    return numeros
  }

  n := make([]int, len(numeros))
  copy(n, numeros)

  indicePivo := len(n) / 2
  pivo := n[indicePivo]
  n = append(n[ : indicePivo], n[indicePivo + 1 : ] ...)

  menores, maiores := particionar(n, pivo)

  return append(append(quicksort(menores), pivo), quicksort(maiores) ...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
  for _, n := range numeros {
    if n <= pivo {
      menores = append(menores, n)
    } else {
      maiores = append(maiores, n)
    }
  }

  return menores, maiores
}

// go run src/cap2-quicksort/quicksort.go 10 30 41 53 78 12 19 22
