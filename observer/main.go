package main

import "fmt"

type Observer interface {
	update()
	getName() string
}
type BaseObserver struct {
	name string
}

type Subject struct {
	observers map[string] Observer
}

func (s *Subject) registerObserver(observer Observer) {
	s.observers[observer.getName()] = observer
}

func (s *Subject) unregisterObserver(observerName string) {
	delete(s.observers, observerName)
}

func (s *Subject) notify() {
	for _, l := range s.observers {
		if l != nil { l.update() }
	}
}

type LoudObserver struct {
	name string
}

func (l *LoudObserver) getName() string {
	return l.name
}

func (l *LoudObserver) update() {
	fmt.Printf("Observer: %s has been updated\n", l.getName())
}

func main() {
	obs1 := &LoudObserver{"obs1"}
	obs2 := &LoudObserver{"obs2"}
	obs3 := &LoudObserver{"obs3"}

	subject := &Subject{observers: map[string]Observer{}}
	subject.registerObserver(obs1)
	subject.registerObserver(obs2)
	subject.registerObserver(obs3)

	subject.notify()

	subject.unregisterObserver(obs2.name)

	subject.notify()
}