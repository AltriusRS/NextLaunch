package screens

import (
	"Nextlaunch/src/translations"
	"crypto/rand"
	"encoding/hex"
	"time"
)

type LoadingState struct {
	ID        string
	task      string
	context   string
	working   bool
	progress  int
	timescale time.Duration
}

func JobID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func NewLoadingState(task string, timescale time.Duration) *LoadingState {
	return &LoadingState{
		ID:        JobID(),
		task:      task,
		working:   false,
		progress:  0,
		timescale: timescale,
	}
}

func (state *LoadingState) SetWorking(working bool) *LoadingState {
	state.working = working
	return state
}

func (state *LoadingState) IncrementProgress() *LoadingState {
	state.progress++
	return state
}

func (state *LoadingState) GetProgress() int {
	return state.progress
}

func (state *LoadingState) GetTask() string {
	return state.task
}

func (state *LoadingState) IsWorking() bool {
	return state.working
}

func (state *LoadingState) GetTimescale() time.Duration {
	return state.timescale
}

func (state *LoadingState) SetTimescale(timescale time.Duration) *LoadingState {
	state.timescale = timescale
	return state
}

func (state *LoadingState) SetProgress(progress int) *LoadingState {
	state.progress = progress

	if progress >= 100 {
		state.working = false
		state.context = translations.Manager.Translate(state.context)
	}

	return state
}

func (state *LoadingState) SetContext(context string) *LoadingState {
	state.context = context
	return state
}

func (state *LoadingState) GetContext() string {
	return state.context
}
