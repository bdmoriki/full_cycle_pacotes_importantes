package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// Gravação em string
	//	tamanho, err := f.WriteString("Hello, World!")

	// Gravação em bytes
	tamanho, err := f.Write([]byte("Hello, World!"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso %d bytes \n", tamanho)

	f.Close()

	// Leitura
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	arquivo2.Close()

	err = os.Remove("arquivo.txt")

	if err != nil {
		panic(err)
	}
}
