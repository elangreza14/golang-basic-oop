package main

import (
	"fmt"

	"github.com/elangreza14/golang-basic-oop/bank"
	"github.com/elangreza14/golang-basic-oop/conventional"
	"github.com/elangreza14/golang-basic-oop/internet"
	elangrezaLogger "github.com/elangreza14/golang-logrus-custom"
)

func main() {
	logger := elangrezaLogger.NewLogger([]elangrezaLogger.PrefixLoggerName{
		{
			Title:       "app",
			Description: "golang basic oop",
		}})

	allBank := []bank.IBank{
		bank.NewBank(conventional.NewConventional("reza conven", 0.1)),
		bank.NewBank(internet.NewInternet("reza internet")),
	}
	fmt.Println(allBank)

	for _, v := range allBank {
		err := v.FlowMoney(1000)
		if err != nil {
			logger.Error("ERR: %v", err)
		}

		a := v.GetBallance()

		logger.Info("INFO BALLANCE: %v", a)

		err = v.FlowMoney(-900)
		if err != nil {
			logger.Error("ERR: %v", err)
		}

		a = v.GetBallance()
		logger.Info("INFO BALLANCE: %v", a)

	}

}
