package main

type Event int64

const (
	KeyloggerActivate Event = iota
	KeyloggerDeactivate
	Quit
)
