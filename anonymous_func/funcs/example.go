package funcs

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type Example interface {
	ExampleAnonymousFunc() (string, error)
}

type example struct {
	called ExampleCalled
}

func NewExample(called ExampleCalled) Example {
	return &example{called: called}
}

func (e *example) ExampleAnonymousFunc() (string, error) {
	log.Print("start ExampleAnonymousFunc()")
	v, err := e.called.WithFunc(func() (i interface{}, err error) {
		v, err := e.called.Example()
		log.Print(fmt.Sprintf("called anonymous func! result: %s", v))
		return v, errors.WithStack(err)
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	log.Print(fmt.Sprintf("end ExampleAnonymousFunc() result: %s", v))
	return v.(string), nil
}
