package main

import "fmt"

func main() {
	fmt.Println("\n*************** Thanks for using our ATM ***************")
	fmt.Println("** We have only 50, 20 and 10 dollar cash **")

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
	d_err := depositProcessor.Process(deposit)

	if d_err != nil {
		fmt.Printf("Error:%s\n", d_err)
	} else {
		fmt.Println("\nYour deposition was successful.")
	}
}
