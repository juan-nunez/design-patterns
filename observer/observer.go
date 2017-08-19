package main

import (
    "log"
)

type Observer interface {
    notify(event Event)
}


type Observable interface {
    registerObserver(o Observer)
    unregisterObserver(o Observer)
}


type Event interface {
    getName() string
}


type StartEvent struct{}
type DoneEvent struct {}
func (s StartEvent) getName() string {
    return "TRANSACTION_STARTED"
}

func (d DoneEvent) getName() string {
    return "TRANSACTION_DONE"
}



type BankTransaction struct {
    observers []Observer
}

func newBankTransaction() *BankTransaction{
    observers := make([]Observer,0)
    return &BankTransaction{observers}
}

func (b *BankTransaction) registerObserver(o Observer) {
    b.observers = append(b.observers, o)  
}

func (b *BankTransaction) unregisterObserver(o Observer) {

}

func (b *BankTransaction) do() {
    b.start()
    log.Println("Doing transaction work")
    b.done()
}

func (b *BankTransaction) start() {
    startEvent := StartEvent{}
    for _, observer := range b.observers {
        observer.notify(startEvent)
    }
}

func (b *BankTransaction) done() {
    doneEvent := DoneEvent{}
    for _, observer := range b.observers {
        observer.notify(doneEvent)
    }
}

type Service struct{}

func (s Service) notify(event Event) {
    log.Printf("Received event: %s", event.getName())
}


func main() {
    s := Service{}
    b := newBankTransaction()
    b.registerObserver(s)
    b.do()

}
