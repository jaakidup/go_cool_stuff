package main

// Abilities ...
type Abilities interface {
	Do()
	Stop()
}

// Run ...
type Run struct {
	running bool
}

// Do ...
func (run *Run) Do() {
	println("running")
	run.running = true
}

// Stop ...
func (run *Run) Stop() {
	run.running = false
	println("stopped running")
}

// Walk ...
type Walk struct {
	walking bool
}

// Do ...
func (walk *Walk) Do() {
	println("walking")
	walk.walking = true
}

// Stop ...
func (walk *Walk) Stop() {
	walk.walking = false
	println("stopped walking")
}
