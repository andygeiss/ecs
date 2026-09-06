package ecs

import "context"

const (
	// StateEngineContinue tells the engine to keep going after this System.
	StateEngineContinue = 0
	// StateEngineStop tells the engine to stop after this System.
	StateEngineStop = 1
)

// Engine drives the systems through their three stages: Setup once, Process
// over and over, Teardown once.
type Engine interface {
	// Run calls Process on each System over and over, until a System returns
	// StateEngineStop or ctx is cancelled. It does not sleep between passes,
	// so a system paces the loop.
	Run(ctx context.Context)
	// Setup calls Setup on each System, once, before the first Process.
	Setup()
	// Teardown calls Teardown on each System, once, after the last Process.
	Teardown()
	// Tick calls Process on each System exactly once, stopping early when a
	// System returns StateEngineStop.
	Tick()
}
