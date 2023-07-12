package cmd

import "TaxiFares/internal/usecase"

type (
	CmdDelivery struct {
		taxiFaresUc usecase.ITaxeFares
	}
)
