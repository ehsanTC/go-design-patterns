package main

type IPattern interface {
	Work()
	Name() string
	GetNumber() int
}
