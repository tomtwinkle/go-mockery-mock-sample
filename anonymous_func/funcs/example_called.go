package funcs

import (
	"github.com/pkg/errors"
	"log"
)

type ExampleCalled interface {
	WithFunc(func() (interface{}, error)) (interface{}, error)
	Example() (string, error)
}
type exampleCalled struct{}

func NewExampleCalled() ExampleCalled {
	return &exampleCalled{}
}

func (e *exampleCalled) WithFunc(
	fn func() (interface{}, error),
) (interface{}, error) {
	log.Print("called with func!")
	v, err := fn()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return v, nil
}

func (e *exampleCalled) Example() (string, error) {
	log.Print("called example!")
	return "example!!!", nil
}
