package main

import (
	"fmt"
)

type NotSupportedHandler struct{}

func (n *NotSupportedHandler) SetNext(next DepositionHandler) DepositionHandler {
	return nil
}

func (n *NotSupportedHandler) Handle(deposit int) error {
	return fmt.Errorf("we dont have suitable cash for %d dollars!", deposit)
}
