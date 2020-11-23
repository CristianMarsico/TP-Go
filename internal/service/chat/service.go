/*
VIDEO_2
-------
CREACION CARPETA SERVICE/CHAT/service.go
creacion de la estructura, servicio privado
---------------------------------------------
*/
package chat

import (
	"github.com/CristianMarsico/TP-Go/internal/config"
	"github.com/jmoiron/sqlx"
)

// Message ...
type Message struct {
	ID   int64
	Text string
}

// Service ...
type Service interface {
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil //instancia de la estructura que respeta la interfaz
}

// AddMessage ...
func (s service) AddMessage(m Message) error {
	return nil
}

// FindByID ...
func (s service) FindByID(ID int) *Message {
	return nil

}

// FindAll ...
func (s service) FindAll() []*Message {
	var list []*Message
	//list = append(list, &Message{0, "Hello World"})
	s.db.Select(&list, "SELECT * FROM messages") //lo aloja en lugar de memoria de list
	return list

}
