package models

type Calculation struct {
	Action    string  `json:"action"`
	FistNum   int32   `json:"fistNum"`
	SecondNum int32   `json:"secondNum"`
	Numbers   []int32 `json:"numbers"`
}

const (
	ActionAdd            = "add"
	ActionMax            = "max"
	ActionAvg            = "average"
	ActionPrimeDecompose = "primeDecompose"
	ActionDivide         = "divide"
)
