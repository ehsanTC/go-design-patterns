package chain_of_responsibility

type DepositProcessor struct {
	handler DepositionHandler
}

func (d *DepositProcessor) SetHandler(handler DepositionHandler) {
	d.handler = handler
}

func (d *DepositProcessor) Process(deposit int) error {
	return d.handler.Handle(deposit)
}
