package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}
type Observer interface {
	getId() string
	updateValue(string)
}
type Item struct {
	observer  []Observer
	name      string
	available bool
}

func newItem(name string) *Item {
	return &Item{name: name}
}
func (i *Item) UpdateAvailable() {
	fmt.Println("Item ", i.name, " is available")
	i.available = true
	i.broadcast()
}
func (i *Item) broadcast()  {
	for _, observer := range i.observer {
		observer.updateValue(i.name)
	}
}
func (i *Item) register(observer Observer)  {
	i.observer = append(i.observer, observer)
}

type EmailClient struct {
	id string
}
func (e *EmailClient) getId() string {
	return e.id
}

func (e *EmailClient) updateValue(name string)  {
	fmt.Println("EmailClient ", e.id, " got notified that ", name, " is available")
}
func main() {
	item := newItem("iPhone")
	firstClient := &EmailClient{id: "1"}
	secondClient := &EmailClient{id: "2"}
	item.register(firstClient)
	item.register(secondClient)
	item.UpdateAvailable()
}