package main

import (
	cmdDelivery "TaxiFares/internal/delivery/cmd"
	taxifaresUC "TaxiFares/internal/usecase/TaxiFares"
	"os"

	"log"
)

func main() {
	//init log file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	//init usecasd
	taxifaresUC, err := taxifaresUC.New()
	if err != nil {
		log.Fatal("failed init taxifares usecase")
	}

	//init delivery
	cmdDelivery, err := cmdDelivery.New(taxifaresUC)
	if err != nil {
		log.Fatal("failed init cmd delivery")
	}

	cmdDelivery.TaxiFares()
}
