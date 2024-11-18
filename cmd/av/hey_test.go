package main

import (
	"testing"
  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
)

func TestAdopt(t *testing.T) {
	initialModel := model{count: 0}

	t.Run("increment", func(t *testing.T) {
		m, _ := initialModel.Update(increment)
		if m.count != 1 {
			t.Errorf("expected count %d, got %d", 1, m.count)
		}
		m, _ = m.Update(increment)
		if m.count != 2 {
			t.Errorf("expected count %d, got %d", 2, m.count)
		}
		m, _ = m.Update(decrement)
		if m.count != 1 {
			t.Errorf("expected count %d, got %d", 1, m.count)
		}
	})
}


type MyMockedObject struct{
  mock.Mock
	stackAdoptViewModel
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {

  args := m.Called(number)
  return args.Bool(0), args.Error(1)

}

func TestSomething(t *testing.T) {

  testObj := new(MyMockedObject)
	testObj.stackAdoptViewModel.adoptionComplete = true
  testObj.On("DoSomething", 123).Return(true, nil)

	s := testObj.stackAdoptViewModel.View()

  // assert that the expectations were met
	expected := "                      \n                      \n  No branch is adopted\n                      \n"
  assert.Equal(t, expected, s)

}
