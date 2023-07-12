package chain_of_responsibility

type DepositionHandler interface {
	SetNext(handler DepositionHandler) DepositionHandler
	Handle(deposit int) error
}
