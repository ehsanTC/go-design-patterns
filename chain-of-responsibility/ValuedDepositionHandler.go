package main

import (
	"errors"
	"fmt"
)

type ValuedDepositionHandler struct {
	depositionSupported int
	next                DepositionHandler
}

func (v *ValuedDepositionHandler) SetNext(next DepositionHandler) DepositionHandler {
	v.next = next
	return next
}

func (v *ValuedDepositionHandler) Handle(deposit int) error {
	var result error = nil // nil means successfull processing deposition
	var numOfCashes int = deposit / v.depositionSupported
	var reminder int = deposit % v.depositionSupported

	if reminder != 0 {
		if v.next != nil {
			result = v.next.Handle(reminder)
		} else {
			result = errors.New("the money can not be processed by the current cashes")
		}
	}

	if result == nil && numOfCashes != 0 {
		fmt.Printf("Your deposit contains %d numbers of %d dollars.\n", numOfCashes, v.depositionSupported)
	}

	return result
}
