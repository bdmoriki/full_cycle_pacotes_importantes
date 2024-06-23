package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"numero"`
	Saldo  float64 `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100.0}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonPuro := []byte(`{"numero": 2, "saldo": 200.55}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Numero: %d, Saldo: %f\n", contaX.Numero, contaX.Saldo)
}
