package models

type ACSFP struct {
	Id               int64   `json:"id"`
	Name             string  `json:"name"`
	QuantityAuth     int64   `json:"quantityAuth"`
	QuantityHeld     int64   `json:"quantityHeld"`
	RegisteredNumber string  `json:"registeredNumber"`
	YearOfProc       int64   `json:"yearOfProc"`
	Cost             float64 `json:"cost"`
	QuantityServed   int64   `json:"quantityServed"`
	ForwardTo        string  `json:"forwardTo"`
	DemandDetails    string  `json:"demandDetails"`
	Remarks          string  `json:"remarks"`
}
