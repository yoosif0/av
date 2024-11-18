package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
}

type msg string

const (
	increment msg = "increment"
	decrement msg = "decrement"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg msg) (model, tea.Cmd) {
	switch msg {
	case increment:
		m.count++
	case decrement:
		m.count--
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Count: %d", m.count)
}

