package chain_of_responsibility

import "fmt"

type ATM struct {
	Number int
}

func (A *ATM) Name() string {
	return "Chain of responsibility"
}

func (A *ATM) GetNumber() int {
	return A.Number
}

func (A *ATM) Work() {
	fmt.Println("\n*************** Thanks for using our ATM ***************")
	fmt.Println("********* We have only 50, 20 and 10 dollar cash *********")

	fmt.Print("\nHow much money do you want? ")
	var deposit int = 15
	_, err := fmt.Scan(&deposit)

	if err != nil {
		fmt.Println("The entered number is wrong!")
		return
	}

	depositProcessor := &DepositProcessor{}
	fiftyDollarsHandler := &ValuedDepositionHandler{depositionSupported: 50}
	twentyDollarsHandler := &ValuedDepositionHandler{depositionSupported: 20}
	tenDollarsHandler := &ValuedDepositionHandler{depositionSupported: 10}
	notSupportedHandler := &NotSupportedHandler{}

	fiftyDollarsHandler.SetNext(twentyDollarsHandler)
	twentyDollarsHandler.SetNext(tenDollarsHandler)
	tenDollarsHandler.SetNext(notSupportedHandler)

	depositProcessor.SetHandler(fiftyDollarsHandler)
	dErr := depositProcessor.Process(deposit)

	if dErr != nil {
		fmt.Printf("Error:%s\n", dErr)
	} else {
		fmt.Println("\nYour deposition was successful.")
	}
}
