package funcs

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-mockery-mock-sample/anonymous_func/mocks"
	"testing"
)

func TestExample_ExampleAnonymousFunc(t *testing.T) {
	t.Run("example NG", func(t *testing.T) {
		mockExampleCalled := new(mocks.ExampleCalled)

		mockExampleCalled.On("WithFunc", mock.Anything).Return("WithFunc mock!", nil)
		mockExampleCalled.On("Example").Return("Example mock!", nil)

		e := NewExample(mockExampleCalled)
		actual, err := e.ExampleAnonymousFunc()
		assert.NoError(t, err)
		assert.Equal(t, "WithFunc mock!", actual)
		mockExampleCalled.AssertExpectations(t)
	})

	t.Run("example OK", func(t *testing.T) {
		mockExampleCalled := new(mocks.ExampleCalled)

		mockWithFunc := mockExampleCalled.On("WithFunc", mock.Anything)
		mockWithFunc.Run(func(args mock.Arguments) {
			// WithFuncの処理をmockしつつ、高階関数の引数の関数実行は行う
			fn := args[0].(func() (interface{}, error))
			v, err := fn()
			// 高階関数の引数の関数の戻り値を返す
			mockWithFunc.ReturnArguments = mock.Arguments{v, err}
		})
		mockExampleCalled.On("Example").Return("Example mock!", nil)

		e := NewExample(mockExampleCalled)
		actual, err := e.ExampleAnonymousFunc()
		assert.NoError(t, err)
		assert.Equal(t, "Example mock!", actual)
		mockExampleCalled.AssertExpectations(t)
	})
}
