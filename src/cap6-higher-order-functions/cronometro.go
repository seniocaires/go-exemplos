package main

import (
  "fmt"
  "time"
)

func GerarFibonacci(n int) func() {
  return func() {
    a, b := 0, 1

    fib := func() int {
      a, b = b, a + b

      return a
    }

    for i := 0 ; i < n ; i++ {
      fmt.Printf("%d ", fib())
    }
  }
}

func Cronometrar(funcao func()) {
  inicio := time.Now()

  funcao()

  fmt.Printf("\nTempo de execucao: %s\n\n", time.Since(inicio))
}

func main() {
  Cronometrar(GerarFibonacci(8))
  Cronometrar(GerarFibonacci(48))
  Cronometrar(GerarFibonacci(88))
}

// go run src/cap6-higher-order-functions/cronometro.go
