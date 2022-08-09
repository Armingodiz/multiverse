package calculatorService

import (
	"errors"
	"multiverse/core/models"
	"multiverse/core/services/calculatorService/client"
	"os"
)

type CalculatorService interface {
	Calculate(calculation models.Calculation) (interface{}, error)
}

func NewCalculatorService() CalculatorService {
	CalculatorServer := os.Getenv("CALCULATOR_SERVER")
	CalculatorPort := os.Getenv("CALCULATOR_PORT")
	CalculatorUseSSl := os.Getenv("CALCULATOR_USE_SSL")
	useSsl := false
	if CalculatorServer == "" {
		CalculatorServer = "Calculator"
	}
	if CalculatorPort == "" {
		CalculatorPort = "8081"
	}
	if CalculatorUseSSl == "" || CalculatorUseSSl == "false" {
		useSsl = false
	} else {
		useSsl = true
	}
	return &calculatorService{
		Host:   CalculatorServer,
		Port:   CalculatorPort,
		UseSSl: useSsl,
	}
}

type calculatorService struct {
	Host   string
	Port   string
	UseSSl bool
}

func (c *calculatorService) Calculate(calculation models.Calculation) (interface{}, error) {
	conn, err := client.NewCalculatorConnection(c.UseSSl, c.Host, c.Port)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cli, err := client.NewCalculatorClient(conn)
	if err != nil {
		return nil, err
	}
	switch calculation.Action {
	case "add":
		return cli.Add(calculation.FistNum, calculation.SecondNum)
	case "primeDecompose":
		return cli.PrimeNumberDecomposition(int64(calculation.FistNum))
	case "max":
		return cli.FindMaximum(calculation.Numbers)
	case "average":
		return cli.ComputeAverage(calculation.Numbers)
	case "divide":
		q, r, err := cli.Divide(calculation.FistNum, calculation.SecondNum)
		if err != nil {
			return nil, err
		}
		return map[string]int32{"q": q, "r": r}, nil
	default:
		return nil, errors.New("unsupported action")
	}
}
