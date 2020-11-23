/*
VIDEO_2
-------
CREACION CARPETA SERVICE/CHAT/service.go
creacion de la estructura, servicio privado
---------------------------------------------
*/
package chat

import "github.com/CristianMarsico/TP-Go/internal/config"

// Message ...
type Message struct {
	ID   int64
	Text string
}

// ChatService ...
type ChatService interface {
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}
type service struct {
	conf *config.Config
}

// New ...
func New(c *config.Config) (ChatService, error) {
	return service{c}, nil //instancia de la estructura que respeta la interfaz
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
	list = append(list, &Message{0, "Hello World"})
	return list

}
