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
	"time"

	"github.com/CristianMarsico/TP-Go/internal/config"
	"github.com/CristianMarsico/TP-Go/internal/database"
	"github.com/CristianMarsico/TP-Go/internal/service/chat"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDataBase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	/*if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}*/

	/*fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}*/
	service, _ := chat.New(db, cfg)
	httpService := chat.NewHTTPTrasnport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()

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

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS messages (
        id integer primary key autoincrement,
        text varchar);`

	//execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}
