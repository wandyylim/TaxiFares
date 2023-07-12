package cmd

import "TaxiFares/internal/usecase"

func New(taxiFaresUc usecase.ITaxeFares) (*CmdDelivery, error) {
	return &CmdDelivery{
		taxiFaresUc: taxiFaresUc,
	}, nil
}
