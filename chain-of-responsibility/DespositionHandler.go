package main

type DepositionHandler interface {
	SetNext(handler DepositionHandler) DepositionHandler
	Handle(deposit int) error
}
