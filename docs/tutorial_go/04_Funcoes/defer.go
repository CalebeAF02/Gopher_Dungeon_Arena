package goguide

import (
	"fmt"
	"os"
)

// ReadFile demonstra o uso de defer para garantir que o arquivo seja fechado
// quando a função terminar, ainda que ocorra um erro ao ler.
func ReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
		fmt.Println("arquivo fechado")
	}()

	// Aqui você faria a leitura real do arquivo.
	return nil
}

// DeferExample mostra que o código colocado em defer é executado no fim da função.
func DeferExample() {
	fmt.Println("antes")
	defer fmt.Println("depois")
	fmt.Println("durante")
}
