package main

type IPattern interface {
	Work()

	Name() string

	Number() int
	SetNumber(int)
}
