package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type State int

const (
	Happy State = iota
	Hungry
	Sad
	Sleepy
)

type ControllableCharacter interface {
	GetState() State
	SetState(state State)
}
type Character struct {
	Mood State
}

func (c *Character) GetState() State {
	return c.Mood
}
func (c *Character) SetState(state State) {
	c.Mood = state
}

type Game struct {
	cc ControllableCharacter
}

func (g *Game) FeedPlayer() {
	if g.cc.GetState() == Hungry {
		g.cc.SetState(Happy)
	}
}

type MockCharacter struct {
	mock.Mock
}

func (mc *MockCharacter) GetState() State {
	args := mc.Called()
	return args.Get(0).(State)
}
func (mc *MockCharacter) SetState(state State) {
	mc.Called(state)
}

func TestMakesHappyIfHungry(t *testing.T) {
	mockCharacter := new(MockCharacter)
	mockCharacter.On("GetState").Return(Hungry)
	mockCharacter.On("SetState", Happy).Return()

	game := &Game{cc: mockCharacter}
	game.FeedPlayer()

	mockCharacter.AssertNumberOfCalls(t, "GetState", 1)
	mockCharacter.AssertCalled(t, "GetState")

	mockCharacter.AssertCalled(t, "SetState", Happy)
	mockCharacter.AssertNumberOfCalls(t, "SetState", 1)

	mockCharacter.AssertExpectations(t)
}

func TestDoesNotMakeHappyIfNotHungry(t *testing.T) {
	mockCharacter := new(MockCharacter)
	mockCharacter.On("GetState").Return(Sleepy)

	game := &Game{cc: mockCharacter}
	game.FeedPlayer()

	mockCharacter.AssertNotCalled(t, "SetState", mock.Anything)
	mockCharacter.AssertExpectations(t)
}
