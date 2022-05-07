package main

import (
	"github.com/elangreza14/golang-basic-oop/bank"
	"github.com/elangreza14/golang-basic-oop/conventional"
	"github.com/elangreza14/golang-basic-oop/internet"
	logger "github.com/elangreza14/golang-logrus-custom"
)

func main() {
	log := logger.NewLogger([]logger.PrefixLoggerName{{
		Title:       "",
		Description: "",
	}})
	allBank := []bank.IBank{
		bank.NewBank(conventional.NewConventional("reza conven", 0.1)),
		bank.NewBank(internet.NewInternet("reza internet")),
	}

	for _, v := range allBank {
		err := v.FlowMoney(1000)
		log.Error("", err)
		// fmt.Println(err)
		// a := v.GetBallance()
		// fmt.Println(a)
		// err = v.FlowMoney(-900)
		// fmt.Println(err)
		// a = v.GetBallance()
		// fmt.Println(a)
	}

}
