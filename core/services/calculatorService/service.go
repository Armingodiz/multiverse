package calculatorService

import (
	"os"
)

type CalculatorService interface {
	Add(a, b int32) (int32, error)
	PrimeNumberDecomposition(number int64) (factores []int64, err error)
	ComputeAverage(nums []int32) (avg float64, err error)
	FindMaximum(inputChann chan int32) (max int32, err error)
	Divide(num, advisor int) (q, r int32, err error)
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

func (c *calculatorService) Add(a, b int32) (int32, error) {
	return 0, nil
}

func (c *calculatorService) PrimeNumberDecomposition(number int64) (factores []int64, err error) {
	return nil, nil
}

func (c *calculatorService) ComputeAverage(nums []int32) (avg float64, err error) {
	return 0, nil
}

func (c *calculatorService) FindMaximum(inputChann chan int32) (max int32, err error) {
	return 0, nil
}

func (c *calculatorService) Divide(num, advisor int) (q, r int32, err error) {
	return 0, 0, nil
}
