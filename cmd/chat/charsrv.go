/**
EJECUCIONES
	go run cmd/chat/charsrv.go
	go run cmd/chat/charsrv.go -config ./config/config.yaml

PASOS P/ INSTALACION DE DEPENDENCIAS
	4) touch config.yaml (crea archivo 'config.yaml', despues hay q crear carpeta 'config' y moverlo ahi)
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CristianMarsico/TP-Go/internal/config"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)
}
