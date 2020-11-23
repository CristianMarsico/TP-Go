/**
VIDEO_1
--------
EJECUCIONES
	go run cmd/chat/charsrv.go
	go run cmd/chat/charsrv.go -config ./config/config.yaml

PASOS P/ INSTALACION DE DEPENDENCIAS
	4) touch config.yaml (crea archivo 'config.yaml', despues hay q crear carpeta 'config' y moverlo ahi)
-------------------------------------
VIDEO_2
-------
	se modifican algunos codigos, impletenta serv.
	no hay nuevas ejecuciones
	se sigue usando go run cmd/chat/charsrv.go -config ./config/config.yaml
--------------------------------------
VIDEO_3
-------
creacion DATABASE/db.go en INTERNAL
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CristianMarsico/TP-Go/internal/config"
	"github.com/CristianMarsico/TP-Go/internal/database"
	"github.com/CristianMarsico/TP-Go/internal/service/chat"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDataBase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	/*fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)*/
	service, _ := chat.New(db, cfg)
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}

}

func readConfig() *config.Config {

	//para leer param
	configFile := flag.String("config", "./config.yaml", "this is service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
