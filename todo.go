package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Tittle      string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(tittle string) {
	todo := Todo{
		Tittle:      tittle,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := (*todos)[index]

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}
