package main

import (
	"fmt"

	"github.com/elangreza14/golang-basic-oop/bank"
	"github.com/elangreza14/golang-basic-oop/conventional"
	"github.com/elangreza14/golang-basic-oop/internet"
)

func main() {

	allBank := []bank.IBank{
		bank.NewBank(conventional.NewConventional("reza conven", 0.1)),
		bank.NewBank(internet.NewInternet("reza internet")),
	}

	for _, v := range allBank {
		err := v.FlowMoney(1000)
		fmt.Println(err)
		a := v.GetBallance()
		fmt.Println(a)
		err = v.FlowMoney(-900)
		fmt.Println(err)
		a = v.GetBallance()
		fmt.Println(a)
	}

}
