package main

import (
  "fmt"
  "math"
  "sync"
  "time"
)

func calcular(base float64, controle *sync.WaitGroup) {
  defer controle.Done()
  n := 0.0

  for i := 0 ; i < 100000000 ; i++ {
    n += base / math.Pi * math.Sin(2)
  }

  fmt.Println(n)
}

func main() {
  inicio := time.Now()
  var controle sync.WaitGroup
  controle.Add(3)

  go calcular(9.37, &controle)
  go calcular(6.94, &controle)
  go calcular(42.57, &controle)

  controle.Wait()
  fmt.Printf("Finalizando em %s.\n", time.Since(inicio))
}

// go run src/cap7-paralelo/paralelo.go
// GOMAXPROCS=2 go run src/cap7-paralelo/paralelo.go
// GOMAXPROCS=3 go run src/cap7-paralelo/paralelo.go
